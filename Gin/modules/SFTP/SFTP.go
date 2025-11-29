package sftpTools

import (
	"archive/zip"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func SetDB(d *sql.DB) { db = d }
func SetDefaultPort(p int) { defaultSFTPPort = p }

type SFTPCreds struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Path     string `json:"path"`
}
type SFTPUser struct {
	ServerIP    string `json:"server_ip"`
	LibraryPath string `json:"library_path"`
}

var db *sql.DB
var (
	cachedSSHConn    *ssh.Client
	cachedSFTPClient *sftp.Client
	cachedCreds      *SFTPCreds
	cacheMu          sync.Mutex
	CurrentUser      User
	Sessions         = make(map[string]int)
	sftpClients      = make(map[int]*sftp.Client)
	mu               sync.Mutex
	sftpMu           sync.Mutex
	defaultSFTPPort  int
	sftpCmd          *exec.Cmd
	sftpCancel       context.CancelFunc
	sftpLastConfig   *SFTPConfig
	sftpLastStarted  time.Time
	sftpLastExitErr  error
)

type User struct {
	ID          int    `json:"id" db:"uid"`
	Username    string `json:"username" db:"username"`
	Password    string `json:"password_hash" db:"password_hash"`
	ServerIP    string `json:"server_ip" db:"server_ip"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	UpdatedAt   string `json:"updated_at" db:"updated_at"`
	IsAdmin     bool   `json:"isAdmin" db:"is_admin"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Email       string `json:"email" db:"email"`
	LibraryPath string `json:"library_path" db:"library_path"`
}
type SFTPConfig struct {
	Folder     string `json:"folder"`
	FolderName string `json:"folder_name"`
	Port       int    `json:"port" binding:"required"`
	User       string `json:"user"`
	Pass       string `json:"pass"`
}

type InstallRcloneRequest struct {
	PackageManager string `json:"package_manager" binding:"required"`
}

type PackageManagerInfo struct {
	Name           string `json:"name"`
	Installed      bool   `json:"installed"`
	ExampleInstall string `json:"example_install"`
}

func SftpEnvHandler(c *gin.Context) {
	goos := runtime.GOOS

	rclonePath, err := exec.LookPath("rclone")
	rcloneInstalled := (err == nil)

	defaultPort := defaultSFTPPort
	if defaultPort == 0 {
		defaultPort = 9824
	}
	if isPortInUse("127.0.0.1", defaultPort) {
		if p, err := pickFreePort(); err == nil {
			defaultPort = p
		}
	}

	var rcloneVersion string
	if rcloneInstalled {
		out, verr := exec.Command(rclonePath, "--version").CombinedOutput()
		if verr == nil {
			lines := strings.SplitN(string(out), "\n", 2)
			rcloneVersion = strings.TrimSpace(lines[0])
		}
	}

	pms := []PackageManagerInfo{}

	if goos == "darwin" {
		if _, err := exec.LookPath("brew"); err == nil {
			pms = append(pms, PackageManagerInfo{
				Name:           "brew",
				Installed:      true,
				ExampleInstall: "brew install rclone",
			})
		}
	} else if goos == "windows" {
		if _, err := exec.LookPath("winget"); err == nil {
			pms = append(pms, PackageManagerInfo{
				Name:           "winget",
				Installed:      true,
				ExampleInstall: `winget install --id Rclone.Rclone -e`,
			})
		}
	} else if goos == "linux" {
		if _, err := exec.LookPath("apt-get"); err == nil {
			pms = append(pms, PackageManagerInfo{
				Name:           "apt",
				Installed:      true,
				ExampleInstall: "sudo apt-get install -y rclone",
			})
		}
		if _, err := exec.LookPath("dnf"); err == nil {
			pms = append(pms, PackageManagerInfo{
				Name:           "dnf",
				Installed:      true,
				ExampleInstall: "sudo dnf install -y rclone",
			})
		}
		if _, err := exec.LookPath("pacman"); err == nil {
			pms = append(pms, PackageManagerInfo{
				Name:           "pacman",
				Installed:      true,
				ExampleInstall: "sudo pacman -S --noconfirm rclone",
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"os":               goos,
		"rclone_installed": rcloneInstalled,
		"rclone_path":      rclonePath,
		"rclone_version":   rcloneVersion,
		"package_managers": pms,
		"default_port":     defaultPort,
	})
}

func InstallRcloneHandler(c *gin.Context) {
	var req InstallRcloneRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload", "details": err.Error()})
		return
	}

	goos := runtime.GOOS
	var cmd *exec.Cmd

	switch goos {
	case "darwin":
		if req.PackageManager != "brew" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "on macOS only 'brew' is supported here"})
			return
		}
		cmd = exec.Command("brew", "install", "rclone")
	case "windows":
		if req.PackageManager != "winget" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "on Windows only 'winget' is supported here"})
			return
		}
		cmd = exec.Command("winget", "install", "--id", "Rclone.Rclone", "-e")
	case "linux":
		switch req.PackageManager {
		case "apt":
			cmd = exec.Command("sudo", "apt-get", "install", "-y", "rclone")
		case "dnf":
			cmd = exec.Command("sudo", "dnf", "install", "-y", "rclone")
		case "pacman":
			cmd = exec.Command("sudo", "pacman", "-S", "--noconfirm", "rclone")
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported package manager for linux"})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported OS for auto-install"})
		return
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "failed to install rclone",
			"details": err.Error(),
			"output":  string(output),
		})
		return
	}

	if _, err := exec.LookPath("rclone"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "rclone still not found after installation",
			"details": err.Error(),
			"output":  string(output),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "installed",
		"os":     goos,
		"output": string(output),
	})
}
func isPortInUse(host string, port int) bool {
	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", addr, 500*time.Millisecond)
	if err == nil {
		_ = conn.Close()
		return true
	}
	return false
}

func StartLocalSFTPHandler(c *gin.Context) {
	var cfg SFTPConfig
	if err := c.ShouldBindJSON(&cfg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload", "details": err.Error()})
		return
	}

	var absFolder string
	if strings.TrimSpace(cfg.Folder) != "" {
		f, err := filepath.Abs(cfg.Folder)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid-folder", "details": err.Error()})
			return
		}
		absFolder = f
	} else if strings.TrimSpace(cfg.FolderName) != "" {
		home, err := os.UserHomeDir()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "no-home-dir", "details": err.Error()})
			return
		}
		absFolder = filepath.Join(home, cfg.FolderName)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "missing-folder",
			"details": "No folder or folder_name provided.",
		})
		return
	}

	if _, err := os.Stat(absFolder); os.IsNotExist(err) {
		if err := os.MkdirAll(absFolder, 0755); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "failed-to-create-folder",
				"details": err.Error(),
			})
			return
		}
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "failed-to-stat-folder",
			"details": err.Error(),
		})
		return
	}

	if strings.TrimSpace(cfg.User) == "" {
		cfg.User = "FlacPlayerUser"
	}
	if strings.TrimSpace(cfg.Pass) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "missing-password",
			"details": "Password is required when starting local SFTP.",
		})
		return
	}

	if cfg.Port <= 0 {
		preferredPort := defaultSFTPPort
		if preferredPort == 0 {
			preferredPort = 9824
		}
		if isPortInUse("127.0.0.1", preferredPort) {
			port, err := pickFreePort()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "port-allocation-failed", "details": err.Error()})
				return
			}
			cfg.Port = port
		} else {
			cfg.Port = preferredPort
		}
	}

	if isPortInUse("127.0.0.1", cfg.Port) {
		c.JSON(http.StatusConflict, gin.H{
			"error":   "port-in-use",
			"details": fmt.Sprintf("TCP port %d is already in use. Stop any existing SFTP server or choose another port.", cfg.Port),
		})
		return
	}

	if _, err := exec.LookPath("rclone"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "rclone-not-found",
			"details": "rclone not found in PATH. Install it or add it to PATH.",
		})
		return
	}

	if sftpCmd != nil && sftpCmd.Process != nil {
		log.Println("[sftp] stopping existing rclone process")
		if sftpCancel != nil {
			sftpCancel()
		}
		_ = sftpCmd.Process.Kill()
		sftpCmd = nil
	}

	ctx, cancel := context.WithCancel(context.Background())

	cmd := exec.CommandContext(ctx,
		"rclone",
		"serve", "sftp",
		absFolder,
		"--addr", fmt.Sprintf(":%d", cfg.Port),
		"--user", cfg.User,
		"--pass", cfg.Pass,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		cancel()
		sftpLastExitErr = err
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "failed-to-start-rclone",
			"details": err.Error(),
		})
		return
	}

	if err := waitForTCPPort("127.0.0.1", cfg.Port, 5*time.Second); err != nil {
		log.Printf("[sftp] rclone started but port not reachable: %v\n", err)
		_ = cmd.Process.Kill()
		cancel()
		sftpCmd = nil
		sftpCancel = nil
		sftpLastExitErr = err

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "sftp-port-unreachable",
			"details": err.Error(),
		})
		return
	}

	sftpCmd = cmd
	sftpCancel = cancel
	sftpLastConfig = &cfg
	sftpLastStarted = time.Now()
	sftpLastExitErr = nil

	log.Printf("[sftp] rclone serving %s on port %d as %s\n", absFolder, cfg.Port, cfg.User)

	go func() {
		err := cmd.Wait()
		sftpMu.Lock()
		defer sftpMu.Unlock()
		if err != nil {
			log.Printf("[sftp] rclone serve sftp exited with error: %v\n", err)
			sftpLastExitErr = err
		}
		sftpCmd = nil
		if sftpCancel != nil {
			sftpCancel()
			sftpCancel = nil
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":     "running",
		"folder":     absFolder,
		"port":       cfg.Port,
		"user":       cfg.User,
		"started_at": sftpLastStarted.Format(time.RFC3339),
	})
}
func waitForTCPPort(host string, port int, timeout time.Duration) error {
	addr := fmt.Sprintf("%s:%d", host, port)
	deadline := time.Now().Add(timeout)

	for {
		conn, err := net.DialTimeout("tcp", addr, 500*time.Millisecond)
		if err == nil {
			_ = conn.Close()
			return nil
		}

		if time.Now().After(deadline) {
			return fmt.Errorf("timeout waiting for %s: %w", addr, err)
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func pickFreePort() (int, error) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}
	defer l.Close()

	_, portStr, err := net.SplitHostPort(l.Addr().String())
	if err != nil {
		return 0, err
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return 0, err
	}

	return port, nil
}
func StopLocalSFTPHandler(c *gin.Context) {
	sftpMu.Lock()
	defer sftpMu.Unlock()

	if sftpCmd == nil || sftpCmd.Process == nil {
		c.JSON(http.StatusOK, gin.H{"status": "not_running"})
		return
	}

	if sftpCancel != nil {
		sftpCancel()
	}
	_ = sftpCmd.Process.Kill()
	sftpCmd = nil
	sftpCancel = nil

	c.JSON(http.StatusOK, gin.H{"status": "stopped"})
}

// Body: { host, port, username, password, path }
func SftpCredsHandler(c *gin.Context) {
	var req SFTPCreds
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"connected": false,
			"error":     "bad-request",
			"details":   err.Error(),
		})
		return
	}

	if strings.TrimSpace(req.Host) == "" || req.Port <= 0 ||
		strings.TrimSpace(req.Username) == "" || strings.TrimSpace(req.Password) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"connected": false,
			"error":     "invalid-input",
			"details":   "host, port, username and password are required",
		})
		return
	}

	if strings.TrimSpace(req.Path) == "" {
		req.Path = "/"
	}

	user, err := GetCurrentUser(c)
	if err != nil {
		log.Println("SftpCredsHandler GetCurrentUser error:", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"connected": false,
			"error":     "unauthorized",
			"details":   err.Error(),
		})
		return
	}

	addr := fmt.Sprintf("%s:%d", req.Host, req.Port)
	sshConfig := &ssh.ClientConfig{
		User:            req.Username,
		Auth:            []ssh.AuthMethod{ssh.Password(req.Password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	}

	conn, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		log.Println("SFTP connect error (ssh dial):", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"connected": false,
			"error":     "connection-failed",
			"details":   err.Error(),
		})
		return
	}
	defer conn.Close()

	client, err := sftp.NewClient(conn)
	if err != nil {
		log.Println("SFTP client error:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"connected": false,
			"error":     "sftp-client-failed",
			"details":   err.Error(),
		})
		return
	}
	defer client.Close()

	if err := EnsureRemoteDir(client, req.Path); err != nil {
		log.Println("EnsureRemoteDir error:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"connected": false,
			"error":     "invalid-remote-path",
			"details":   err.Error(),
		})
		return
	}

	encPass, err := EncryptWithMasterKey(req.Password)
	if err != nil {
		log.Println("EncryptWithMasterKey error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"connected": false,
			"error":     "encryption-failed",
			"details":   err.Error(),
		})
		return
	}

	_, err = db.Exec(`
		INSERT INTO user_sftp (uid, server_ip, server_port, sftp_user, sftp_password_enc, library_path)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (uid) DO UPDATE
		SET server_ip = EXCLUDED.server_ip,
		    server_port = EXCLUDED.server_port,
		    sftp_user = EXCLUDED.sftp_user,
		    sftp_password_enc = EXCLUDED.sftp_password_enc,
		    library_path = EXCLUDED.library_path
	`, user.ID, req.Host, req.Port, req.Username, encPass, req.Path)

	if err != nil {
		log.Println("DB upsert user_sftp error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"connected": false,
			"error":     "db-error",
			"details":   err.Error(),
		})
		return
	}

	mu.Lock()
	sftpClients[user.ID] = client
	mu.Unlock()

	c.JSON(http.StatusOK, gin.H{
		"connected": true,
		"details":   "SFTP reachable and credentials saved",
	})
}

func LocalSFTPStatusHandler(c *gin.Context) {
	sftpMu.Lock()
	defer sftpMu.Unlock()

	running := sftpCmd != nil && sftpCmd.Process != nil
	resp := gin.H{"running": running}

	if sftpLastConfig != nil {
		resp["folder"] = sftpLastConfig.Folder
		resp["port"] = sftpLastConfig.Port
		resp["user"] = sftpLastConfig.User
		resp["started_at"] = sftpLastStarted.Format(time.RFC3339)
	}

	if sftpLastExitErr != nil {
		resp["last_error"] = sftpLastExitErr.Error()
	}

	c.JSON(http.StatusOK, resp)
}

func DownloadFromSFTP(userID int, remotePath, localPath string) error {
	return withSFTPClient(userID, func(client *sftp.Client) error {
		src, err := client.Open(remotePath)
		if err != nil {
			ResetSFTPConnection()
			return err
		}
		defer src.Close()

		dst, err := os.Create(localPath)
		if err != nil {
			return err
		}
		defer dst.Close()

		_, err = io.Copy(dst, src)
		return err
	})
}

func WithSFTPFile(userID int, remotePath string, fn func(*sftp.File, os.FileInfo) error) error {
	return withSFTPClient(userID, func(client *sftp.Client) error {
		f, err := client.Open(remotePath)
		if err != nil {
			ResetSFTPConnection()
			return err
		}
		defer f.Close()

		info, err := f.Stat()
		if err != nil {
			return err
		}

		return fn(f, info)
	})
}

func RemoteFileExists(userID int, remotePath string) (bool, error) {
	exists := false
	err := withSFTPClient(userID, func(client *sftp.Client) error {
		_, err := client.Stat(remotePath)
		if err != nil {
			if os.IsNotExist(err) {
				exists = false
				return nil
			}
			ResetSFTPConnection()
			return err
		}
		exists = true
		return nil
	})
	return exists, err
}
func DownloadSFTPDirectory(userID int, remoteDir, localDir string) error {
	return withSFTPClient(userID, func(client *sftp.Client) error {
		walker := client.Walk(remoteDir)
		for walker.Step() {
			if walker.Err() != nil {
				return walker.Err()
			}

			rel, _ := filepath.Rel(remoteDir, walker.Path())
			localPath := filepath.Join(localDir, rel)

			if walker.Stat().IsDir() {
				os.MkdirAll(localPath, 0755)
				continue
			}

			src, err := client.Open(walker.Path())
			if err != nil {
				return err
			}
			defer src.Close()

			os.MkdirAll(filepath.Dir(localPath), 0755)
			dst, err := os.Create(localPath)
			if err != nil {
				return err
			}

			io.Copy(dst, src)
			dst.Close()
		}
		return nil
	})
}
func ZipDirectory(src, dst string) error {
	zipfile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		rel, _ := filepath.Rel(src, path)
		header, _ := zip.FileInfoHeader(info)
		header.Name = rel
		header.Method = zip.Deflate

		writer, _ := archive.CreateHeader(header)
		file, _ := os.Open(path)
		defer file.Close()

		io.Copy(writer, file)
		return nil
	})
	return nil
}

func UploadAtomic(userID int, localPath, remotePath string) error {
	return withSFTPClient(userID, func(client *sftp.Client) error {
		tmp := remotePath + ".uploading"

		src, err := os.Open(localPath)
		if err != nil {
			return err
		}
		defer src.Close()

		dst, err := client.Create(tmp)
		if err != nil {
			return err
		}

		_, err = io.Copy(dst, src)
		dst.Close()
		if err != nil {
			client.Remove(tmp)
			return err
		}

		if err := client.PosixRename(tmp, remotePath); err != nil {
			client.Remove(tmp)
			return err
		}
		return nil
	})
}

func UploadToSFTP(userID int, localPath, remotePath string) error {
	client, err := GetSFTPClient(userID)
	if err != nil {
		return err
	}

	src, err := os.Open(localPath)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := client.OpenFile(remotePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE)
	if err != nil {
		ResetSFTPConnection()
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}
func GetCurrentUser(c *gin.Context) (*User, error) {
	userIDVal, ok := c.Get("userID")
	if !ok {
		return nil, fmt.Errorf("unauthorized: missing userID in context")
	}

	userID, ok := userIDVal.(int)
	if !ok {
		return nil, fmt.Errorf("invalid userID type")
	}

	user, err := GetUserByID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user: %w", err)
	}

	return user, nil
}
func GetUserByID(id int) (*User, error) {
	var u User
	err := db.QueryRow(`SELECT uid, username, password_hash, is_admin FROM users WHERE uid=$1`, id).
		Scan(&u.ID, &u.Username, &u.Password, &u.IsAdmin)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func GetUserBySession(sessionID string) (int, bool) {
	userID, ok := Sessions[sessionID]
	return userID, ok
}

func EncryptWithMasterKey(plainText string) (string, error) {
	key := []byte(os.Getenv("SFTP_MASTER_KEY"))
	if len(key) == 0 {
		return "", fmt.Errorf("missing SFTP_MASTER_KEY")
	}

	if decoded, err := base64.StdEncoding.DecodeString(string(key)); err == nil && len(decoded) > 0 {
		key = decoded
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(plainText), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func DecryptWithMasterKey(b64 string) (string, error) {
	keyEnv := os.Getenv("SFTP_MASTER_KEY")
	if keyEnv == "" {
		return "", fmt.Errorf("missing SFTP_MASTER_KEY")
	}

	key := []byte(keyEnv)
	if decoded, err := base64.StdEncoding.DecodeString(keyEnv); err == nil && len(decoded) > 0 {
		key = decoded
	}

	raw, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesgcm.NonceSize()
	if len(raw) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce := raw[:nonceSize]
	ciphertext := raw[nonceSize:]
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
func GetUserSFTPCredsFromDB(userID int) (*SFTPCreds, error) {
	var host, username, encPass, remotePath sql.NullString
	var port sql.NullInt64
	err := db.QueryRow(`SELECT server_ip, server_port, sftp_user, sftp_password_enc, library_path FROM user_sftp WHERE uid = $1`, userID).
		Scan(&host, &port, &username, &encPass, &remotePath)
	if err == sql.ErrNoRows {
		return nil, errors.New("no-credentials")
	}
	if err != nil {
		return nil, err
	}
	if !host.Valid {
		return nil, errors.New("no-credentials")
	}
	p := 2222
	if port.Valid && port.Int64 > 0 {
		p = int(port.Int64)
	}
	pass := ""
	if encPass.Valid && encPass.String != "" {
		clear, derr := DecryptWithMasterKey(encPass.String)
		if derr != nil {
			return nil, derr
		}
		pass = clear
	} else {
		return nil, errors.New("no-credentials")
	}
	rpath := "/"
	if remotePath.Valid && strings.TrimSpace(remotePath.String) != "" {
		rpath = remotePath.String
	}
	// fmt.Println("DEBUG CREDS:", host, p, username.String, rpath)
	return &SFTPCreds{
		Host:     host.String,
		Port:     p,
		Username: username.String,
		Password: pass,
		Path:     rpath,
	}, nil
}

func GetSFTPClient(userID int) (*sftp.Client, error) {
	creds, err := GetUserSFTPCredsFromDB(userID)
	if err != nil {
		fmt.Println("DEBUG: GetSFTPClient no creds for user", userID)
		return nil, errors.New("no-credentials")
	}
	// fmt.Println("DEBUG CREDS:", creds)
	fmt.Println("DEBUG: Connecting to SFTP", creds.Host, creds.Port, creds.Username)

	addr := fmt.Sprintf("%s:%d", creds.Host, creds.Port)
	conn, err := ssh.Dial("tcp", addr, &ssh.ClientConfig{
		User:            creds.Username,
		Auth:            []ssh.AuthMethod{ssh.Password(creds.Password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	})
	if err != nil {
		fmt.Println("DEBUG: SSH Dial failed:", err)
		return nil, err
	}

	client, err := sftp.NewClient(conn)
	if err != nil {
		fmt.Println("DEBUG: NewClient failed:", err)
		return nil, err
	}

	return client, nil
}
func withSFTPClient(userID int, fn func(*sftp.Client) error) error {
	mu.Lock()
	client, ok := sftpClients[userID]
	mu.Unlock()

	if !ok {
		fmt.Println("DEBUG: no cached client, creating new one for user", userID)
		newClient, err := GetSFTPClient(userID)
		if err != nil {
			fmt.Println("DEBUG: GetSFTPClient failed:", err)
			return err
		}
		mu.Lock()
		sftpClients[userID] = newClient
		mu.Unlock()
		client = newClient
	} else {
		fmt.Println("DEBUG: reusing cached client for user", userID)
	}

	return fn(client)
}

func ResetSFTPConnection() {
	cacheMu.Lock()
	defer cacheMu.Unlock()

	if cachedSFTPClient != nil {
		cachedSFTPClient.Close()
	}
	if cachedSSHConn != nil {
		cachedSSHConn.Close()
	}

	cachedSFTPClient = nil
	cachedSSHConn = nil
}

func UploadToUserSFTPSync(userID int, localPath, remoteFilename string) error {
	client, err := GetSFTPClient(userID)
	if err != nil {
		return err
	}
	defer client.Close()

	tempName := remoteFilename + ".uploading"

	localFile, err := os.Open(localPath)
	if err != nil {
		return err
	}
	defer localFile.Close()

	remoteTemp, err := client.Create(tempName)
	if err != nil {
		return err
	}

	_, err = io.Copy(remoteTemp, localFile)
	remoteTemp.Close()
	if err != nil {
		client.Remove(tempName)
		return err
	}

	err = client.PosixRename(tempName, remoteFilename)
	if err != nil {
		client.Remove(tempName)
		return err
	}

	return nil
}
func deterministicStorageKey(uid int64, origFilename string, bytes []byte) string {
	h := sha256.New()
	h.Write([]byte(strconv.FormatInt(uid, 10)))
	h.Write([]byte{0xFF})
	h.Write([]byte(origFilename))
	h.Write([]byte{0xFF})

	if len(bytes) > 32 {
		h.Write(bytes[:32])
	} else {
		h.Write(bytes)
	}

	return hex.EncodeToString(h.Sum(nil))
}

func EnsureRemoteDir(client *sftp.Client, remotePath string) error {
	remotePath = path.Clean(remotePath)
	if remotePath == "." || remotePath == "/" {
		return nil
	}
	parts := strings.Split(remotePath, "/")
	cur := "/"
	for _, p := range parts {
		if p == "" {
			continue
		}
		cur = path.Join(cur, p)
		if _, err := client.Stat(cur); err != nil {
			if os.IsNotExist(err) || strings.Contains(err.Error(), "no such file") {
				if err := client.Mkdir(cur); err != nil {
					if !strings.Contains(err.Error(), "file exists") {
						return err
					}
				}
			} else {
				return err
			}
		}
	}
	return nil
}

func UploadToUserSFTP(userID int, localPath, filename string) error {
	log.Printf("[sftp] UploadToUserSFTP user=%d local=%s filename=%s\n", userID, localPath, filename)

	creds, err := GetUserSFTPCredsFromDB(userID)
	if err != nil {
		log.Printf("[sftp] GetUserSFTPCredsFromDB error for user=%d: %v\n", userID, err)
		return fmt.Errorf("no creds: %w", err)
	}

	log.Printf(
		"[sftp] creds for user=%d host=%s port=%d user=%s path=%s\n",
		userID, creds.Host, creds.Port, creds.Username, creds.Path,
	)

	config := &ssh.ClientConfig{
		User:            creds.Username,
		Auth:            []ssh.AuthMethod{ssh.Password(creds.Password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	}

	addr := fmt.Sprintf("%s:%d", creds.Host, creds.Port)
	log.Printf("[sftp] ssh.Dial addr=%s\n", addr)
	conn, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Printf("[sftp] ssh.Dial FAILED addr=%s err=%v\n", addr, err)
		return fmt.Errorf("ssh dial: %w", err)
	}
	log.Printf("[sftp] ssh.Dial OK addr=%s\n", addr)
	defer conn.Close()

	sftpClient, err := sftp.NewClient(conn)
	if err != nil {
		log.Printf("[sftp] sftp.NewClient FAILED: %v\n", err)
		return fmt.Errorf("sftp client: %w", err)
	}
	log.Printf("[sftp] sftp.NewClient OK\n")
	defer sftpClient.Close()

	log.Printf("[sftp] EnsureRemoteDir base=%s\n", creds.Path)
	if err := EnsureRemoteDir(sftpClient, creds.Path); err != nil {
		log.Printf("[sftp] EnsureRemoteDir FAILED: %v\n", err)
		return fmt.Errorf("ensure remote dir: %w", err)
	}

	remoteFilePath := path.Join(creds.Path, filename)
	log.Printf("[sftp] remoteFilePath=%s\n", remoteFilePath)

	src, err := os.Open(localPath)
	if err != nil {
		log.Printf("[sftp] open local FAILED: %v\n", err)
		return fmt.Errorf("open local: %w", err)
	}
	defer src.Close()

	dst, err := sftpClient.OpenFile(remoteFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC)
	if err != nil {
		log.Printf("[sftp] open remote FAILED: %v\n", err)
		return fmt.Errorf("open remote: %w", err)
	}
	defer dst.Close()

	n, err := io.Copy(dst, src)
	if err != nil {
		log.Printf("[sftp] copy FAILED after %d bytes: %v\n", n, err)
		return fmt.Errorf("copy: %w", err)
	}
	log.Printf("[sftp] ✅ UploadToUserSFTP done user=%d remote=%s bytes=%d\n", userID, remoteFilePath, n)

	return nil
}

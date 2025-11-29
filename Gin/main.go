package main

import (
	sftpTools "MusicAppGin/modules/SFTP"
	"MusicAppGin/modules/tagedit"
	"archive/zip"
	"log"
	"net"
	"path"
	"path/filepath"

	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mewkiz/flac"
	"github.com/mewkiz/flac/meta"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// Struct
type Album struct {
	Album_id        int    `json:"album_id"`
	Album_name      string `json:"album_name"`
	Album_cover     string `json:"album_cover"`
	Album_genre     string `json:"album_genre"`
	Album_subgenres string `json:"album_subgenres"`
	Album_artist    string `json:"album_artist"`
	Album_date      string `json:"album_date"`
	Album_comment   string `json:"album_comment"`
	Album_copyright string `json:"album_copyright"`
	Songs           []Song `json:"songs"`
	DateAdded       string `json:"date_added"`
	LastModified    string `json:"last_modified"`
}

type Song struct {
	SongID      int    `json:"song_id"`
	Title       string `json:"title"`
	Duration    string `json:"duration"`
	TrackNumber int    `json:"track_number"`
	FileName    string `json:"file_name"`
	Data        []byte `json:"-"`
	Missing     bool   `json:"missing,omitempty"`
}

type PendingAlbum struct {
	AlbumName string `json:"album_name"`
	Cover     string `json:"album_cover"`
	Genre     string `json:"album_genre"`
	Artist    string `json:"album_artist"`
	Year      string `json:"album_date"`
	Comment   string `json:"album_comment"`
	Copyright string `json:"album_copyright"`
	SongsPA   []Song `json:"songs"`
}

type PortsConfig struct {
	APIPort      int    `json:"api_port"`
	FrontendPort int    `json:"frontend_port"`
	SFTPPort     int    `json:"sftp_port"`
	DBURL        string `json:"db_url"`
}

var portsConfig PortsConfig

type UploadStatus struct {
	Artist           string `json:"artist"`
	Title            string `json:"title"`
	Done             bool   `json:"done"`
	OriginalFilename string `json:"-"`
}
type TreeSong struct {
	Artist string              `json:"artist"`
	Albums map[string][]string `json:"albums"`
}

var (
	db             *sql.DB
	uploadProgress = make(map[int][]UploadStatus)
	mu             sync.Mutex
	pendingAlbums  = map[int]map[string]*PendingAlbum{}
)

const (
	defaultFirstName   = "Guest"
	defaultUsername    = "guest"
	defaultServerIP    = "localhost"
	defaultLastName    = "User"
	defaultEmail       = "guest@example.com"
	defaultLibraryPath = "/guest/library"
)

func getUserByUsername(username string) (*sftpTools.User, error) {
	var user sftpTools.User
	err := db.QueryRow(
		"SELECT uid, username, password_hash, is_admin FROM users WHERE username = $1",
		username,
	).Scan(&user.ID, &user.Username, &user.Password, &user.IsAdmin)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func hashPasswordSHA256(pw string) string {
	h := sha256.Sum256([]byte(pw))
	return hex.EncodeToString(h[:])
}

// func checkPasswordSHA256(stored, received string) bool {
// 	fmt.Printf("🔑 Stored hash:   %s\n", stored)
// 	fmt.Printf("➡ Received hash: %s\n", received)
// 	return stored == received
// }

// FLAC Cover from file
func extractAlbumCoverFromReader(r io.Reader) string {
	stream, err := flac.Parse(r)
	if err != nil {
		return ""
	}

	for _, block := range stream.Blocks {
		if picture, ok := block.Body.(*meta.Picture); ok {
			b64 := base64.StdEncoding.EncodeToString(picture.Data)
			return "data:image/jpeg;base64," + b64
		}
	}

	return ""
}

func extractMetadataFromReader(r io.Reader, flacName string) (
	title, album, genre, artist, date, comment, copyright, tracknumber string,
) {
	stream, err := flac.Parse(r)
	if err != nil {
		return flacName, "Unknown Album", "Unknown Genre", "Unknown Artist", "Unknown Year", "", "", ""
	}

	for _, block := range stream.Blocks {
		if vc, ok := block.Body.(*meta.VorbisComment); ok {
			for _, tag := range vc.Tags {
				key := strings.ToLower(tag[0])
				value := tag[1]
				switch key {
				case "title":
					title = value
				case "album":
					album = value
				case "genre":
					genre = value
				case "artist":
					artist = value
				case "albumartist":
					if artist == "" {
						artist = value
					}
				case "date", "year":
					date = value
				case "label", "publisher":
					if copyright == "" {
						copyright = "Published by " + value
					}
				case "comment", "description":
					comment = value
				case "copyright":
					copyright = value
				case "tracknumber":
					fixed := strings.TrimSpace(value)
					re := regexp.MustCompile(`^(\d+)`)
					match := re.FindStringSubmatch(fixed)
					if len(match) > 1 {
						tracknumber = match[1]
					} else {
						tracknumber = fixed
					}

				}
			}
		}
	}

	if title == "" {
		title = flacName
	}
	if album == "" {
		album = "Unknown Album"
	}
	if genre == "" {
		genre = "Unknown Genre"
	}
	if artist == "" {
		artist = "Unknown Artist"
	}
	if date == "" {
		date = "Unknown Year"
	}
	if copyright == "" {
		copyright = "Unknown legal owner, refer to " + artist
	}

	return
}

// func extractSongInfoFromReader(r io.Reader, flacName string) Song {
// 	title, _, _, _, _, _, _, _ := extractMetadataFromReader(r, flacName)
// 	duration := getFLACDuration(r)
// 	return Song{Title: title, Duration: duration}
// }

func getFLACDuration(r io.Reader) string {
	data, err := io.ReadAll(r)
	if err != nil {
		return "00:00"
	}

	stream, err := flac.Parse(bytes.NewReader(data))
	if err != nil || stream.Info == nil {
		return "00:00"
	}

	ns := stream.Info.NSamples
	sr := stream.Info.SampleRate
	if ns == 0 || sr == 0 {
		return "00:00"
	}

	totalSecs := float64(ns) / float64(sr)
	return fmt.Sprintf("%02d:%02d", int(totalSecs)/60, int(totalSecs)%60)
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("session_id")
		// fmt.Println("DEBUG: sessionID cookie =", sessionID, "err =", err)
		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		userID, ok := sftpTools.GetUserBySession(sessionID)
		// fmt.Println("DEBUG: userID from sessions =", userID, "ok =", ok)
		if !ok {
			c.JSON(401, gin.H{"error": "Invalid session"})
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}

func ensureGuestUser() (*sftpTools.User, error) {
	const defaultUsername = "guest"
	const defaultFirstName = "Guest"
	const defaultLastName = "User"
	const defaultEmail = "guest@example.com"

	defaultPassword := hashPasswordSHA256("guest")

	var oldGuestID int
	err := db.QueryRow(`SELECT uid FROM users WHERE username=$1`, defaultUsername).Scan(&oldGuestID)
	if err == nil {
		rows, qErr := db.Query(`SELECT file_path FROM user_music WHERE user_id=$1`, oldGuestID)
		if qErr == nil {
			defer rows.Close()

			client, cErr := sftpTools.GetSFTPClient(oldGuestID)
			if cErr == nil {
				defer client.Close()

				for rows.Next() {
					var remotePath string
					if err := rows.Scan(&remotePath); err != nil {
						continue
					}
					_ = client.Remove(remotePath)
				}
			}
		}

		if _, err := db.Exec(`DELETE FROM user_music WHERE user_id=$1`, oldGuestID); err != nil {
			fmt.Println("❌ Failed to delete old guest songs:", err)
		}
		if _, err := db.Exec(`DELETE FROM user_sftp WHERE uid=$1`, oldGuestID); err != nil {
			fmt.Println("❌ Failed to delete old guest sftp creds:", err)
		}
		if _, err := db.Exec(`DELETE FROM users WHERE uid=$1`, oldGuestID); err != nil {
			fmt.Println("❌ Failed to delete old guest user:", err)
		}
	}

	var userID int
	err = db.QueryRow(`
        INSERT INTO users(username, password_hash, first_name, last_name, email)
        VALUES($1,$2,$3,$4,$5)
        RETURNING uid
    `, defaultUsername, defaultPassword, defaultFirstName, defaultLastName, defaultEmail).Scan(&userID)
	if err != nil {
		fmt.Println("❌ Failed to create guest user:", err)
		return nil, err
	}

	fmt.Printf("✅ Created guest user: ID=%d, username=%s\n", userID, defaultUsername)
	sftpTools.CurrentUser = sftpTools.User{
		ID:       userID,
		Username: defaultUsername,
		IsAdmin:  false,
	}
	return &sftpTools.CurrentUser, nil
}

func findPgBin() string {
	candidates := []string{
		"/opt/homebrew/opt/postgresql@14/bin",
		"/opt/homebrew/opt/postgresql/bin",
		"/usr/local/opt/postgresql@14/bin",
		"/usr/local/opt/postgresql/bin",
	}
	for _, dir := range candidates {
		if info, err := os.Stat(dir); err == nil && info.IsDir() {
			return dir
		}
	}
	return ""
}

func initLocalDBIfNeeded() {
	pgBin := findPgBin()
	if pgBin == "" {
		log.Printf("postgres binaries not found; skipping auto DB init")
		return
	}
	createuser := filepath.Join(pgBin, "createuser")
	createdb := filepath.Join(pgBin, "createdb")
	psql := filepath.Join(pgBin, "psql")

	run := func(cmd string, args ...string) {
		c := exec.Command(cmd, args...)
		_ = c.Run()
	}

	run(createuser, "-s", "musicuser")
	run(createdb, "-O", "musicuser", "musicdb")

	schemaCandidates := []string{
		filepath.Join("sql", "schema.sql"),
	}
	if exe, err := os.Executable(); err == nil {
		base := filepath.Dir(exe)
		schemaCandidates = append(schemaCandidates,
			filepath.Join(base, "..", "share", "musicapp", "sql", "schema.sql"),
			filepath.Join(base, "..", "sql", "schema.sql"),
		)
	}
	for _, schemaPath := range schemaCandidates {
		if _, err := os.Stat(schemaPath); err == nil {
			run(psql, "-d", "musicdb", "-f", schemaPath)
			break
		}
	}
}

func tryStartPostgresService() {
	services := []string{"postgresql@14", "postgresql"}
	for _, svc := range services {
		cmd := exec.Command("brew", "services", "start", svc)
		if err := cmd.Run(); err == nil {
			log.Printf("attempted to start %s via brew services", svc)
			return
		}
	}
}

func openDBWithRetry(databaseURL string) (*sql.DB, error) {
	var lastErr error
	for i := 0; i < 3; i++ {
		db, err := sql.Open("postgres", databaseURL)
		if err != nil {
			lastErr = err
			time.Sleep(2 * time.Second)
			continue
		}
		if err = db.Ping(); err != nil {
			lastErr = err
			if strings.Contains(err.Error(), "connection refused") && i == 0 {
				tryStartPostgresService()
			}
			time.Sleep(2 * time.Second)
			continue
		}
		return db, nil
	}
	return nil, lastErr
}
func buildTreeFromPending(pending map[string]*PendingAlbum) []TreeSong {
	tree := []TreeSong{}
	for _, album := range pending {
		var artistNode *TreeSong
		for i := range tree {
			if tree[i].Artist == album.Artist {
				artistNode = &tree[i]
				break
			}
		}
		if artistNode == nil {
			artistNode = &TreeSong{
				Artist: album.Artist,
				Albums: make(map[string][]string),
			}
			tree = append(tree, *artistNode)
			artistNode = &tree[len(tree)-1]
		}

		if _, ok := artistNode.Albums[album.AlbumName]; !ok {
			artistNode.Albums[album.AlbumName] = []string{}
		}

		existing := map[string]bool{}
		for _, s := range artistNode.Albums[album.AlbumName] {
			existing[s] = true
		}
		for _, song := range album.SongsPA {
			if !existing[song.Title] {
				artistNode.Albums[album.AlbumName] = append(artistNode.Albums[album.AlbumName], song.Title)
				existing[song.Title] = true
			}
		}
	}

	return tree
}

func buildTreeString(tree []TreeSong) string {
	var sb strings.Builder

	for i, artistNode := range tree {
		lastArtist := i == len(tree)-1
		artistPrefix := "├── "
		if lastArtist {
			artistPrefix = "└── "
		}
		sb.WriteString(artistPrefix + artistNode.Artist + "\n")

		albums := make([]string, 0, len(artistNode.Albums))
		for album := range artistNode.Albums {
			albums = append(albums, album)
		}
		sort.Strings(albums)

		for j, album := range albums {
			lastAlbum := j == len(albums)-1
			albumPrefix := "│   ├── "
			if lastArtist {
				albumPrefix = "    ├── "
			}
			if lastAlbum {
				if lastArtist {
					albumPrefix = "    └── "
				} else {
					albumPrefix = "│   └── "
				}
			}
			sb.WriteString(albumPrefix + album + "\n")

			songs := artistNode.Albums[album]
			sort.Strings(songs)
			for k, song := range songs {
				lastSong := k == len(songs)-1
				songPrefix := "│   │   ├── "
				if lastAlbum {
					if lastArtist {
						songPrefix = "        ├── "
					} else {
						songPrefix = "│       ├── "
					}
					if lastSong {
						if lastArtist {
							songPrefix = "        └── "
						} else {
							songPrefix = "│       └── "
						}
					}
				}
				sb.WriteString(songPrefix + song + "\n")
			}
		}
	}

	return sb.String()
}

// Takes Discogs result map["genre"], map["style"], adds:
//   - "main_genre": string
//   - "subgenres": []string
func enrichGenres(r map[string]interface{}) {
	var all []string

	if graw, ok := r["genre"].([]interface{}); ok {
		for _, v := range graw {
			if s, ok := v.(string); ok && strings.TrimSpace(s) != "" {
				all = append(all, s)
			}
		}
	}
	if sraw, ok := r["style"].([]interface{}); ok {
		for _, v := range sraw {
			if s, ok := v.(string); ok && strings.TrimSpace(s) != "" {
				all = append(all, s)
			}
		}
	}

	raw := strings.TrimSpace(strings.Join(all, ", "))
	if raw == "" {
		r["main_genre"] = "Unknown"
		r["subgenres"] = []string{}
		return
	}

	parts := strings.FieldsFunc(raw, func(ch rune) bool {
		return ch == ',' || ch == ';' || ch == '/' || ch == '&'
	})

	cleaned := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			cleaned = append(cleaned, p)
		}
	}

	if len(cleaned) == 0 {
		r["main_genre"] = "Unknown"
		r["subgenres"] = []string{}
		return
	}

	main := cleaned[0]
	sub := []string{}
	for _, p := range cleaned[1:] {
		if p != "" && p != main {
			sub = append(sub, p)
		}
	}

	r["main_genre"] = main
	r["subgenres"] = sub
}

func enrichCopyright(r map[string]interface{}) {
	year, _ := r["year"].(string)
	labelsRaw, _ := r["label"].([]interface{})
	labels := make([]string, 0, len(labelsRaw))
	seen := map[string]struct{}{}

	for _, v := range labelsRaw {
		if s, ok := v.(string); ok {
			s = strings.TrimSpace(s)
			if s == "" {
				continue
			}
			if _, exists := seen[s]; exists {
				continue
			}
			seen[s] = struct{}{}
			labels = append(labels, s)
		}
	}

	if len(labels) == 0 && year == "" {
		return
	}

	technical := map[string]struct{}{
		"MPO": {},
	}

	primary := make([]string, 0, len(labels))
	for _, l := range labels {
		if _, isTech := technical[l]; !isTech {
			primary = append(primary, l)
		}
	}

	main := ""
	extra := []string{}

	if len(primary) > 0 {
		main = primary[0]
		if len(primary) > 1 {
			extra = primary[1:]
		}
	} else if len(labels) > 0 {
		main = labels[0]
		if len(labels) > 1 {
			extra = labels[1:]
		}
	}

	var copyright string
	if main != "" {
		if len(extra) > 0 {
			if year != "" {
				copyright = fmt.Sprintf("℗ %s %s & %s", year, main, strings.Join(extra, ", "))
			} else {
				copyright = fmt.Sprintf("℗ %s & %s", main, strings.Join(extra, ", "))
			}
		} else {
			if year != "" {
				copyright = fmt.Sprintf("℗ %s %s", year, main)
			} else {
				copyright = fmt.Sprintf("℗ %s", main)
			}
		}
	} else if year != "" {
		copyright = fmt.Sprintf("℗ %s", year)
	}

	if copyright != "" {
		r["copyright"] = copyright
	}
}

func audioContentTypeFromName(name string) string {
	switch strings.ToLower(filepath.Ext(name)) {
	case ".flac":
		return "audio/flac"
	case ".mp3":
		return "audio/mpeg"
	case ".wav":
		return "audio/wav"
	case ".ogg":
		return "audio/ogg"
	case ".m4a":
		return "audio/mp4"
	default:
		return "application/octet-stream"
	}
}

func DownloadSongByID(c *gin.Context) {
	userID := c.GetInt("userID")
	if userID == 0 {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	songIDStr := c.Query("songId")
	if songIDStr == "" {
		c.JSON(400, gin.H{"error": "missing songId"})
		return
	}

	songID, err := strconv.Atoi(songIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid songId"})
		return
	}

	var remotePath string
	err = db.QueryRow(`SELECT file_path FROM user_music WHERE user_id=$1 AND song_id=$2`,
		userID, songID).Scan(&remotePath)

	if err != nil {
		c.JSON(404, gin.H{"error": "song not found"})
		return
	}

	name := filepath.Base(remotePath)
	tmp := filepath.Join(os.TempDir(), name)

	err = sftpTools.DownloadFromSFTP(userID, remotePath, tmp)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.FileAttachment(tmp, name)
}
func StreamSong(c *gin.Context) {
	userID := c.GetInt("userID")

	songIDStr := c.Param("songId")
	songID, err := strconv.Atoi(songIDStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid song id"})
		return
	}

	var remotePath string
	if err := db.QueryRow(
		`SELECT file_path FROM user_music WHERE user_id=$1 AND song_id=$2`,
		userID, songID,
	).Scan(&remotePath); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{"error": "song not found"})
			return
		}
		c.JSON(500, gin.H{"error": "failed to resolve song path"})
		return
	}

	name := filepath.Base(remotePath)
	err = sftpTools.WithSFTPFile(userID, remotePath, func(f *sftp.File, info os.FileInfo) error {
		fmt.Println("[StreamSong] user=%d song=%d path=%s size=%d bytes\n",
			userID, songID, remotePath, info.Size())
		c.Header("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", name))
		c.Header("Content-Type", audioContentTypeFromName(name))
		c.Header("Accept-Ranges", "bytes")
		http.ServeContent(c.Writer, c.Request, name, info.ModTime(), f)
		return nil
	})
	if err != nil {
		c.JSON(500, gin.H{"error": "stream failed", "details": err.Error()})
		return
	}
}

func DownloadAlbum(c *gin.Context) {
	albumID := c.Query("albumId")
	filePath := "/path/to/albums/" + albumID + ".zip"

	c.Header("Content-Disposition", "attachment; filename=album_"+albumID+".zip")
	c.Header("Content-Type", "application/zip")
	c.File(filePath)
}

func generateRandomSessionID() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}

var consumerKey = "VUOxAzdEicDBOftGbPyW"
var consumerSecret = "ORrKNsJuKlUOzemlXjdMiYfpGNLoisoi"

func SearchReleases(artist, album string) ([]map[string]interface{}, error) {
	strictResults, err := searchStrict(artist, album)
	if err != nil {
		return nil, err
	}

	fuzzyResults, err := searchFuzzy(artist, album)
	if err != nil {
		return nil, err
	}

	merged := mergeResults(strictResults, fuzzyResults)

	if len(merged) > 12 {
		merged = merged[:12]
	}

	return merged, nil
}
func searchStrict(artist, album string) ([]map[string]interface{}, error) {
	base := "https://api.discogs.com/database/search"

	params := url.Values{}
	params.Set("artist", artist)
	params.Set("release_title", album)
	params.Set("per_page", "6")

	reqURL := fmt.Sprintf("%s?%s", base, params.Encode())
	return performDiscogsRequest(reqURL)
}

func searchFuzzy(artist, album string) ([]map[string]interface{}, error) {
	base := "https://api.discogs.com/database/search"

	q := strings.TrimSpace(artist + " " + album)

	params := url.Values{}
	params.Set("q", q)
	params.Set("per_page", "12")

	reqURL := fmt.Sprintf("%s?%s", base, params.Encode())
	fmt.Println("Request URL to Discogs:", reqURL)

	return performDiscogsRequest(reqURL)
}
func performDiscogsRequest(reqURL string) ([]map[string]interface{}, error) {
	if consumerKey == "" || consumerSecret == "" {
		return nil, fmt.Errorf("discogs credentials not configured")
	}

	req, _ := http.NewRequest("GET", reqURL, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Discogs key=%s, secret=%s", consumerKey, consumerSecret))
	req.Header.Set("User-Agent", "YourAppName/1.0")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var body map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, err
	}

	raw, _ := body["results"].([]interface{})

	results := []map[string]interface{}{}
	for _, x := range raw {
		if m, ok := x.(map[string]interface{}); ok {
			results = append(results, m)
		}
	}

	return results, nil
}
func mergeResults(strict, fuzzy []map[string]interface{}) []map[string]interface{} {
	seen := map[int]bool{}
	merged := []map[string]interface{}{}

	add := func(list []map[string]interface{}) {
		for _, item := range list {
			id, _ := item["id"].(float64)
			iid := int(id)
			if !seen[iid] {
				seen[iid] = true
				merged = append(merged, item)
			}
		}
	}

	add(strict)
	add(fuzzy)

	return merged
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

func portsConfigPath() (string, error) {
	if custom := os.Getenv("MUSICAPP_CONFIG"); custom != "" {
		return custom, nil
	}
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "musicapp", "ports.json"), nil
}

func loadPortsConfig() (PortsConfig, error) {
	cfg := PortsConfig{}
	path, err := portsConfigPath()
	if err != nil {
		return cfg, err
	}

	_ = os.MkdirAll(filepath.Dir(path), 0o755)
	if data, err := os.ReadFile(path); err == nil {
		_ = json.Unmarshal(data, &cfg)
	}

	if cfg.APIPort == 0 {
		if p, err := pickFreePort(); err == nil {
			cfg.APIPort = p
		}
	}
	if cfg.FrontendPort == 0 {
		if p, err := pickFreePort(); err == nil {
			cfg.FrontendPort = p
		}
	}
	if cfg.SFTPPort == 0 {
		if p, err := pickFreePort(); err == nil {
			cfg.SFTPPort = p
		}
	}
	if cfg.DBURL == "" {
		cfg.DBURL = "postgres://musicuser:musicuser@localhost/musicdb?sslmode=disable"
	}

	if data, err := json.MarshalIndent(cfg, "", "  "); err == nil {
		_ = os.WriteFile(path, data, 0o644)
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
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

func main() {
	var err error
	_ = godotenv.Load()
	portsConfig, err = loadPortsConfig()
	if err != nil {
		log.Printf("failed to load ports config, using defaults: %v", err)
		portsConfig = PortsConfig{}
	}
	if portsConfig.APIPort == 0 {
		portsConfig.APIPort = 8080
	}
	if portsConfig.FrontendPort == 0 {
		portsConfig.FrontendPort = 4173
	}
	if portsConfig.SFTPPort == 0 {
		portsConfig.SFTPPort = 9824
	}

	consumerKey = os.Getenv("DISCOGS_KEY")
	consumerSecret = os.Getenv("DISCOGS_SECRET")
	databaseURL := getEnv("DATABASE_URL", portsConfig.DBURL)

	// Attempt local DB bootstrap if using defaults
	if databaseURL == "" || strings.Contains(databaseURL, "musicuser") {
		initLocalDBIfNeeded()
	}

	db, err = openDBWithRetry(databaseURL)
	if err != nil {
		log.Printf("Failed to connect to DB: %v", err)
	} else {
		tagedit.SetDB(db)
		sftpTools.SetDB(db)
		sftpTools.SetDefaultPort(portsConfig.SFTPPort)
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	r.MaxMultipartMemory = 512 << 20
	r.Use(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Accept, Origin, Cache-Control, X-CSRF-Token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	r.GET("/config/ports", func(c *gin.Context) {
		c.JSON(200, portsConfig)
	})
	r.GET("/api/sftp/song", AuthRequired(), DownloadSongByID)
	r.GET("/stream/:songId", AuthRequired(), StreamSong)
	r.GET("/api/sftp/album", AuthRequired(), func(c *gin.Context) {
		albumIDParam := c.Query("albumId")
		var albumID int
		fmt.Sscanf(albumIDParam, "%d", &albumID)
		currentUser, err := sftpTools.GetCurrentUser(c)
		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		rows, err := db.Query(`        SELECT um.song_id, um.file_path 
        FROM user_music um
        JOIN songs s ON um.song_id = s.song_id
        WHERE s.album_id=$1 AND um.user_id=$2`, albumID, currentUser.ID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch songs"})
			return
		}
		defer rows.Close()

		buf := new(bytes.Buffer)
		zipWriter := zip.NewWriter(buf)

		for rows.Next() {
			var songID int
			var remotePath string
			rows.Scan(&songID, &remotePath)

			tmpFile := fmt.Sprintf("/tmp/%d.tmp", songID)
			err := sftpTools.DownloadFromSFTP(sftpTools.CurrentUser.ID, remotePath, tmpFile)
			if err != nil {
				continue
			}

			f, _ := zipWriter.Create(path.Base(remotePath))
			data, _ := os.ReadFile(tmpFile)
			f.Write(data)

			os.Remove(tmpFile)
		}
		zipWriter.Close()

		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=album_%d.zip", albumID))
		c.Data(200, "application/zip", buf.Bytes())
	})
	r.POST("/api/albums/:id/metadata", func(c *gin.Context) {
		albumID := c.Param("id")
		if albumID == "" {
			c.JSON(400, gin.H{"error": "missing album id"})
			return
		}

		var reqBody tagedit.ApplyMetadataRequest
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(400, gin.H{"error": "invalid JSON", "details": err.Error()})
			return
		}

		if reqBody.UseDiscogsNotes && reqBody.DiscogsID != 0 {
			releaseURL := fmt.Sprintf("https://api.discogs.com/releases/%d", reqBody.DiscogsID)
			req, _ := http.NewRequest("GET", releaseURL, nil)
			req.Header.Set("Authorization",
				fmt.Sprintf("Discogs key=%s, secret=%s", consumerKey, consumerSecret))
			req.Header.Set("User-Agent", "YourAppName/1.0")

			resp, err := http.DefaultClient.Do(req)
			if err == nil && resp.StatusCode == http.StatusOK {
				defer resp.Body.Close()
				var full map[string]interface{}
				if err := json.NewDecoder(resp.Body).Decode(&full); err == nil {
					if notes, ok := full["notes"].(string); ok {
						notes = strings.TrimSpace(notes)
						if notes != "" {
							reqBody.AlbumComment = notes //
						}
					}
				}
			}
		}
	})
	r.GET("/api/sftp/album/zip", AuthRequired(), func(c *gin.Context) {
		albumIDParam := c.Query("albumId")
		var albumID int
		if _, err := fmt.Sscanf(albumIDParam, "%d", &albumID); err != nil {
			c.JSON(400, gin.H{"error": "Invalid albumId"})
			return
		}

		currentUser, err := sftpTools.GetCurrentUser(c)
		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		rows, err := db.Query(`
        SELECT um.song_id, um.file_path, s.title
        FROM user_music um
        JOIN songs s ON um.song_id = s.song_id
        WHERE s.album_id=$1 AND um.user_id=$2
    `, albumID, currentUser.ID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch songs"})
			return
		}
		defer rows.Close()

		c.Header("Content-Type", "application/zip")
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=album_%d.zip", albumID))
		zipWriter := zip.NewWriter(c.Writer)
		defer zipWriter.Close()

		for rows.Next() {
			var songID int
			var remotePath, title string
			if err := rows.Scan(&songID, &remotePath, &title); err != nil {
				continue
			}

			tmpPath := fmt.Sprintf("./temp_album_%d_%d", albumID, songID)
			if err := sftpTools.DownloadFromSFTP(currentUser.ID, remotePath, tmpPath); err != nil {
				continue
			}
			defer os.Remove(tmpPath)

			f, err := os.Open(tmpPath)
			if err != nil {
				continue
			}

			fw, err := zipWriter.Create(filepath.Base(title + filepath.Ext(remotePath)))
			if err != nil {
				f.Close()
				continue
			}

			_, _ = io.Copy(fw, f)
			f.Close()
		}
	})

	r.GET("/api/discogs/search", func(c *gin.Context) {
		artist := strings.TrimSpace(c.Query("artist"))
		album := strings.TrimSpace(c.Query("album"))
		if artist == "" || album == "" {
			c.JSON(400, gin.H{"error": "missing query parameters"})
			return
		}

		strictResults, err := SearchReleases(artist, album)
		if err != nil {
			c.JSON(500, gin.H{"error": "Discogs search failed", "details": err.Error()})
			return
		}

		generalResults := strictResults
		if len(strictResults) == 0 {
			generalResults, err = SearchReleases("", album)
			if err != nil {
				c.JSON(500, gin.H{"error": "Discogs fallback search failed", "details": err.Error()})
				return
			}
		}

		for _, r := range generalResults {
			enrichCopyright(r)
			enrichGenres(r)
		}

		c.JSON(200, generalResults)
	})
	r.Static("/static", "./static")

	distDir := getEnv("DIST_DIR", "./dist")
	indexFile := filepath.Join(distDir, "index.html")
	assetsDir := filepath.Join(distDir, "assets")
	if _, err := os.Stat(assetsDir); err == nil {
		r.Static("/assets", assetsDir)
	}
	if _, err := os.Stat(filepath.Join(distDir, "vite.svg")); err == nil {
		r.StaticFile("/vite.svg", filepath.Join(distDir, "vite.svg"))
	}

	go startFrontendServer(distDir, indexFile)

	r.POST("/sftp/creds", AuthRequired(), func(c *gin.Context) {
		user, err := sftpTools.GetCurrentUser(c)
		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorized"})
			return
		}

		var creds sftpTools.SFTPCreds
		body, _ := io.ReadAll(c.Request.Body)
		fmt.Println("RAW BODY:", string(body))
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		if err := c.ShouldBindJSON(&creds); err != nil {
			c.JSON(400, gin.H{
				"error":   "invalid-payload",
				"details": err.Error(),
			})
			return
		}

		if creds.Host == "" || creds.Port == 0 || creds.Username == "" || creds.Password == "" {
			c.JSON(400, gin.H{"error": "missing-fields"})
			return
		}

		addr := fmt.Sprintf("%s:%d", creds.Host, creds.Port)
		config := &ssh.ClientConfig{
			User:            creds.Username,
			Auth:            []ssh.AuthMethod{ssh.Password(creds.Password)},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         10 * time.Second,
		}

		conn, err := ssh.Dial("tcp", addr, config)
		if err != nil {
			log.Println("SFTP ssh.Dial error:", err)

			var msg string
			switch {
			case strings.Contains(err.Error(), "unable to authenticate"):
				msg = "SSH login failed: wrong username/password for this SFTP server."
			case strings.Contains(err.Error(), "connection refused"):
				msg = "SFTP server is not listening on this host/port."
			case strings.Contains(err.Error(), "handshake failed"),
				strings.Contains(err.Error(), "EOF"):
				msg = "SSH handshake failed (server rejected the connection). " +
					"Make sure the username/password in the SFTP modal match the ones rclone was started with."
			default:
				msg = err.Error()
			}

			c.JSON(http.StatusBadRequest, gin.H{
				"connected": false,
				"error":     "connection-failed",
				"details":   msg,
			})
			return
		}

		defer conn.Close()

		encPass, err := sftpTools.EncryptWithMasterKey(creds.Password)
		if err != nil {
			c.JSON(500, gin.H{"error": "encryption-failed", "details": err.Error()})
			return
		}

		_, err = db.Exec(`
        INSERT INTO user_sftp (server_ip, server_port, sftp_user, sftp_password_enc, library_path, uid)
        VALUES ($1, $2, $3, $4, $5, $6)
        ON CONFLICT (uid) DO UPDATE
        SET server_ip = EXCLUDED.server_ip,
            server_port = EXCLUDED.server_port,
            sftp_user = EXCLUDED.sftp_user,
            sftp_password_enc = EXCLUDED.sftp_password_enc,
            library_path = EXCLUDED.library_path
    `, creds.Host, creds.Port, creds.Username, encPass, creds.Path, user.ID)
		if err != nil {
			c.JSON(500, gin.H{"error": "db-error", "details": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"connected": true,
			"status":    "ok",
		})
	})

	r.GET("/api/sftp/status", AuthRequired(), func(c *gin.Context) {
		user, err := sftpTools.GetCurrentUser(c)
		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorized"})
			return
		}
		var host, username, path sql.NullString
		var port sql.NullInt64
		err = db.QueryRow(`SELECT server_ip, server_port, sftp_user, library_path FROM user_sftp WHERE uid=$1`, user.ID).
			Scan(&host, &port, &username, &path)
		if err == sql.ErrNoRows {
			c.JSON(200, gin.H{
				"status":   "missing",
				"host":     "",
				"port":     0,
				"username": "",
				"path":     "",
			})
			return
		}
		if err != nil {
			c.JSON(500, gin.H{"error": "query failed"})
			return
		}
		c.JSON(200, gin.H{
			"status":   "exists",
			"host":     host.String,
			"port":     port.Int64,
			"username": username.String,
			"path":     path.String,
		})
	})

	r.POST("/api/register", func(c *gin.Context) {
		var form struct {
			FirstName       string `json:"firstName"`
			LastName        string `json:"lastName"`
			Email           string `json:"email"`
			Username        string `json:"username"`
			Password        string `json:"password"`
			ConfirmPassword string `json:"confirmPassword"`
			LibraryPath     string `json:"libraryPath"`
		}
		if err := c.ShouldBindJSON(&form); err != nil {
			log.Printf("[register] invalid payload: %v", err)
			c.JSON(400, gin.H{
				"error":   "Invalid payload",
				"details": err.Error(),
			})
			return
		}

		var exists bool
		err := db.QueryRow(
			`SELECT EXISTS(SELECT 1 FROM users WHERE username=$1 OR email=$2)`,
			form.Username, form.Email,
		).Scan(&exists)
		if err != nil {
			log.Printf("[register] exists check failed for %s: %v", form.Username, err)
			c.JSON(500, gin.H{
				"error":   "Database error",
				"details": err.Error(),
			})
			return
		}
		if exists {
			c.JSON(400, gin.H{"error": "Email already exists(if not try different usermname)"})
			return
		}

		hashed := hashPasswordSHA256(form.Password)
		var userID int
		err = db.QueryRow(`
        INSERT INTO users(username, password_hash, first_name, last_name, email, library_path, server_ip)
        VALUES($1,$2,$3,$4,$5,$6,$7)
        RETURNING uid
    `, form.Username, hashed, form.FirstName, form.LastName, form.Email, form.LibraryPath, defaultServerIP).Scan(&userID)
		if err != nil {
			log.Printf("[register] create user failed username=%s email=%s: %v", form.Username, form.Email, err)
			c.JSON(500, gin.H{
				"error":   "Failed to create user",
				"details": err.Error(),
			})
			return
		}

		sftpTools.CurrentUser = sftpTools.User{ID: userID, Username: form.Username, IsAdmin: false}

		c.JSON(200, gin.H{
			"id":       userID,
			"username": form.Username,
			"isAdmin":  false,
		})
	})

	r.POST("/login", func(c *gin.Context) {
		var payload struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := c.BindJSON(&payload); err != nil {
			c.JSON(400, gin.H{"error": "Invalid payload"})
			return
		}

		fmt.Println("Username", payload.Username)

		// 1) Special case: guest login
		if payload.Username == "guest" {
			const guestHash = "84983c60f7daadc1cb8698621f802c0d9f9a3c3c295c810748fb048115c186ec"

			if payload.Password != guestHash {
				c.JSON(401, gin.H{"error": "Invalid guest token"})
				return
			}

			guestUser, err := ensureGuestUser()
			if err != nil {
				log.Printf("[login/guest] init guest failed: %v", err)
				c.JSON(500, gin.H{
					"error":   "Failed to initialize guest",
					"details": err.Error(),
				})
				return
			}
			sessionID := generateRandomSessionID()
			sftpTools.Sessions[sessionID] = guestUser.ID

			c.SetCookie("session_id", sessionID, 3600*24, "/", "", false, true)

			c.JSON(200, gin.H{
				"id":       guestUser.ID,
				"username": guestUser.Username,
				"isAdmin":  guestUser.IsAdmin,
			})
			return
		}

		// 2) Normal users
		user, err := getUserByUsername(payload.Username)
		if err != nil || user.Password != payload.Password {
			c.JSON(401, gin.H{"error": "Invalid username or password"})
			return
		}

		sessionID := generateRandomSessionID()
		sftpTools.Sessions[sessionID] = user.ID

		c.SetCookie("session_id", sessionID, 3600*24, "/", "", false, true)

		c.JSON(200, gin.H{
			"id":       user.ID,
			"username": user.Username,
			"isAdmin":  user.IsAdmin,
		})
	})

	r.POST("/logout", func(c *gin.Context) {
		sessionID, err := c.Cookie("session_id")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not logged in"})
			return
		}

		delete(sftpTools.Sessions, sessionID)

		c.SetCookie("session_id", "", -1, "/", "", false, true)

		c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
	})

	r.GET("/profile", AuthRequired(), func(c *gin.Context) {
		userID := c.GetInt("userID")
		// fmt.Println("Debug: userID =", userID)

		var u struct {
			ID          int    `json:"id"`
			Username    string `json:"username"`
			FirstName   string `json:"first_name"`
			LastName    string `json:"last_name"`
			Email       string `json:"email"`
			Server      string `json:"server"`
			LibraryPath string `json:"library_path"`
			SFTPUser    string `json:"sftp_user"`
			IsAdmin     bool   `json:"is_admin"`
		}

		var serverIP, serverPort, sftpUser, libraryPath sql.NullString

		// fmt.Println("Debug: running query")
		err := db.QueryRow(`
		SELECT u.uid, u.username, u.first_name, u.last_name, u.email,
		       s.server_ip, s.server_port, s.sftp_user, s.library_path, u.is_admin
		FROM users u
		LEFT JOIN user_sftp s ON u.uid = s.uid
		WHERE u.uid=$1
	`, userID).Scan(
			&u.ID, &u.Username, &u.FirstName, &u.LastName, &u.Email,
			&serverIP, &serverPort, &sftpUser, &libraryPath, &u.IsAdmin,
		)

		if err != nil {
			// fmt.Println("Debug: query error:", err)
			c.JSON(500, gin.H{"error": "Failed to fetch user"})
			return
		}

		ip := ""
		port := ""
		if serverIP.Valid {
			ip = serverIP.String
		}
		if serverPort.Valid {
			port = serverPort.String
		}
		if ip != "" && port != "" {
			u.Server = fmt.Sprintf("%s:%s", ip, port)
		} else if ip != "" {
			u.Server = ip
		}

		if sftpUser.Valid {
			u.SFTPUser = sftpUser.String
		}
		if libraryPath.Valid {
			u.LibraryPath = libraryPath.String
		}

		c.JSON(200, u)
	})

	r.PUT("/profile", AuthRequired(), func(c *gin.Context) {
		userID := c.GetInt("userID")
		// fmt.Println("Debug: userID =", userID)

		var payload struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Email     string `json:"email"`
			Server    string `json:"server"`
			SFTPUser  string `json:"sftp_user"`
		}

		if err := c.BindJSON(&payload); err != nil {
			// fmt.Println("Debug: bind error:", err)
			c.JSON(400, gin.H{"error": "Invalid payload"})
			return
		}

		_, err := db.Exec(`
		UPDATE users
		SET first_name=$1, last_name=$2, email=$3, updated_at=NOW()
		WHERE uid=$4
	`, payload.FirstName, payload.LastName, payload.Email, userID)
		if err != nil {
			// fmt.Println("Debug: update users error:", err)
			c.JSON(500, gin.H{"error": "Failed to update user info"})
			return
		}

		serverIP := ""
		serverPort := ""
		if payload.Server != "" {
			parts := strings.Split(payload.Server, ":")
			serverIP = parts[0]
			if len(parts) > 1 {
				serverPort = parts[1]
			}
		}

		_, err = db.Exec(`
		INSERT INTO user_sftp(uid, server_ip, server_port, sftp_user)
		VALUES($1, $2, $3, $4)
		ON CONFLICT(uid)
		DO UPDATE SET server_ip=$2, server_port=$3, sftp_user=$4
	`, userID, serverIP, serverPort, payload.SFTPUser)
		if err != nil {
			// fmt.Println("Debug: upsert user_sftp error:", err)
			c.JSON(500, gin.H{"error": "Failed to update SFTP info"})
			return
		}

		c.JSON(200, gin.H{"success": true})
	})

	r.GET("/api/check-username", func(c *gin.Context) {
		username := c.Query("username")
		if username == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "username parameter required"})
			return
		}

		var exists bool
		err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)", username).Scan(&exists)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"exists": exists})
	})

	r.GET("/library", AuthRequired(), func(c *gin.Context) {
		userID := c.GetInt("userID")
		// fmt.Println("DEBUG: userID from context =", userID)

		var user sftpTools.User
		err := db.QueryRow(`SELECT uid, username, is_admin FROM users WHERE uid=$1`, userID).
			Scan(&user.ID, &user.Username, &user.IsAdmin)
		if err != nil {
			fmt.Println("ERROR fetching user:", err)
			c.JSON(500, gin.H{"error": "Failed to fetch user"})
			return
		}
		// fmt.Println("DEBUG: fetched user:", user)

		albumsMap := make(map[int]Album)
		albumKeys := []int{}
		rows, err := db.Query(`
SELECT s.song_id, s.title, s.album_id, a.name, a.cover_base64, a.year, a.genre, ar.name, s.duration, a.comment, a.copyright, um.file_path
FROM user_music um
JOIN songs s ON um.song_id = s.song_id
JOIN albums a ON s.album_id = a.album_id
JOIN artists ar ON a.artist_id = ar.artist_id
WHERE um.user_id = $1
ORDER BY um.uploaded_at DESC
    `, user.ID)
		if err != nil {
			fmt.Println("ERROR fetching albums:", err)
			c.JSON(500, gin.H{"error": "Failed to fetch albums"})
			return
		}
		defer rows.Close()
		// fmt.Println("DEBUG: rows fetched successfully")

		missingCache := make(map[string]bool)
		for rows.Next() {
			var songID, albumID int
			var title, albumName, artistName, albumGenre, duration string
			var coverBase64, year, comment, copyright sql.NullString
			var remotePath sql.NullString

			if err := rows.Scan(&songID, &title, &albumID, &albumName, &coverBase64, &year, &albumGenre, &artistName, &duration, &comment, &copyright, &remotePath); err != nil {
				fmt.Println("ERROR scanning row:", err)
				continue
			}
			// fmt.Printf("DEBUG: scanned row - songID=%d, albumID=%d, title=%s\n", songID, albumID, title)

			missing := false
			if remotePath.Valid && strings.TrimSpace(remotePath.String) != "" {
				if cached, ok := missingCache[remotePath.String]; ok {
					missing = !cached
				} else {
					exists, err := sftpTools.RemoteFileExists(user.ID, remotePath.String)
					if err != nil {
						fmt.Println("ERROR checking SFTP file:", err)
					}
					missingCache[remotePath.String] = exists
					missing = !exists
				}
			} else {
				missing = true
			}

			song := Song{SongID: songID, Title: title, Duration: duration, Missing: missing}
			if existing, ok := albumsMap[albumID]; ok {
				existing.Songs = append(existing.Songs, song)
				albumsMap[albumID] = existing
			} else {
				albumsMap[albumID] = Album{
					Album_id:        albumID,
					Album_name:      albumName,
					Album_genre:     albumGenre,
					Album_artist:    artistName,
					Album_comment:   comment.String,
					Album_copyright: copyright.String,
					Album_cover:     coverBase64.String,
					Album_date:      year.String,
					Songs:           []Song{song},
				}
				albumKeys = append(albumKeys, albumID)
			}
		}
		dateRows, err := db.Query(`
    SELECT s.album_id,
           MIN(um.uploaded_at) AS date_added,
           MAX(um.modified_at) AS last_modified
    FROM user_music um
    JOIN songs s ON um.song_id = s.song_id
    WHERE um.user_id = $1
    GROUP BY s.album_id
`, user.ID)
		if err != nil {
			fmt.Println("ERROR fetching album dates:", err)
		} else {
			defer dateRows.Close()

			for dateRows.Next() {
				var albumID int
				var dateAdded, lastModified sql.NullTime

				if err := dateRows.Scan(&albumID, &dateAdded, &lastModified); err != nil {
					fmt.Println("ERROR scanning album date row:", err)
					continue
				}

				if alb, ok := albumsMap[albumID]; ok {
					if dateAdded.Valid {
						alb.DateAdded = dateAdded.Time.Format(time.RFC3339)
					}
					if lastModified.Valid {
						alb.LastModified = lastModified.Time.Format(time.RFC3339)
					}
					albumsMap[albumID] = alb
				}
			}

			if err := dateRows.Err(); err != nil {
				fmt.Println("ERROR iterating date rows:", err)
			}
		}

		albums := []Album{}
		for _, key := range albumKeys {
			albums = append(albums, albumsMap[key])
		}

		fmt.Println("DEBUG: albums map built, count =", len(albums))

		if len(albums) == 0 {
			imgFile := "./static/img/default.jpg"
			f, err := os.Open(imgFile)
			if err == nil {
				defer f.Close()
				img, _, _ := image.Decode(f)
				buf := new(bytes.Buffer)
				jpeg.Encode(buf, img, &jpeg.Options{Quality: 80})
				coverBase64 := "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(buf.Bytes())
				albums = append(albums, Album{Album_name: "No albums", Album_cover: coverBase64})
			}
		}

		c.JSON(200, gin.H{
			"albums":   albums,
			"username": user.Username,
			"user_id":  user.ID,
		})

	})

	r.GET("/api/albums/:id", AuthRequired(), func(c *gin.Context) {
		id := c.Param("id")
		idStr := c.Param("id")
		albumID, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid album ID"})
			return
		}

		userID := c.GetInt("userID")
		if userID == 0 {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		var album Album

		err = db.QueryRow(`
    SELECT a.album_id, a.name, a.year, a.genre, a.cover_base64, ar.name, a.comment, a.copyright, a.subgenres
    FROM albums a
    JOIN artists ar ON a.artist_id = ar.artist_id
    WHERE a.album_id=$1
`, albumID).Scan(
			&album.Album_id,
			&album.Album_name,
			&album.Album_date,
			&album.Album_genre,
			&album.Album_cover,
			&album.Album_artist,
			&album.Album_comment,
			&album.Album_copyright,
			&album.Album_subgenres,
		)

		if err != nil {
			c.JSON(404, gin.H{"error": "Album not found"})
			return
		}

		rows, err := db.Query(`
        SELECT s.song_id, s.title, s.duration, s.track_number, um.file_path
        FROM songs s
        JOIN user_music um ON s.song_id = um.song_id
        WHERE s.album_id = $1 AND um.user_id = $2
        ORDER BY s.track_number
    `, id, userID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch songs"})
			return
		}
		defer rows.Close()

		album.Songs = []Song{}
		missingCache := make(map[string]bool)
		for rows.Next() {
			var s Song
			var remotePath sql.NullString
			if err := rows.Scan(&s.SongID, &s.Title, &s.Duration, &s.TrackNumber, &remotePath); err != nil {
				continue
			}
			if remotePath.Valid && strings.TrimSpace(remotePath.String) != "" {
				if cached, ok := missingCache[remotePath.String]; ok {
					s.Missing = !cached
				} else {
					exists, err := sftpTools.RemoteFileExists(userID, remotePath.String)
					if err != nil {
						fmt.Println("ERROR checking SFTP file:", err)
					}
					missingCache[remotePath.String] = exists
					s.Missing = !exists
				}
			} else {
				s.Missing = true
			}
			album.Songs = append(album.Songs, s)
		}
		var dateAdded, dateModified string
		err = db.QueryRow(`
    SELECT 
        COALESCE(MIN(um.uploaded_at)::text, '') AS date_added,
        COALESCE(MAX(um.modified_at)::text, '') AS date_modified
    FROM user_music um
    JOIN songs s ON um.song_id = s.song_id
    WHERE s.album_id = $1 AND um.user_id = $2
`, albumID, userID).Scan(&dateAdded, &dateModified)
		if err != nil {
			c.JSON(200, gin.H{"error": "Failed to fetch album dates"})
			return
		}

		c.JSON(200, gin.H{
			"album":         album,
			"date_added":    dateAdded,
			"date_modified": dateModified,
			"user_id":       userID,
		})

	})
	sftpGroup := r.Group("/sftp", AuthRequired())
	{
		sftpGroup.GET("/env", sftpTools.SftpEnvHandler)
		sftpGroup.POST("/install-rclone", sftpTools.InstallRcloneHandler)
		sftpGroup.POST("/start-local", sftpTools.StartLocalSFTPHandler)
		sftpGroup.POST("/stop-local", sftpTools.StopLocalSFTPHandler)
		sftpGroup.GET("/status-local", sftpTools.LocalSFTPStatusHandler)
	}
	r.POST("/song/delete/:song_id", AuthRequired(), func(c *gin.Context) {
		user, err := sftpTools.GetCurrentUser(c)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}

		songID := c.Param("song_id")

		_, err = db.Exec(`DELETE FROM user_music WHERE user_id = $1 AND song_id = $2`, user.ID, songID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete from usermusic"})
			return
		}

		var count int
		err = db.QueryRow(`SELECT COUNT(*) FROM user_music WHERE song_id = $1`, songID).Scan(&count)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to check song usage"})
			return
		}

		if count == 0 {
			_, err = db.Exec(`DELETE FROM songs WHERE song_id = $1`, songID)
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to delete song"})
				return
			}
		}

		c.JSON(200, gin.H{"message": "Song deleted successfully"})
	})

	r.POST("/album/delete/:album_id", AuthRequired(), func(c *gin.Context) {
		user, err := sftpTools.GetCurrentUser(c)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}

		albumID := c.Param("album_id")

		tx, err := db.Begin()
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to start transaction"})
			return
		}
		defer tx.Rollback()

		_, err = tx.Exec(`
		DELETE FROM user_music 
		WHERE user_id = $1 
		  AND song_id IN (SELECT song_id FROM songs WHERE album_id = $2)
	`, user.ID, albumID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete usermusic relations"})
			return
		}

		_, err = tx.Exec(`
		DELETE FROM songs
		WHERE album_id = $1
		  AND song_id NOT IN (SELECT song_id FROM user_music)
	`, albumID)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to delete songs"})
			return
		}

		var remaining int
		err = tx.QueryRow(`SELECT COUNT(*) FROM songs WHERE album_id = $1`, albumID).Scan(&remaining)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to check remaining songs"})
			return
		}

		if remaining == 0 {
			_, err = tx.Exec(`DELETE FROM albums WHERE album_id = $1`, albumID)
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to delete album"})
				return
			}
		}

		if err := tx.Commit(); err != nil {
			c.JSON(500, gin.H{"error": "Failed to commit transaction"})
			return
		}

		c.JSON(200, gin.H{"message": "Album and related songs deleted successfully"})
	})
	r.POST("/song/update-metadata/:song_id", AuthRequired(), tagedit.UpdateSongMetadata)

	r.GET("/upload/status", AuthRequired(), func(c *gin.Context) {
		user, err := sftpTools.GetCurrentUser(c)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}
		// fmt.Println("❌ USER (%s): %v", user, err)

		mu.Lock()
		progress := uploadProgress[user.ID]
		userPending := pendingAlbums[user.ID]
		mu.Unlock()
		// fmt.Println("Current progress for user", user.ID, ":", progress)
		// fmt.Println("Current userPending for user", user.ID, ":", userPending)
		tree := buildTreeFromPending(userPending)
		// fmt.Println("TREE:", tree)

		treeStr := buildTreeString(tree)
		// fmt.Println("TREE:", treeStr)
		allDone := true
		for _, s := range progress {
			if !s.Done {
				allDone = false
				break
			}
		}
		c.JSON(200, gin.H{
			"progress": progress,
			"tree":     treeStr,
			"finished": allDone,
		})
	})

	r.POST("/upload", AuthRequired(), func(c *gin.Context) {
		fmt.Println(">>> Upload handler entered")
		user, err := sftpTools.GetCurrentUser(c)
		if err != nil {
			fmt.Println("❌ sftpTools.GetCurrentUser failed:", err)
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}

		form, err := c.MultipartForm()
		if err != nil {
			fmt.Println("❌ Failed to parse form:", err)
			c.JSON(400, gin.H{"error": "Failed to parse form"})
			return
		}

		files := form.File["files[]"]
		if len(files) == 0 {
			fmt.Println("⚠️ No files found in form")
			c.JSON(400, gin.H{"error": "No files uploaded"})
			return
		}

		mu.Lock()
		uploadProgress[user.ID] = []UploadStatus{}
		for _, file := range files {
			if !strings.HasSuffix(strings.ToLower(file.Filename), ".flac") {
				continue
			}
			uploadProgress[user.ID] = append(uploadProgress[user.ID], UploadStatus{
				Artist:           "Unknown",
				Title:            file.Filename,
				Done:             false,
				OriginalFilename: file.Filename,
			})
		}
		mu.Unlock()

		uploaded := []Song{}
		albumMeta := map[string]*PendingAlbum{}
		albumGenres := map[string][]string{} // albumName -> list of genres
		albumSubgenres := map[string][]string{}

		for _, file := range files {
			if !strings.HasSuffix(strings.ToLower(file.Filename), ".flac") {
				fmt.Println("⚠️ Skipping non-FLAC file:", file.Filename)
				continue
			}

			f, err := file.Open()
			if err != nil {
				fmt.Printf("❌ Failed to open file %s: %v\n", file.Filename, err)
				continue
			}
			data, _ := io.ReadAll(f)
			f.Close()

			title, albumName, albumGenre, artistName, year, albumComment, albumCopyright, tracknumber :=
				extractMetadataFromReader(bytes.NewReader(data), file.Filename)
			coverBase64 := extractAlbumCoverFromReader(bytes.NewReader(data))

			mainGenre, subs := tagedit.NormalizeGenre(albumGenre)
			albumGenres[albumName] = append(albumGenres[albumName], mainGenre)
			albumSubgenres[albumName] = append(albumSubgenres[albumName], subs...)

			if coverBase64 == "" {
				imgFile := "./static/img/default.jpg"
				f, err := os.Open(imgFile)
				if err == nil {
					defer f.Close()
					img, _, _ := image.Decode(f)
					buf := new(bytes.Buffer)
					jpeg.Encode(buf, img, &jpeg.Options{Quality: 80})
					coverBase64 = "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(buf.Bytes())
				}
			}

			mu.Lock()
			for i := range uploadProgress[user.ID] {
				if uploadProgress[user.ID][i].Title == file.Filename {
					uploadProgress[user.ID][i].Artist = artistName
					uploadProgress[user.ID][i].Title = title
					break
				}
			}
			mu.Unlock()

			if _, ok := albumMeta[albumName]; !ok {
				albumMeta[albumName] = &PendingAlbum{
					AlbumName: albumName,
					Artist:    artistName,
					Cover:     coverBase64,
					Year:      year,
					Comment:   albumComment,
					Copyright: albumCopyright,
					SongsPA:   []Song{},
				}
			}

			duration := getFLACDuration(bytes.NewReader(data))
			tn := 0
			if tracknumber != "" {
				if n, err := strconv.Atoi(strings.TrimSpace(tracknumber)); err == nil {
					tn = n
				} else {
					re := regexp.MustCompile(`^(\d+)`)
					if m := re.FindStringSubmatch(tracknumber); len(m) > 1 {
						if n, err := strconv.Atoi(m[1]); err == nil {
							tn = n
						}
					}
				}
			}

			albumMeta[albumName].SongsPA = append(albumMeta[albumName].SongsPA, Song{
				Title:       title,
				Duration:    duration,
				TrackNumber: tn,
				FileName:    file.Filename,
				Data:        data,
			})
		}

		mu.Lock()
		pendingAlbums[user.ID] = albumMeta
		mu.Unlock()

		for albumName, meta := range albumMeta {
			freq := map[string]int{}
			for _, g := range albumGenres[albumName] {
				freq[g]++
			}
			mainGenre := "Unknown"
			max := 0
			for g, count := range freq {
				if count > max {
					mainGenre = g
					max = count
				}
			}

			subSet := map[string]bool{}
			for _, sg := range albumSubgenres[albumName] {
				sg = strings.TrimSpace(sg)
				if sg != "" {
					subSet[sg] = true
				}
			}
			subList := []string{}
			for sg := range subSet {
				if sg != mainGenre {
					subList = append(subList, sg)
				}
			}
			subgenresStr := strings.Join(subList, ", ")

			artistID, err := tagedit.GetOrCreateArtist(meta.Artist)
			if err != nil {
				fmt.Println("❌ Artist insert failed:", err)
				continue
			}

			albumID, err := tagedit.GetOrCreateAlbum(
				meta.AlbumName, artistID, meta.Year, meta.Cover,
				mainGenre, subgenresStr, meta.Comment, meta.Copyright,
			)
			if err != nil {
				fmt.Println("❌ Album insert failed:", err)
				continue
			}

			for _, s := range meta.SongsPA {
				var songID int
				err := db.QueryRow(`
                INSERT INTO songs(title, album_id, duration, track_number)
                VALUES($1,$2,$3,$4)
                ON CONFLICT(title, album_id) DO UPDATE
                    SET duration=EXCLUDED.duration,
                        track_number=EXCLUDED.track_number
                RETURNING song_id
            `, s.Title, albumID, s.Duration, s.TrackNumber).Scan(&songID)
				if err != nil {
					fmt.Println("❌ Song insert failed:", err)
					continue
				}

				creds, err := sftpTools.GetUserSFTPCredsFromDB(user.ID)
				if err != nil {
					fmt.Println("❌ GetUserSFTPCredsFromDB failed:", err)
					continue
				}
				remotePath := path.Join(creds.Path, s.FileName)

				_, err = db.Exec(`
                INSERT INTO user_music(user_id, song_id, file_path)
                VALUES($1,$2,$3)
                ON CONFLICT(user_id, song_id) DO NOTHING
            `, user.ID, songID, remotePath)
				if err != nil {
					fmt.Println("❌ user_music insert failed:", err)
					continue
				}

				os.MkdirAll("./temp_uploads", 0755)

				localTempPath := fmt.Sprintf("./temp_uploads/u%d_%s", user.ID, s.FileName)
				if err := os.WriteFile(localTempPath, s.Data, 0644); err != nil {
					fmt.Printf("❌ Failed to save temp file for SFTP %s: %v\n", s.FileName, err)
				} else {
					fmt.Printf("[sftp] UploadToUserSFTP start user=%d local=%s filename=%s\n",
						user.ID, localTempPath, s.FileName)

					if err := sftpTools.UploadToUserSFTP(user.ID, localTempPath, s.FileName); err != nil {
						fmt.Printf("❌ SFTP upload failed user=%d file=%s err=%v\n",
							user.ID, s.FileName, err)
					} else {
						fmt.Printf("[sftp] ✅ UploadToUserSFTP done user=%d local=%s\n",
							user.ID, localTempPath)
					}

					if err := os.Remove(localTempPath); err != nil {
						fmt.Printf("⚠️ Failed to remove temp file %s: %v\n", localTempPath, err)
					}
				}

				mu.Lock()
				for i := range uploadProgress[user.ID] {
					if uploadProgress[user.ID][i].OriginalFilename == s.FileName {
						uploadProgress[user.ID][i].Done = true
						break
					}
				}
				mu.Unlock()
			}
		}

		c.JSON(200, gin.H{
			"message":  "Upload finished",
			"uploaded": uploaded,
		})
		uploaded = nil
	})

	r.NoRoute(func(c *gin.Context) {
		if c.Request.Method != http.MethodGet {
			c.JSON(404, gin.H{"error": "not found"})
			return
		}
		if _, err := os.Stat(indexFile); err != nil {
			c.JSON(500, gin.H{"error": "frontend missing"})
			return
		}
		c.File(indexFile)
	})

	apiAddr := fmt.Sprintf(":%d", portsConfig.APIPort)
	log.Printf("API listening at http://localhost:%d", portsConfig.APIPort)

	r.Run(apiAddr)
}

func startFrontendServer(distDir, indexFile string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/config/ports", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(portsConfig)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestPath := r.URL.Path
		clean := filepath.Clean(requestPath)
		target := filepath.Join(distDir, clean)
		if info, err := os.Stat(target); err == nil && !info.IsDir() {
			http.ServeFile(w, r, target)
			return
		}
		http.ServeFile(w, r, indexFile)
	})

	addr := fmt.Sprintf(":%d", portsConfig.FrontendPort)
	log.Printf("Serving frontend on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil && err != http.ErrServerClosed {
		log.Printf("frontend server error: %v", err)
	}
}

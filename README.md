# MusicApp Multiuser Server

This repository contains the Go/Gin backend prepared for multiuser deployment (separate from the main monorepo). It is a clean copy of the `Gin/` service to make hardening, scaling, and security work easier.

## Layout
- `Gin/` — Go server source copied from the main project.

## Next steps (recommended)
1. Add multiuser hardening
   - Replace SHA256 password hashing with bcrypt (golang.org/x/crypto/bcrypt).
   - Wrap all access to shared maps (`uploadProgress`, `pendingAlbums`) with mutex; consider moving state to Postgres for multi-instance scaling.
   - Set DB pool tuning (`SetMaxOpenConns`, `SetMaxIdleConns`, `SetConnMaxLifetime`).
   - Add request size limits for uploads (`gin.Engine.MaxMultipartMemory`) and validate file names/extensions.
2. Security checks
   - Run `govulncheck ./...` and `gosec ./...`.
   - Lock CORS to prod origins; avoid `*` in production.
   - Ensure auth middleware on all mutating routes and avoid logging secrets.
3. Performance testing
   - Use `wrk` or `k6` to capture p50/p95/p99 latency, RPS, and error rates on key endpoints.
   - Enable Postgres slow-query logging for profiling hot paths.
4. Deployment
   - Build the binary: `cd Gin && go build -o bin/musicapp`.
   - Run under systemd (example service in prior instructions) behind nginx with TLS.

## Quick start
```
cd Gin
go mod tidy
go test ./...
go build -o bin/musicapp
./bin/musicapp
```

Ensure `DATABASE_URL`, `GIN_MODE=release`, `PORT`, and other secrets are set in your environment or a `.env` loaded by the app.

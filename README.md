# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Run the server
go build -o out/chirpy && ./out/chirpy

# Or run directly
go run .

# Run tests
go test ./...

# Run a single test
go test ./... -run TestName

# Generate database code from SQL (requires sqlc)
sqlc generate

# Apply database migrations (requires goose)
goose -dir sql/schema postgres "postgres://<user>:@localhost:5432/chirpy" up
```

## Environment

Requires a `.env` file (or environment variables):

```
DB_URL=postgres://user:password@host:port/chirpy
PLATFORM=dev   # or "prod"; "dev" enables the /admin/reset endpoint
```

Connect directly to the database:

```bash
psql "postgres://<user>:@localhost:5432/chirpy"
```

## Architecture

This is a Go HTTP API server (`package main`, flat structure) backed by PostgreSQL.

**Request flow:** `main.go` sets up `apiConfig` (holds DB connection + platform + hit counter) and registers all routes on a plain `net/http.ServeMux`. Handlers are either standalone functions or methods on `*apiConfig` when they need DB/config access.

**Database layer:** SQL queries live in `sql/queries/`, schema migrations in `sql/schema/` (goose format). `sqlc` generates type-safe Go code into `internal/database/` — never edit files in that directory directly; edit the `.sql` source files and re-run `sqlc generate`.

**Response helpers:** `json.go` provides `respondWithJSON` and `respondWithError` used by all handlers.

**Routes:**

- `GET /api/healthz` — readiness check
- `POST /api/users` — create user
- `POST /api/chirps` — validate and clean chirp body (profanity filter, 140-char limit)
- `GET /admin/metrics` — fileserver hit count (HTML)
- `POST /admin/reset` — reset hit counter + wipe DB (dev only)
- `/app/` — static file server

# Chirpy

A simple microblogging REST API built with Go and PostgreSQL. Users can post short "chirps" (up to 140 characters) with automatic profanity filtering.

## Tech Stack

- **Go** — HTTP server using `net/http`
- **PostgreSQL** — primary data store
- **sqlc** — generates type-safe Go code from SQL queries
- **goose** — database migrations

## Prerequisites

- Go 1.21+
- PostgreSQL
- [sqlc](https://sqlc.dev) (for regenerating database code)
- [goose](https://github.com/pressly/goose) (for running migrations)

## Setup

1. Clone the repo and create a `.env` file:

```
DB_URL=postgres://user:password@localhost:5432/chirpy
PLATFORM=dev
```

> Set `PLATFORM=prod` to disable the `/admin/reset` endpoint.

2. Apply database migrations:

```bash
goose -dir sql/schema postgres "postgres://user:password@localhost:5432/chirpy" up
```

## Running

```bash
go run .
```

Or build first:

```bash
go build -o out/chirpy && ./out/chirpy
```

## API

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/api/healthz` | Health check |
| `POST` | `/api/users` | Create a user |
| `POST` | `/api/chirps` | Post a chirp (max 140 chars) |
| `GET` | `/api/chirps` | Get all chirps |
| `GET` | `/api/chirps/{chirpID}` | Get a single chirp |
| `GET` | `/admin/metrics` | View request hit count |
| `POST` | `/admin/reset` | Reset metrics and wipe DB (dev only) |

## Tests

```bash
go test ./...
```

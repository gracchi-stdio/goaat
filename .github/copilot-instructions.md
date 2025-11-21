# Goaat Project - AI Coding Agent Instructions

## Architecture Overview
This is a Go project enabling the users to edit markdowns with frontmatter service using **Echo v4** framework with **Templ** templates and **PostgreSQL** via **pgx/v5**. The project uses **SQLC** for type-safe database queries and runs in **Podman 5** containers.

**Module**: `github.com/gracchi-stdio/goaat`  
**Entry Point**: `cmd/server/main.go` (for now I have only one server serving web with Templ, API server planned later)

## Container Environment (Podman 5)

### Critical: Database Connection
- Podman 5 has working DNS, so use container hostname: `goaat-db:5432`
- Database URL: `postgres://user:password@goaat-db:5432/goaatDB?sslmode=disable`
- Always set 5s timeout on `pgx.Connect()` - fails gracefully if DB unavailable
- DB connection pattern in `main.go` lines 60-73: timeout context, warning on failure, continues without DB


## Web Framework (Echo v4)

### Custom Colorful Logger
Uses ANSI colors in request logs (lines 21-50 in `main.go`):
- **Cyan** methods
- **Green** 2xx status
- **Yellow** 4xx status  
- **Red** 5xx status
- Format: `METHOD | STATUS | LATENCY | IP | URI`

### Middleware Order
1. Custom RequestLogger (with colors)
2. Recover (panic recovery)

### Environment Variables
- `DATABASE_URL` - PostgreSQL connection string
- `PORT` - Server port (default: 8080)

## Project Structure (Current)
```
cmd/server/main.go          # Entry point
internal/
  db/                       # SQLC generated code
  web/                      # Web layer (in progress)
    handlers/               # HTTP handlers
    templates/              # Templ files (*.templ)
    routes.go              # Route definitions
compose.yaml                # Podman compose setup
schema.sql                  # DB schema
query.sql                   # SQLC queries
sqlc.yaml                   # SQLC config
```

## Planned Migration
Moving to cleaner structure with:
- `cmd/web/main.go` - Web server with Templ
- `internal/web/` - Web handlers and templates
- Future: `cmd/api/` for REST API
- `internal/shared/` for common middleware

## Key Dependencies
- **echo/v4**: Web framework
- **pgx/v5**: PostgreSQL driver (direct, not database/sql)
- **templ**: Go templates (install: `go install github.com/a-h/templ/cmd/templ@latest`)
- **air**: Hot reload (in container only)

## Common Patterns

### Error Handling
- DB connection errors: log warning, continue without DB
- Use Echo's logger: `e.Logger.Info()`, `e.Logger.Warn()`
- HTTP errors: return `c.JSON(status, map[string]string{"error": msg})`

### Context Usage
- Always use `context.WithTimeout()` for DB operations (5s timeout)
- Defer cancel: `defer cancel()`
- Don't reuse request context for cleanup - use `context.Background()`

## Testing
- Access API: `curl http://localhost:8080/`
- Check DB: DBeaver to `localhost:5432`, user: `user`, password: `password`, db: `goaatDB`

## Important Notes
- Podman 5 required (DNS doesn't work in Podman 4)
- Use `podman-compose` not `podman compose` (installed via `uv tool install podman-compose`)
- Module name is `goaat` but folder is still `goauth` (will rename)
- Empty dirs exist for future features: `config/`, `services/`, `middleware/`

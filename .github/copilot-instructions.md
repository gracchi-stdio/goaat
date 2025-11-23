# Goaat Project - AI Coding Agent Instructions

## Architecture Overview
This is a Go project enabling the users to edit markdowns with frontmatter service using **Echo v4** framework with **Templ** templates and **PostgreSQL** via **pgx/v5**. The project uses **SQLC** for type-safe database queries and runs in **Podman 5** containers.

**Module**: `github.com/gracchi-stdio/goaat`  
**Entry Point**: `cmd/server/main.go`

## Container Environment (Podman 5)

### Critical: Database Connection
- **App to DB**: Uses `host.containers.internal` in `DATABASE_URL` for Podman networking.
- **Migrations**: Run via `tern` inside the app container.
- **Command**: `podman exec -it goaat-app yarn migrate` (uses `host.containers.internal`).

## Web Framework (Echo v4)

### Authentication (OAuth)
- **Service Pattern**: `internal/auth/service.go` defines the interface.
- **Implementation**: `internal/auth/auth.go` uses **Goth** (GitHub provider).
- **Session**: **Gorilla Sessions** (Cookie-based) via `echo-contrib`.
- **User Storage**: Users are upserted into `users` table on login.

### Project Structure
```
cmd/server/main.go          # Entry point (wires Service -> Handlers)
internal/
  auth/                     # Auth Service & Goth implementation
  platform/
    db/                     # SQLC generated code
    logger/                 # Custom colorful logger
  web/
    handlers/               # HTTP handlers (injected with AuthService)
    templates/              # Templ files (*.templ)
    routes.go               # Route definitions
db/
  migrations/               # SQL migrations (Tern)
  queries/                  # SQLC queries
compose.yaml                # Podman compose setup
tern.conf                   # Migration config
```

## Key Workflows

### Database Migrations
- **Tool**: `tern` (installed in dev container).
- **Run**: `podman exec -it goaat-app yarn migrate`
- **Create New**: `podman exec -it goaat-app yarn migrate:new -n name_of_migration`

### CSS & UI
- **Framework**: **Shoelace** (Web Components) + Custom CSS.
- **Structure**:
  - `assets/css/main.css`: Entry point.
  - `assets/css/pages/`: Page-specific styles.
  - `variables.css`: Global design tokens.
- **Templ Layout**: Accepts `class` param for body styling (e.g., `@Layout("Title", "page-login")`).

### Development
- **Hot Reload**: `air` runs automatically in the container.
- **Templ Generation**: `air` triggers `templ generate` on file save.

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
- explain the plan before writing code after confirming with me

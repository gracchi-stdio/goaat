# Goaat Project - AI Coding Agent Instructions

## Project Vision
A Go backend for editing **Astro Starlight** content repositories (markdown files with frontmatter). The system authenticates users via GitHub OAuth, clones their repos, and provides an editing interface for documentation sites.

## Architecture Overview

**Stack**: Echo v4 + Templ + PostgreSQL (pgx/v5) + SQLC  
**Module**: `github.com/gracchi-stdio/goaat`  
**Entry Point**: `cmd/server/main.go`  
**Container Runtime**: Podman 5

### Project Structure
```
cmd/
  server/main.go           # Entry point, wiring, graceful shutdown
internal/
  auth/                    # OAuth service (interface + Goth impl)
    service.go             # Interface definition
    auth.go                # Goth/GitHub implementation
  config/                  # Environment config loader
  middleware/              # Echo middleware (auth, logging)
  platform/                # Infrastructure adapters
    db/                    # SQLC generated code (DO NOT EDIT)
    logger/                # Custom Echo logger
  web/                     # HTTP layer
    handlers/              # Request handlers (thin, delegate to services)
    templates/             # Templ components
      layouts/             # Base layouts (Layout, AuthedLayout)
      pages/               # Page components
      components/          # Reusable UI components
    routes.go              # Route registration
db/
  migrations/              # Tern SQL migrations
  queries/                 # SQLC query definitions
  sqlc.yaml                # SQLC config
```

### Future Structure (Content Editing)
```
internal/
  content/                 # Content service (planned)
    service.go             # Interface: ParseMarkdown, ValidateFrontmatter
    markdown.go            # goldmark implementation
  repository/              # Git operations (planned)
    service.go             # Interface: Clone, Pull, Push, Commit
    git.go                 # go-git implementation
  github/                  # GitHub API client (planned)
    client.go              # Repo listing, PR creation
```

## Code Style & Patterns

### Dependency Injection
- Use constructor functions: `New(deps) *Service`
- Handlers receive services via struct fields
- Interfaces defined where consumed, not where implemented

```go
// Good: Interface in consumer package
type Handler struct {
    Auth auth.Service  // Interface from auth package
    DB   *db.Queries
}
```

### Error Handling
- Wrap errors with context: `fmt.Errorf("failed to fetch user: %w", err)`
- Use `echo.NewHTTPError(status, message)` for HTTP errors
- Check nil before DB operations: `if h.DB == nil { return error }`
- Log warnings for non-fatal errors, continue gracefully

### Context Usage
- Always use `context.WithTimeout()` for external calls (5s default)
- Defer cancel immediately: `defer cancel()`
- Use `c.Request().Context()` in handlers
- Use `context.Background()` for cleanup operations

### Handler Pattern
```go
func (h *Handler) GetThing(c echo.Context) error {
    if h.DB == nil {
        return echo.NewHTTPError(http.StatusServiceUnavailable, "database unavailable")
    }
    
    ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
    defer cancel()
    
    thing, err := h.DB.GetThing(ctx, id)
    if err != nil {
        return echo.NewHTTPError(http.StatusNotFound, "thing not found")
    }
    
    return Render(c, pages.ThingPage(thing))
}
```

### Templ Rendering Helper
```go
// Add to handlers/handler.go
func Render(c echo.Context, component templ.Component) error {
    return component.Render(c.Request().Context(), c.Response().Writer)
}
```

## Container Environment (Podman 5)

### Database Connection
- **App → DB**: Use `host.containers.internal` in `DATABASE_URL`
- **Local → DB**: Use `localhost:5432`

### Commands
```bash
# Start services
podman-compose up -d

# Run migrations
podman exec -it goaat-app yarn migrate

# Create new migration
podman exec -it goaat-app yarn migrate:new -n migration_name

# View logs
podman-compose logs -f app
```

### Important
- Podman 5 required (DNS broken in v4)
- Use `podman-compose` not `podman compose`

## Development Workflow

### Hot Reload
- `air` runs in container, watches `.go` and `.templ` files
- Templ generates on save automatically

### CSS & UI
- **Framework**: Shoelace (Web Components)
- **Entry**: `assets/css/main.css`
- **Pages**: `assets/css/pages/*.css`
- **Tokens**: `variables.css`
- **Templ**: `@Layout("Title", "page-classname")`

### Testing
```bash
curl http://localhost:8080/health
curl http://localhost:8080/api/authors
```

## Database

### SQLC Workflow
1. Write query in `db/queries/*.sql`
2. Run `sqlc generate` (via `yarn sqlc`)
3. Use generated methods in handlers

### Migration Naming
```
001_initial_schema.sql
002_create_users_table.sql
003_add_repositories_table.sql  # Future
```

## Authentication Flow

1. User clicks "Login with GitHub" → `/auth/github`
2. Goth redirects to GitHub OAuth
3. Callback → `/auth/github/callback`
4. User upserted to `users` table
5. Session created with `user_id`, `email`, `name`
6. Redirect to `return_to` URL or `/`

### Session Keys
- `user_id`: Database user ID
- `email`: User email
- `name`: Display name
- `avatar_url`: GitHub avatar
- `return_to`: Redirect after login

## Key Dependencies
| Package | Purpose |
|---------|---------|
| `echo/v4` | Web framework |
| `pgx/v5` | PostgreSQL driver (direct, not database/sql) |
| `templ` | Type-safe HTML templates |
| `goth` | OAuth providers |
| `gorilla/sessions` | Cookie sessions |
| `sqlc` | SQL → Go code generator |
| `tern` | Database migrations |

## Coding Guidelines

### DO
- Explain plan before writing code, confirm with user
- Use interfaces for testability
- Handle nil DB gracefully
- Add context to errors
- Keep handlers thin, logic in services
- Use meaningful HTTP status codes

### DON'T
- Edit SQLC generated files (`internal/platform/db/*.go`)
- Use `fmt.Printf` for logging (use Echo logger)
- Hardcode configuration values
- Ignore context cancellation
- Put business logic in handlers

## Planned Features (Astro Starlight Backend)
1. **Repository Management**: Clone/sync user's GitHub repos
2. **Content Parsing**: Parse markdown frontmatter (title, description, sidebar)
3. **File Browser**: Navigate repo structure, edit files
4. **Live Preview**: Render markdown as Starlight would
5. **Git Operations**: Commit, push, create PRs
6. **Collaboration**: Multiple editors, review workflow

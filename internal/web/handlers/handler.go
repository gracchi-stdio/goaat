package handlers

import (
	"github.com/a-h/templ"
	"github.com/gracchi-stdio/goaat/internal/auth"
	"github.com/gracchi-stdio/goaat/internal/platform/db"
	"github.com/labstack/echo/v4"
	"github.com/starfederation/datastar-go/datastar"
)

// Handler holds dependencies for HTTP handlers.
// All handlers should check for nil DB before database operations.
type Handler struct {
	DB          *db.Queries
	AuthService auth.Service
}

// New creates a new Handler with dependencies.
// DB can be nil if database is unavailable.
func New(db *db.Queries, authService auth.Service) *Handler {
	return &Handler{
		DB:          db,
		AuthService: authService,
	}
}

// Render is a helper to render templ components with proper context.
func Render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response().Writer)
}

// SSE creates a new Datastar Server-Sent Event generator for streaming responses.
// Use this for AJAX-style updates via Datastar's @get(), @post(), etc. actions.
func SSE(c echo.Context) *datastar.ServerSentEventGenerator {
	return datastar.NewSSE(c.Response(), c.Request())
}

// SSEWithContext creates an SSE generator with a custom context.
// Useful for passing values to templ components during SSE rendering.
func SSEWithContext(c echo.Context) *datastar.ServerSentEventGenerator {
	return datastar.NewSSE(c.Response(), c.Request(), datastar.WithContext(c.Request().Context()))
}

package handlers

import (
	"github.com/a-h/templ"
	"github.com/gracchi-stdio/goaat/internal/auth"
	"github.com/gracchi-stdio/goaat/internal/platform/db"
	"github.com/gracchi-stdio/goaat/internal/web/templates/layouts"
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

// RenderWithDatastar renders a templ component and handles Datastar partial updates
func RenderWithDatastar(c echo.Context, component templ.Component) error {
	// Check if this is a Datastar request
	if c.Request().Header.Get("datastar-request") != "" {
		sse := datastar.NewSSE(c.Response().Writer, c.Request())
		// Wrap the content in the PageContent layout to preserve the <main> tag and ID
		return sse.PatchElementTempl(layouts.PageContentWrapper(component),
			datastar.WithSelectorID("page-content"),
			datastar.WithViewTransitions(),
		)
	}

	// Regular full page render
	return component.Render(c.Request().Context(), c.Response().Writer)
}

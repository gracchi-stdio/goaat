package handlers

import (
	"github.com/gracchi-stdio/goaat/internal/web/templates/pages"
	"github.com/labstack/echo/v4"
)

// DashboardPage renders the main dashboard page with Datastar support
func (h *Handler) DashboardPage(c echo.Context) error {
	// Example: Show a welcome alert on Datastar navigation
	// This will be picked up by the frontend JS
	if c.Request().Header.Get("datastar-request") != "" {
		return RenderWithDatastar(c, pages.DashboardContent())
	}
	return Render(c, pages.Dashboard())
}

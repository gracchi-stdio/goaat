package handlers

import (
	"github.com/gracchi-stdio/goaat/internal/web/templates/pages"
	"github.com/labstack/echo/v4"
)

// SettingsPage renders the settings page
func (h *Handler) SettingsPage(c echo.Context) error {
	if c.Request().Header.Get("datastar-request") != "" {
		return RenderWithDatastar(c, pages.SettingsContent())
	}
	return Render(c, pages.Settings())
}

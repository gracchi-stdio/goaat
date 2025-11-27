package handlers

import (
	"github.com/gracchi-stdio/goaat/internal/web/templates/pages"
	"github.com/labstack/echo/v4"
)

// RepositoriesPage renders the repositories page
func (h *Handler) RepositoriesPage(c echo.Context) error {
	if c.Request().Header.Get("datastar-request") != "" {
		return RenderWithDatastar(c, pages.RepositoriesContent())
	}
	return Render(c, pages.Repositories())
}

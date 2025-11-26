package handlers

import (
	"github.com/gracchi-stdio/goaat/internal/web/templates/pages"
	"github.com/labstack/echo/v4"
)

// LoginPage renders the login page with OAuth options.
func (h *Handler) LoginPage(c echo.Context) error {
	return Render(c, pages.Login())
}

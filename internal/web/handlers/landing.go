package handlers

import (
	"github.com/gracchi-stdio/goaat/internal/web/templates/pages"
	"github.com/labstack/echo/v4"
)

// LandingPage renders the main landing page.
func (h *Handler) LandingPage(c echo.Context) error {
	return Render(c, pages.Landing())
}

package handlers

import (
	"github.com/gracchi-stdio/goaat/internal/web/templates/pages"
	"github.com/labstack/echo/v4"
)

// HelloPage renders the hello page with optional name parameter.
func (h *Handler) HelloPage(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		name = "World"
	}
	return Render(c, pages.Hello(name))
}

// AuthorListPage renders the author list page with Datastar support
func (h *Handler) AuthorListPage(c echo.Context) error {
	// Mock data for now - will use real service later
	authors := []string{"Alice", "Bob", "Charlie"}
	if c.Request().Header.Get("datastar-request") != "" {
		return RenderWithDatastar(c, pages.AuthorListContent(authors))
	}
	return Render(c, pages.AuthorList(authors))
}

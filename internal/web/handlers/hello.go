package handlers

import (
	"github.com/gracchi-stdio/goaat/internal/web/templates/pages"
	"github.com/labstack/echo/v4"
)

func (h *Handler) HelloPage(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		name = "World"
	}

	component := pages.Hello(name)
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func (h *Handler) AuthorListPage(c echo.Context) error {
	// Mock data for now - will use real service later
	authors := []string{"Alice", "Bob", "Charlie"}

	component := pages.AuthorList(authors)
	return component.Render(c.Request().Context(), c.Response().Writer)
}

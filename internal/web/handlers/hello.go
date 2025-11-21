package handlers

import (
	"github.com/gracchi-stdio/goaat/internal/services"
	"github.com/gracchi-stdio/goaat/internal/web/templates"
	"github.com/labstack/echo/v4"
)

func HelloPage(svc *services.Services, c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		name = "World"
	}

	component := templates.Hello(name)
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func AuthorListPage(svc *services.Services, c echo.Context) error {
	// Mock data for now - will use real service later
	authors := []string{"Alice", "Bob", "Charlie"}

	component := templates.AuthorList(authors)
	return component.Render(c.Request().Context(), c.Response().Writer)
}

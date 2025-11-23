package handlers

import (
	"github.com/gracchi-stdio/goaat/internal/web/templates"
	"github.com/labstack/echo/v4"
)

func (h *Handler) LoginPage(c echo.Context) error {
	component := templates.Login()
	return component.Render(c.Request().Context(), c.Response().Writer)
}

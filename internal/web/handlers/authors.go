package handlers

import (
	"net/http"

	"github.com/gracchi-stdio/goaat/internal/services"
	"github.com/labstack/echo/v4"
)

func ListAuthors(svc *services.Services, c echo.Context) error {
	authors, err := svc.Author.List(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, authors)
}

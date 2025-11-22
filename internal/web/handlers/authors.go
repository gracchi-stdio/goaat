package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) ListAuthors(c echo.Context) error {
	authors, err := h.DB.ListAuthors(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, authors)
}

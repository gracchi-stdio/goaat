package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) ListAuthors(c echo.Context) error {
	if h.DB == nil {
		return echo.NewHTTPError(http.StatusServiceUnavailable, "database unavailable")
	}

	authors, err := h.DB.ListAuthors(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch authors")
	}
	return c.JSON(http.StatusOK, authors)
}

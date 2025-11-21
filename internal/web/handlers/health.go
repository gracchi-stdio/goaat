package handlers

import (
	"net/http"

	"github.com/gracchi-stdio/goaat/internal/services"
	"github.com/labstack/echo/v4"
)

// Health handles health check requests
func Health(svc *services.Services, c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "healthy",
	})
}

package web

import (
	"github.com/gracchi-stdio/goaat/internal/services"
	"github.com/gracchi-stdio/goaat/internal/web/handlers"
	"github.com/labstack/echo/v4"
)

// RegisterRoutes sets up all application routes
func RegisterRoutes(svc *services.Services, e *echo.Echo) {
	// Health checks
	e.GET("/", With(svc, handlers.Health))
	e.GET("/health", With(svc, handlers.Health))

	// Templ pages
	e.GET("/hello", With(svc, handlers.HelloPage))
	e.GET("/authors", With(svc, handlers.AuthorListPage))
}

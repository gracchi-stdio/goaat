package web

import (
	"github.com/gracchi-stdio/goaat/internal/auth"
	"github.com/gracchi-stdio/goaat/internal/middleware"
	"github.com/gracchi-stdio/goaat/internal/platform/db"
	"github.com/gracchi-stdio/goaat/internal/web/handlers"
	"github.com/labstack/echo/v4"
)

// RegisterRoutes sets up all application routes
func RegisterRoutes(e *echo.Echo, queries *db.Queries, authService auth.Service) {
	// Initialize handlers with dependencies
	h := handlers.New(queries, authService)

	// Health checks
	e.GET("/", h.Health)
	e.GET("/health", h.Health)

	// Templ pages
	e.GET("/hello", h.HelloPage)
	e.GET("/login", h.LoginPage)

	// Auth routes
	e.GET("/auth/:provider", h.Auth)
	e.GET("/auth/:provider/callback", h.AuthCallback)
	e.GET("/logout/:provider", h.Logout)
	e.GET("/logout", h.Logout)

	// Authenticated routes
	authGroup := e.Group("/admin")
	authGroup.Use(middleware.RequireAuth)
	authGroup.GET("/authors", h.AuthorListPage)

	// API
	e.GET("/api/authors", h.ListAuthors)

}

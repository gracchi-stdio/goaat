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

	// Public pages with user context
	publicPages := e.Group("")
	publicPages.Use(middleware.InjectUser)
	publicPages.GET("/", h.LandingPage)
	publicPages.GET("/hello", h.HelloPage)
	publicPages.GET("/login", h.LoginPage)

	// Health checks (no user needed)
	e.GET("/health", h.Health)

	// Auth routes
	e.GET("/auth/:provider", h.Auth)
	e.GET("/auth/:provider/callback", h.AuthCallback)
	e.POST("/logout/:provider", h.Logout)
	e.POST("/logout", h.Logout)

	// Authenticated routes
	authGroup := e.Group("/admin")
	authGroup.Use(middleware.RequireAuth)
	authGroup.GET("/dashboard", h.DashboardPage)
	authGroup.GET("/authors", h.AuthorListPage)
	authGroup.GET("/profile", h.ProfilePage)
	authGroup.POST("/profile/update", h.UpdateProfile)
	authGroup.GET("/repositories", h.RepositoriesPage)
	authGroup.GET("/settings", h.SettingsPage)

	// API
	e.GET("/api/authors", h.ListAuthors)

}

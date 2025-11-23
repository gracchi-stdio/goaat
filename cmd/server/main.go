package main

import (
	"context"
	"os"
	"time"

	"github.com/gorilla/sessions"
	"github.com/gracchi-stdio/goaat/internal/auth"
	"github.com/gracchi-stdio/goaat/internal/config"
	"github.com/gracchi-stdio/goaat/internal/middleware"
	"github.com/gracchi-stdio/goaat/internal/platform/db"
	"github.com/gracchi-stdio/goaat/internal/platform/logger"
	"github.com/gracchi-stdio/goaat/internal/web"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize Auth
	if err := auth.Init(cfg); err != nil {
		// Log warning but don't fail if auth is not configured in dev
		if cfg.Environment == "production" {
			panic(err)
		}
		// In dev, we might proceed without auth or log it
		// For now, let's just print it
		// fmt.Println("Auth init failed:", err)
	}

	// Initialize Echo
	e := echo.New()
	e.Logger = logger.NewColorful()
	e.Logger.SetOutput(logger.NewColorWriter(os.Stdout))

	// Middleware
	e.Use(middleware.ColorfulLogger())
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.Static("public"))

	// Session Middleware
	if cfg.SessionSecret != "" {
		e.Use(session.Middleware(sessions.NewCookieStore([]byte(cfg.SessionSecret))))
	} else {
		e.Logger.Warn("SESSION_SECRET not set, session middleware disabled")
	}

	// Database connection
	queries := initDatabase(e, cfg.DatabaseURL)

	// Initialize Auth Service
	authService := auth.NewService()

	// Routes
	web.RegisterRoutes(e, queries, authService)

	// Start server
	e.Logger.Fatal(e.Start(":" + cfg.Port))
}

func initDatabase(e *echo.Echo, dbURL string) *db.Queries {
	if dbURL == "" {
		e.Logger.Warn("DATABASE_URL not set, skipping database connection")
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	e.Logger.Info("Connecting to database...")
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		e.Logger.Warn("Database connection failed:", err)
		e.Logger.Info("Starting server without database connection")
		return nil
	}

	// Ping to verify connection
	if err := pool.Ping(ctx); err != nil {
		e.Logger.Warn("Database ping failed:", err)
		return nil
	}

	e.Logger.Info("Successfully connected to database")
	return db.New(pool)
}

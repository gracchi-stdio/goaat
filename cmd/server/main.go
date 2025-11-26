package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	pool, queries := initDatabase(e, cfg.DatabaseURL)

	// Initialize Auth Service
	authService := auth.NewService()

	// Routes
	web.RegisterRoutes(e, queries, authService)

	// Graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Start server
	go func() {
		if err := e.Start(":" + cfg.Port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal
	<-ctx.Done()
	e.Logger.Info("Shutting down gracefully...")

	// Shutdown with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(shutdownCtx); err != nil {
		e.Logger.Fatal(err)
	}

	// Close database pool
	if pool != nil {
		pool.Close()
		e.Logger.Info("Database connection closed")
	}

	e.Logger.Info("Server stopped")
}

func initDatabase(e *echo.Echo, dbURL string) (*pgxpool.Pool, *db.Queries) {
	if dbURL == "" {
		e.Logger.Warn("DATABASE_URL not set, skipping database connection")
		return nil, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	e.Logger.Info("Connecting to database...")
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		e.Logger.Warn("Database connection failed:", err)
		e.Logger.Info("Starting server without database connection")
		return nil, nil
	}

	// Ping to verify connection
	if err := pool.Ping(ctx); err != nil {
		e.Logger.Warn("Database ping failed:", err)
		return nil, nil
	}

	e.Logger.Info("Successfully connected to database")
	return pool, db.New(pool)
}

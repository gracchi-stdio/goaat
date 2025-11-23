package handlers

import (
	"github.com/gracchi-stdio/goaat/internal/auth"
	"github.com/gracchi-stdio/goaat/internal/platform/db"
)

// Handler holds dependencies for HTTP handlers
type Handler struct {
	DB          *db.Queries
	AuthService auth.Service
}

// New creates a new Handler with dependencies
func New(db *db.Queries, authService auth.Service) *Handler {
	return &Handler{
		DB:          db,
		AuthService: authService,
	}
}

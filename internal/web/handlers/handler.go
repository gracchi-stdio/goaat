package handlers

import "github.com/gracchi-stdio/goaat/internal/db"

// Handler holds dependencies for HTTP handlers
type Handler struct {
	DB *db.Queries
}

// New creates a new Handler with dependencies
func New(db *db.Queries) *Handler {
	return &Handler{
		DB: db,
	}
}

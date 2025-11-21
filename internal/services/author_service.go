package services

import (
	"context"

	"github.com/gracchi-stdio/goaat/internal/db"
)

type AuthorService struct {
	queries *db.Queries
}

func NewAuthorService(queries *db.Queries) *AuthorService {
	return &AuthorService{queries: queries}
}

func (s *AuthorService) List(ctx context.Context) ([]db.Author, error) {
	return s.queries.ListAuthors(ctx)
}

// func (s *AuthorService) GetByID(ctx context.Context, id int64) (db.Author, error) {
//     return s.queries.GetAuthor(ctx, id)
// }

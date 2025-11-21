package services

import "github.com/gracchi-stdio/goaat/internal/db"

type Services struct {
	Author *AuthorService
	// Image   *ImageService
	// Post    *PostService
}

func NewServices(queries *db.Queries) *Services {
	return &Services{
		Author: NewAuthorService(queries),
	}
}

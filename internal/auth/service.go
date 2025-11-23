package auth

import (
	"net/http"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

// Service defines the authentication flow interface.
// It abstracts the underlying OAuth implementation (Goth).
type Service interface {
	// BeginAuth starts the authentication process (redirects to provider)
	BeginAuth(w http.ResponseWriter, r *http.Request)

	// CompleteAuth finishes the authentication process and returns the user profile
	CompleteAuth(w http.ResponseWriter, r *http.Request) (goth.User, error)

	// Logout clears the session
	Logout(w http.ResponseWriter, r *http.Request) error
}

type service struct{}

// NewService creates a new instance of the authentication service
func NewService() Service {
	return &service{}
}

func (s *service) BeginAuth(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
}

func (s *service) CompleteAuth(w http.ResponseWriter, r *http.Request) (goth.User, error) {
	return gothic.CompleteUserAuth(w, r)
}

func (s *service) Logout(w http.ResponseWriter, r *http.Request) error {
	return gothic.Logout(w, r)
}

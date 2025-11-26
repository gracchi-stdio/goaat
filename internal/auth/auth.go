package auth

import (
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/gracchi-stdio/goaat/internal/config"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
)

// Init initializes the authentication providers.
// Returns an error if required environment variables are missing.
func Init(cfg *config.Config) error {
	if cfg.GithubClientID == "" || cfg.GithubClientSecret == "" {
		return fmt.Errorf("GITHUB_CLIENT_ID and GITHUB_CLIENT_SECRET must be set")
	}

	if cfg.SessionSecret == "" {
		return fmt.Errorf("SESSION_SECRET must be set")
	}

	// Set up the store for gothic (handles OAuth state)
	store := sessions.NewCookieStore([]byte(cfg.SessionSecret))
	gothic.Store = store

	callbackURL := cfg.BaseURL + "/auth/github/callback"

	goth.UseProviders(
		github.New(cfg.GithubClientID, cfg.GithubClientSecret, callbackURL),
	)

	return nil
}

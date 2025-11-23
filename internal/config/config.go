package config

import "os"

// Config holds application configuration
type Config struct {
	DatabaseURL        string
	Port               string
	Environment        string // "development" or "production"
	GithubClientID     string
	GithubClientSecret string
	SessionSecret      string
	LoginRedirectURL   string
}

// Load reads configuration from environment variables
func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "production"
	}

	return &Config{
		DatabaseURL:        os.Getenv("DATABASE_URL"),
		Port:               port,
		Environment:        env,
		GithubClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		GithubClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		SessionSecret:      os.Getenv("SESSION_SECRET"),
		LoginRedirectURL:   "/admin",
	}
}

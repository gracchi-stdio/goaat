package config

import (
	"fmt"
	"os"
)

// Config holds application configuration loaded from environment variables.
type Config struct {
	DatabaseURL        string
	Port               string
	Environment        string // "development" or "production"
	BaseURL            string // Base URL for callbacks (e.g., "http://localhost:8080")
	GithubClientID     string
	GithubClientSecret string
	SessionSecret      string
}

// Load reads configuration from environment variables with sensible defaults.
func Load() *Config {
	port := getEnvOrDefault("PORT", "8080")
	env := getEnvOrDefault("ENV", "development")
	baseURL := getEnvOrDefault("BASE_URL", fmt.Sprintf("http://localhost:%s", port))

	return &Config{
		DatabaseURL:        os.Getenv("DATABASE_URL"),
		Port:               port,
		Environment:        env,
		BaseURL:            baseURL,
		GithubClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		GithubClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		SessionSecret:      os.Getenv("SESSION_SECRET"),
	}
}

// getEnvOrDefault returns the environment variable value or a default.
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

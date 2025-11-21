package config

import "os"

// Config holds application configuration
type Config struct {
	DatabaseURL string
	Port        string
}

// Load reads configuration from environment variables
func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        port,
	}
}

package config

import "os"

// Config holds application configuration values.
type Config struct {
	Port string
}

// Load reads configuration from environment variables, applying
// sensible defaults when they are not set.
func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return Config{Port: port}
}

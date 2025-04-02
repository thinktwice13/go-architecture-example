package bootstrap

import (
	"log"
	"os"
)

// Config holds application configuration
type Config struct {
	DBConnString   string
	NATSConnString string
	HTTPPort       string
	Logger         *log.Logger
}

// NewConfig creates a new configuration with defaults and environment overrides
func NewConfig() *Config {
	return &Config{
		DBConnString:   getEnv("DB_CONN_STRING", "postgres://postgres:postgres@localhost:5432/documents?sslmode=disable"),
		NATSConnString: getEnv("NATS_CONN_STRING", "nats://localhost:4222"),
		HTTPPort:       getEnv("HTTP_PORT", "8080"),
		Logger:         log.New(os.Stdout, "[FCIS] ", log.LstdFlags),
	}
}

// getEnv retrieves environment variables with defaults
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

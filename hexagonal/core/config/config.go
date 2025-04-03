package config

// Config holds application configuration
type Config struct {
}

// Load loads configuration from environment variables
func Load() *Config {
	return &Config{}
}

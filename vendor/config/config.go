package config

import (
	"os"
)

// Config stores the configuration values for the application.
type Config struct {
	DatabaseURL string
	StripeKey   string
}

// NewConfig creates a new configuration object by reading values from environment variables.
func NewConfig() *Config {
	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", ""),
		StripeKey:   getEnv("STRIPE_KEY", ""),
	}
}

// getEnv gets the value of an environment variable with the specified key. If the variable
// is not set, it returns the specified default value.
func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}

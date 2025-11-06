package config

import (
	"os"
	"strconv"
)

// Config holds the application configuration
type Config struct {
	Port            string
	LogLevel        string
	ReadBufferSize  int
	WriteBufferSize int
}

// Load reads configuration from environment variables with defaults
func Load() *Config {
	return &Config{
		Port:            getEnv("PORT", "8080"),
		LogLevel:        getEnv("LOG_LEVEL", "info"),
		ReadBufferSize:  getEnvAsInt("READ_BUFFER_SIZE", 1024),
		WriteBufferSize: getEnvAsInt("WRITE_BUFFER_SIZE", 1024),
	}
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt retrieves an environment variable as int or returns a default value
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}

	return value
}

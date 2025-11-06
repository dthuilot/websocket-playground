package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name     string
		envVars  map[string]string
		expected Config
	}{
		{
			name:    "default values",
			envVars: map[string]string{},
			expected: Config{
				Port:            "8080",
				LogLevel:        "info",
				ReadBufferSize:  1024,
				WriteBufferSize: 1024,
			},
		},
		{
			name: "custom values",
			envVars: map[string]string{
				"PORT":              "9000",
				"LOG_LEVEL":         "debug",
				"READ_BUFFER_SIZE":  "2048",
				"WRITE_BUFFER_SIZE": "2048",
			},
			expected: Config{
				Port:            "9000",
				LogLevel:        "debug",
				ReadBufferSize:  2048,
				WriteBufferSize: 2048,
			},
		},
		{
			name: "invalid buffer size defaults",
			envVars: map[string]string{
				"READ_BUFFER_SIZE":  "invalid",
				"WRITE_BUFFER_SIZE": "invalid",
			},
			expected: Config{
				Port:            "8080",
				LogLevel:        "info",
				ReadBufferSize:  1024,
				WriteBufferSize: 1024,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clear environment
			os.Clearenv()

			// Set test environment variables
			for key, value := range tt.envVars {
				os.Setenv(key, value)
			}

			// Load configuration
			cfg := Load()

			// Verify
			if cfg.Port != tt.expected.Port {
				t.Errorf("Port: got %s, want %s", cfg.Port, tt.expected.Port)
			}
			if cfg.LogLevel != tt.expected.LogLevel {
				t.Errorf("LogLevel: got %s, want %s", cfg.LogLevel, tt.expected.LogLevel)
			}
			if cfg.ReadBufferSize != tt.expected.ReadBufferSize {
				t.Errorf("ReadBufferSize: got %d, want %d", cfg.ReadBufferSize, tt.expected.ReadBufferSize)
			}
			if cfg.WriteBufferSize != tt.expected.WriteBufferSize {
				t.Errorf("WriteBufferSize: got %d, want %d", cfg.WriteBufferSize, tt.expected.WriteBufferSize)
			}
		})
	}
}

func TestGetEnv(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		defaultValue string
		envValue     string
		expected     string
	}{
		{
			name:         "uses default when env not set",
			key:          "TEST_KEY",
			defaultValue: "default",
			envValue:     "",
			expected:     "default",
		},
		{
			name:         "uses env value when set",
			key:          "TEST_KEY",
			defaultValue: "default",
			envValue:     "custom",
			expected:     "custom",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Clearenv()
			if tt.envValue != "" {
				os.Setenv(tt.key, tt.envValue)
			}

			result := getEnv(tt.key, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("got %s, want %s", result, tt.expected)
			}
		})
	}
}

func TestGetEnvAsInt(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		defaultValue int
		envValue     string
		expected     int
	}{
		{
			name:         "uses default when env not set",
			key:          "TEST_INT",
			defaultValue: 100,
			envValue:     "",
			expected:     100,
		},
		{
			name:         "uses env value when set",
			key:          "TEST_INT",
			defaultValue: 100,
			envValue:     "200",
			expected:     200,
		},
		{
			name:         "uses default when env value is invalid",
			key:          "TEST_INT",
			defaultValue: 100,
			envValue:     "invalid",
			expected:     100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Clearenv()
			if tt.envValue != "" {
				os.Setenv(tt.key, tt.envValue)
			}

			result := getEnvAsInt(tt.key, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}

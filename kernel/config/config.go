package config

import "os"

// GetEnv
func GetEnv(key string) (value string) {
	return os.Getenv(key)
}

// GetDefaultEnv
func GetDefaultEnv(key, defaultValue string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

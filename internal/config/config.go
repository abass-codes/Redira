package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName     string
	Environment string
	ServerPort  string
	DatabaseURL string
	RedisURL    string
}

func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		AppName:     getEnv("APP_NAME", "Redira"),
		Environment: getEnv("APP_ENV", "development"),
		ServerPort:  getEnv("SERVER_PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", ""),
		RedisURL:    getEnv("REDIS_URL", ""),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

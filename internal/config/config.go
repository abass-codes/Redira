package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName     string
	AppEnv      string
	ServerPort  string
	DatabaseURL string
	RedisURL    string
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		AppName:     getEnv("APP_NAME", "Redira"),
		AppEnv:      getEnv("APP_ENV", "development"),
		ServerPort:  getEnv("SERVER_PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", ""),
		RedisURL:    getEnv("REDIS_URL", ""),
	}

	log.Printf("Loaded config (%s)", cfg.AppEnv)

	return cfg
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

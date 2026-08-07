package config

import (
	"log"
	"os"
)

type Config struct {
	AppName     string
	ServerPort  string
	DatabaseURL string
	RedisURL    string
	JWTSecret   string
	Environment string
}

func Load() Config {

	cfg := Config{
		AppName:     getEnv("APP_NAME", "Redira"),
		ServerPort:  getEnv("SERVER_PORT", "8080"),
		DatabaseURL: getEnvRequired("DATABASE_URL"),
		RedisURL:    getEnvRequired("REDIS_URL"),
		JWTSecret:   getEnvRequired("JWT_SECRET"),
		Environment: getEnv("ENVIRONMENT", "development"),
	}

	log.Printf(
		"Loaded config (%s)",
		cfg.Environment,
	)

	return cfg
}

func getEnv(key string, fallback string) string {

	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}

func getEnvRequired(key string) string {

	value := os.Getenv(key)

	if value == "" {
		log.Fatalf(
			"missing required environment variable: %s",
			key,
		)
	}

	return value
}

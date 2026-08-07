package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string

	Environment string

	ServerPort string

	DatabaseURL string

	RedisURL string

	JWTSecret string
}

func Load() Config {

	godotenv.Load()

	cfg := Config{

		AppName: getEnv(
			"APP_NAME",
			"Redira",
		),

		Environment: getEnv(
			"APP_ENV",
			"development",
		),

		ServerPort: getEnv(
			"SERVER_PORT",
			"8080",
		),

		DatabaseURL: required(
			"DATABASE_URL",
		),

		RedisURL: required(
			"REDIS_URL",
		),

		JWTSecret: required(
			"JWT_SECRET",
		),
	}

	log.Printf(
		"Loaded config (%s)",
		cfg.Environment,
	)

	return cfg

}

func getEnv(
	key string,
	fallback string,
) string {

	value := os.Getenv(key)

	if value == "" {

		return fallback

	}

	return value

}

func required(
	key string,
) string {

	value := os.Getenv(key)

	if value == "" {

		log.Fatalf(
			"missing required environment variable: %s",
			key,
		)

	}

	return value

}

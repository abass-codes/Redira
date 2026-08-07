package main

import (
	"log"

	"github.com/abass-codes/redira/internal/analytics"
	"github.com/abass-codes/redira/internal/cache"
	"github.com/abass-codes/redira/internal/config"
	"github.com/abass-codes/redira/internal/database"
	apphttp "github.com/abass-codes/redira/internal/http"
	"github.com/abass-codes/redira/internal/links"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.Load()

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	log.Println("✅ Connected to PostgreSQL")

	redisCache, err := cache.New(cfg.RedisURL)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Connected to Redis")

	linkRepository := links.NewRepository(
		db.Queries,
	)

	analyticsRepository := analytics.NewRepository(
		db.Queries,
	)

	analyticsService := analytics.NewService(
		analyticsRepository,
	)

	linkService := links.NewService(
		linkRepository,
		redisCache,
		analyticsService,
	)

	linkHandler := links.NewHandler(
		linkService,
	)

	analyticsHandler := analytics.NewHandler(
		analyticsService,
	)

	router := gin.New()

	apphttp.RegisterRoutes(
		router,
		linkHandler,
		analyticsHandler,
		redisCache.Client,
	)

	log.Printf(
		"🚀 %s starting on http://localhost:%s",
		cfg.AppName,
		cfg.ServerPort,
	)

	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal(err)
	}
}

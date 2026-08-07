package main

import (
	"log"

	"github.com/abass-codes/redira/internal/analytics"
	"github.com/abass-codes/redira/internal/auth"
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

	// Link dependencies

	linkRepository := links.NewRepository(
		db.Queries,
	)

	userLinkRepository := links.NewUserRepository(
		db.Queries,
	)

	// Analytics dependencies

	analyticsRepository := analytics.NewRepository(
		db.Queries,
	)

	analyticsService := analytics.NewService(
		analyticsRepository,
	)

	// Link service

	linkService := links.NewService(
		linkRepository,
		userLinkRepository,
		redisCache,
		analyticsService,
	)

	linkHandler := links.NewHandler(
		linkService,
	)

	userLinkHandler := links.NewUserHandler(
		userLinkRepository,
	)

	// Feature 6: Link Management

	managementService := links.NewManagementService(
		db.Queries,
	)

	managementHandler := links.NewManagementHandler(
		managementService,
	)

	// Authentication

	authRepository := auth.NewRepository(
		db.Queries,
	)

	authService := auth.NewService(
		authRepository,
	)

	authHandler := auth.NewHandler(
		authService,
	)

	// Analytics handler

	analyticsHandler := analytics.NewHandler(
		analyticsService,
	)

	router := gin.New()

	apphttp.RegisterRoutes(
		router,
		linkHandler,
		userLinkHandler,
		managementHandler,
		analyticsHandler,
		authHandler,
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

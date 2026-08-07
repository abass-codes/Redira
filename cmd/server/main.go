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
	"github.com/abass-codes/redira/internal/redirect"

	"github.com/gin-contrib/cors"
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

	// Existing link system

	linkRepository := links.NewRepository(
		db.Queries,
	)

	userLinkRepository := links.NewUserRepository(
		db.Queries,
	)

	analyticsRepository := analytics.NewRepository(
		db.Queries,
	)

	analyticsService := analytics.NewService(
		analyticsRepository,
	)

	// Feature 8 - Analytics Dashboard

	dashboardRepository := analytics.NewDashboardRepository(
		db.Queries,
	)

	dashboardService := analytics.NewDashboardService(
		dashboardRepository,
	)

	dashboardHandler := analytics.NewDashboardHandler(
		dashboardService,
	)

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

	// Feature 6 - Link Management

	managementService := links.NewManagementService(
		db.Queries,
	)

	managementHandler := links.NewManagementHandler(
		managementService,
	)

	// Feature 7 - Production Redirect Engine

	redirectService := redirect.NewService(
		db.Queries,
		redisCache.Client,
	)

	redirectHandler := redirect.NewHandler(
		redirectService,
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

	analyticsHandler := analytics.NewHandler(
		analyticsService,
	)

	router := gin.New()

	// Feature 9 - Frontend CORS

	router.Use(
		cors.New(cors.Config{
			AllowOrigins: []string{
				"http://localhost:3000",
			},
			AllowMethods: []string{
				"GET",
				"POST",
				"PATCH",
				"DELETE",
				"OPTIONS",
			},
			AllowHeaders: []string{
				"Origin",
				"Content-Type",
				"Authorization",
			},
			AllowCredentials: true,
		}),
	)

	apphttp.RegisterRoutes(
		router,
		redirectHandler,
		linkHandler,
		userLinkHandler,
		managementHandler,
		analyticsHandler,
		dashboardHandler,
		authHandler,
		redisCache.Client,
	)

	log.Printf(
		"🚀 %s starting on http://localhost:%s",
		cfg.AppName,
		cfg.ServerPort,
	)

	if err := router.Run(
		":" + cfg.ServerPort,
	); err != nil {
		log.Fatal(err)
	}
}

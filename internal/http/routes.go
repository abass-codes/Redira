package http

import (
	"time"

	"github.com/abass-codes/redira/internal/analytics"
	"github.com/abass-codes/redira/internal/auth"
	"github.com/abass-codes/redira/internal/health"
	"github.com/abass-codes/redira/internal/links"
	"github.com/abass-codes/redira/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RegisterRoutes(
	router *gin.Engine,
	linkHandler *links.Handler,
	analyticsHandler *analytics.Handler,
	authHandler *auth.Handler,
	redisClient *redis.Client,
) {

	router.GET(
		"/health",
		health.Handler,
	)

	router.GET(
		"/r/:shortCode",
		linkHandler.Redirect,
	)

	v1 := router.Group("/api/v1")

	v1.Use(
		middleware.RateLimiter(
			redisClient,
			100,
			time.Minute,
		),
	)

	{

		// Authentication

		v1.POST(
			"/auth/register",
			authHandler.Register,
		)

		v1.POST(
			"/auth/login",
			authHandler.Login,
		)

		// Protected link routes

		v1.POST(
			"/links",
			linkHandler.Create,
		)

		v1.GET(
			"/analytics/:id",
			analyticsHandler.Get,
		)
	}
}

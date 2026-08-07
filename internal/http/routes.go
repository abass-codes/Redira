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
	userLinkHandler *links.UserHandler,
	managementHandler *links.ManagementHandler,
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

	// Authentication

	v1.POST(
		"/auth/register",
		authHandler.Register,
	)

	v1.POST(
		"/auth/login",
		authHandler.Login,
	)

	// Protected routes

	protected := v1.Group("")

	protected.Use(
		middleware.AuthRequired(),
	)

	// Link creation

	protected.POST(
		"/links",
		linkHandler.Create,
	)

	// User link list/delete

	protected.GET(
		"/links",
		userLinkHandler.GetMyLinks,
	)

	protected.DELETE(
		"/links/:id",
		userLinkHandler.DeleteMyLink,
	)

	// Feature 6 - Link Management

	protected.GET(
		"/links/:id",
		managementHandler.Get,
	)

	protected.PATCH(
		"/links/:id",
		managementHandler.Update,
	)

	protected.POST(
		"/links/:id/disable",
		managementHandler.Disable,
	)

	protected.POST(
		"/links/:id/enable",
		managementHandler.Enable,
	)

	// Analytics

	v1.GET(
		"/analytics/:id",
		analyticsHandler.Get,
	)
}

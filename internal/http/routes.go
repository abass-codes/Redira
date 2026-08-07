package http

import (
	"time"

	"github.com/abass-codes/redira/internal/analytics"
	"github.com/abass-codes/redira/internal/auth"
	"github.com/abass-codes/redira/internal/links"
	"github.com/abass-codes/redira/internal/middleware"
	"github.com/abass-codes/redira/internal/redirect"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RegisterRoutes(
	router *gin.Engine,

	redirectHandler *redirect.Handler,

	linkHandler *links.Handler,

	userLinkHandler *links.UserHandler,

	managementHandler *links.ManagementHandler,

	analyticsHandler *analytics.Handler,

	authHandler *auth.Handler,

	redisClient *redis.Client,
) {

	router.GET(
		"/health",
		func(c *gin.Context) {
			c.JSON(
				200,
				gin.H{
					"status": "ok",
				},
			)
		},
	)

	// Feature 7 production redirect

	router.GET(
		"/r/:shortCode",
		redirectHandler.Redirect,
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

	protected := v1.Group("")

	protected.Use(
		middleware.AuthRequired(),
	)

	// Create links

	protected.POST(
		"/links",
		linkHandler.Create,
	)

	// User links

	protected.GET(
		"/links",
		userLinkHandler.GetMyLinks,
	)

	protected.DELETE(
		"/links/:id",
		userLinkHandler.DeleteMyLink,
	)

	// Feature 6

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

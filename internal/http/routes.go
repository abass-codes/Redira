package http

import (
	"github.com/abass-codes/redira/internal/health"
	"github.com/abass-codes/redira/internal/links"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, linkHandler *links.Handler) {
	router.GET("/health", health.Handler)

	router.GET("/r/:shortCode", linkHandler.Redirect)

	v1 := router.Group("/api/v1")
	{
		v1.POST("/links", linkHandler.Create)
	}
}

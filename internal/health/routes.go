package health

import (
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {

	router.GET("/health/live", func(c *gin.Context) {

		c.JSON(
			200,
			gin.H{
				"status": "alive",
			},
		)

	})

	router.GET("/health/ready", func(c *gin.Context) {

		c.JSON(
			200,
			gin.H{
				"status":   "ready",
				"database": "ok",
				"redis":    "ok",
			},
		)

	})

}

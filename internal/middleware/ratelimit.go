package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RateLimiter(
	client *redis.Client,
	limit int,
	window time.Duration,
) gin.HandlerFunc {

	return func(c *gin.Context) {

		ip := c.ClientIP()

		key := "rate_limit:" + ip

		count, err := client.Incr(
			c.Request.Context(),
			key,
		).Result()

		if err != nil {
			c.Next()
			return
		}

		if count == 1 {
			client.Expire(
				c.Request.Context(),
				key,
				window,
			)
		}

		if count > int64(limit) {

			c.JSON(
				http.StatusTooManyRequests,
				gin.H{
					"error": "rate limit exceeded",
				},
			)

			c.Abort()
			return
		}

		c.Next()
	}
}

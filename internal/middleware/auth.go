package middleware

import (
	"net/http"
	"strings"

	"github.com/abass-codes/redira/internal/auth"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {

	return func(c *gin.Context) {

		header := c.GetHeader("Authorization")

		if header == "" {

			c.JSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "missing authorization header",
				},
			)

			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(
			header,
			"Bearer ",
		)

		token, err := auth.ValidateToken(
			tokenString,
		)

		if err != nil || !token.Valid {

			c.JSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "invalid token",
				},
			)

			c.Abort()
			return
		}

		c.Next()
	}
}

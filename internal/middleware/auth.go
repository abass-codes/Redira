package middleware

import (
	"net/http"
	"strings"

	"github.com/abass-codes/redira/internal/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func AuthRequired() gin.HandlerFunc {

	return func(c *gin.Context) {

		header := c.GetHeader("Authorization")

		if header == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "missing authorization header",
			})
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
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid claims",
			})
			c.Abort()
			return
		}

		id, ok := claims["user_id"].(string)

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "missing user id",
			})
			c.Abort()
			return
		}

		userID, err := uuid.Parse(id)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid user id",
			})
			c.Abort()
			return
		}

		c.Set(
			"user_id",
			userID,
		)

		c.Next()
	}
}

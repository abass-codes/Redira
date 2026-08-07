package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("redira-secret")

func GenerateToken(
	userID string,
) (string, error) {

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": userID,
			"exp": time.Now().
				Add(24 * time.Hour).
				Unix(),
		},
	)

	return token.SignedString(jwtSecret)
}

func ValidateToken(
	tokenString string,
) (*jwt.Token, error) {

	return jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		},
	)
}

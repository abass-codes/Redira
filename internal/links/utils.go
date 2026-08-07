package links

import (
	"math/rand"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateShortCode() string {

	rand.Seed(time.Now().UnixNano())

	code := make([]byte, 6)

	for i := range code {
		code[i] = letters[rand.Intn(len(letters))]
	}

	return string(code)
}

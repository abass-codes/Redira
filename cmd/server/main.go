package main

import (
	"log"

	"github.com/abass-codes/redira/internal/health"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.GET("/health", health.Handler)

	log.Println("🚀 Redira server starting on http://localhost:8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

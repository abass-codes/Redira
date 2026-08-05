package main

import (
	"log"

	"github.com/abass-codes/redira/internal/config"
	"github.com/abass-codes/redira/internal/database"
	"github.com/abass-codes/redira/internal/health"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("✅ Connected to PostgreSQL")

	router := gin.New()

	router.GET("/health", health.Handler)

	log.Printf("🚀 %s starting on http://localhost:%s", cfg.AppName, cfg.ServerPort)

	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal(err)
	}
}

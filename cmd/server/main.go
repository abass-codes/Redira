package main

import (
	"log"

	"github.com/abass-codes/redira/internal/config"
	"github.com/abass-codes/redira/internal/database"
	apphttp "github.com/abass-codes/redira/internal/http"
	"github.com/abass-codes/redira/internal/links"
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

	repository := links.NewRepository(db.Queries)
	service := links.NewService(repository)
	handler := links.NewHandler(service)

	router := gin.New()

	apphttp.RegisterRoutes(router, handler)

	log.Printf("🚀 %s starting on http://localhost:%s", cfg.AppName, cfg.ServerPort)

	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pg-geoshard/services/library-store/internal/application"
	"github.com/pg-geoshard/services/library-store/internal/domain"
	"github.com/pg-geoshard/services/library-store/internal/infrastructure/postgres"
	"github.com/pg-geoshard/services/library-store/internal/interfaces/http"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Database connection
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=library_store port=5432 sslmode=disable"
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate the schema
	if err := db.AutoMigrate(&domain.Book{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize dependencies
	bookRepo := postgres.NewBookRepository(db)
	bookService := application.NewBookService(bookRepo)
	bookHandler := http.NewBookHandler(bookService)

	// Setup Gin router
	router := gin.Default()

	// Register routes
	bookHandler.RegisterRoutes(router)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

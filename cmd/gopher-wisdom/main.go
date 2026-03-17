package main

import (
	"fmt"
	"log"
	"os"

	"github.com/EternalHalve/gopher-wisdom/internal/quotes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Warning: No .env file found")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "wisdom.db"
	}

	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatalf("Fatal: Could not connect to database: %v", err)
	}

	db.AutoMigrate(&quotes.Quote{})

	quoteHandler := quotes.NewQuoteHandler(db)

	router := gin.Default()

	router.GET("/quotes", quoteHandler.GetQuotes)
	router.GET("/quotes/:id", quoteHandler.GetQuotesByID)
	router.POST("/quotes", quoteHandler.PostQuotes)

	router.Run("localhost:8080")
}

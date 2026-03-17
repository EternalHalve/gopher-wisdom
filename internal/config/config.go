package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	DB *gorm.DB
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Info: No .env file found")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "wisdom.db"
	}

	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatalf("Fatal: Could not connect to database: %v", err)
	}

	return &Config{DB: db}
}

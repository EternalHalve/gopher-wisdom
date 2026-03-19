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
		log.Println("Gopher Wisdom: No external environment (.env) found. Relying on internal defaults.")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "wisdom.db"
	}

	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gopher Wisdom: The archives are buried too deep! Error: %v", err)
	}

	log.Printf("Gopher Wisdom: Connected to the source of truth [%s]\n", dbName)

	return &Config{DB: db}
}

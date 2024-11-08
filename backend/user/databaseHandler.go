package user

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func DatabaseInitialization() error {
	godotenv.Load()
	dsn := os.Getenv("DATABASE_DATA")

	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("failed to connect database", err)
	}

	return nil
}

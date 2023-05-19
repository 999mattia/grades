package db

import (
	"fmt"
	"log"
	"os"

	"github.com/999mattia/grades/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	fmt.Println("Connection to database...")

	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	DB = db

	fmt.Println("Connected!")
}

func Migrate() {
	fmt.Println("Migrating...")
	DB.AutoMigrate(&models.User{}, &models.Module{}, &models.Grade{})
	fmt.Println("Migrated!")
}

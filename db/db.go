package db

import (
	"fmt"
	"log"

	"github.com/999mattia/grades/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	fmt.Println("Connection to database...")

	db, err := gorm.Open(postgres.Open("postgres://jwnybpvo:tvKkBSR1LoQiQSBtsw5IA2o9bUNfTXuY@kandula.db.elephantsql.com/jwnybpvo"), &gorm.Config{})

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

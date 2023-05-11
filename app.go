package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	DB, _ = gorm.Open(postgres.Open("postgres://jwnybpvo:tvKkBSR1LoQiQSBtsw5IA2o9bUNfTXuY@kandula.db.elephantsql.com/jwnybpvo"), &gorm.Config{})

	r := gin.Default()

	r.Run()
}

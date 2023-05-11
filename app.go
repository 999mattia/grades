package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Grade struct {
	gorm.Model
	Name     string
	Grade    float32
	ModuleId int
	UserId   int
}

type Module struct {
	gorm.Model
	Name     string
	Semester int
	UserId   int
}

type User struct {
	gorm.Model
	Name     string
	Password string
}

var DB *gorm.DB

func main() {
	DB, _ = gorm.Open(postgres.Open("postgres://jwnybpvo:tvKkBSR1LoQiQSBtsw5IA2o9bUNfTXuY@kandula.db.elephantsql.com/jwnybpvo"), &gorm.Config{})

	DB.AutoMigrate(&User{}, &Module{}, &Grade{})

	r := gin.Default()

	r.Run()
}

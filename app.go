package main

import (
	"fmt"

	"github.com/999mattia/grades/controllers"
	"github.com/999mattia/grades/db"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Connection to database...")
	db.Connect()
	fmt.Println("Connected!")

	fmt.Println("Migrating...")
	db.Migrate()
	fmt.Println("Migrated!")

	r := gin.Default()

	r.POST("/signup", controllers.SignUp)

	r.Run()
}

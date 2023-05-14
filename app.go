package main

import (
	"github.com/999mattia/grades/controllers"
	"github.com/999mattia/grades/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	//db.Migrate()

	r := gin.Default()

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)

	r.Run()
}

package main

import (
	"github.com/999mattia/grades/controllers"
	"github.com/999mattia/grades/db"
	"github.com/999mattia/grades/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	//db.Migrate()

	r := gin.Default()

	r.POST("/api/signup", controllers.SignUp)
	r.POST("/api/login", controllers.Login)

	r.GET("/api/user/:id", middleware.Auth, controllers.GetUserById)

	r.POST("/api/module", middleware.Auth, controllers.CreateModule)

	r.Static("/assets", "./wwwroot/assets")
	r.StaticFile("/", "./wwwroot/index.html")

	r.Run()
}

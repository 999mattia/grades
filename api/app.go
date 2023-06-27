package main

import (
	"github.com/999mattia/grades/controllers"
	"github.com/999mattia/grades/db"
	"github.com/999mattia/grades/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	db.Connect()

	//db.Migrate()

	r := gin.Default()

	r.Use(middleware.CORS)

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)

	r.GET("/user/:id", middleware.Auth, controllers.GetUserById)

	r.POST("/module", middleware.Auth, controllers.CreateModule)

	r.Run()
}

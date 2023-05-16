package main

import (
	"fmt"
	"os"

	"github.com/999mattia/grades/controllers"
	"github.com/999mattia/grades/db"
	"github.com/999mattia/grades/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	fmt.Println("test: ", os.Getenv("TEST"))

	db.Connect()

	//db.Migrate()

	r := gin.Default()

	r.Use(middleware.CORS)

	r.POST("/api/signup", controllers.SignUp)
	r.POST("/api/login", controllers.Login)

	r.GET("/api/user/:id", middleware.Auth, controllers.GetUserById)

	r.POST("/api/module", middleware.Auth, controllers.CreateModule)

	r.GET("/api/spotify/current", controllers.GetCurrentOrLast)

	r.Static("/assets", "./client/dist/assets")
	r.StaticFile("/", "./client/dist/index.html")

	r.Run()
}

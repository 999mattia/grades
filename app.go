package main

import (
	"net/http"

	"github.com/999mattia/grades/controllers"
	"github.com/999mattia/grades/db"
	"github.com/999mattia/grades/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	//db.Migrate()

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	})

	r.POST("/api/signup", controllers.SignUp)
	r.POST("/api/login", controllers.Login)

	r.GET("/api/user/:id", middleware.Auth, controllers.GetUserById)

	r.POST("/api/module", middleware.Auth, controllers.CreateModule)

	r.GET("/api/spotify/current", controllers.GetCurrentOrLast)

	r.Static("/assets", "./client/dist/assets")
	r.StaticFile("/", "./client/dist/index.html")

	r.Run()
}

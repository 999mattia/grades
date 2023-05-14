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

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)

	r.GET("/auth", middleware.Auth, func(c *gin.Context) {
		user, _ := c.Get("user")
		c.JSON(200, gin.H{"user": user})
	})

	r.Run()
}

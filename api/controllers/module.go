package controllers

import (
	"github.com/999mattia/grades/db"
	"github.com/999mattia/grades/models"
	"github.com/gin-gonic/gin"
)

func CreateModule(c *gin.Context) {
	var body struct {
		Name     string
		Semester int
		UserId   uint
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{"error": "Failed to read body"})
		return
	}

	user, _ := c.Get("user")

	if user.(models.User).ID != body.UserId {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	module := models.Module{Name: body.Name, Semester: body.Semester, UserId: body.UserId}

	result := db.DB.Create(&module)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create module"})
	}

	c.JSON(200, gin.H{"message": "Module created"})
}

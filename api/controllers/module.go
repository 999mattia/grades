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
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{"error": "Failed to read body"})
		return
	}

	user, _ := c.Get("user")

	module := models.Module{Name: body.Name, Semester: body.Semester, UserId: user.(models.User).ID}

	result := db.DB.Create(&module)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create module"})
	}

	c.JSON(200, gin.H{"message": "Module created"})
}

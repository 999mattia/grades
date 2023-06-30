package controllers

import (
	"github.com/999mattia/grades/db"
	"github.com/999mattia/grades/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func DeleteModule(c *gin.Context) {
	id := c.Param("id")

	var module models.Module

	if err := db.DB.First(&module, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"error": "Module not found"})
			return
		}
	}

	user, _ := c.Get("user")

	if module.UserId != user.(models.User).ID {
		c.JSON(403, gin.H{"error": "You don't own this module"})
		return
	}

	db.DB.Delete(&module)

	c.JSON(200, gin.H{"message": "Module deleted"})
}

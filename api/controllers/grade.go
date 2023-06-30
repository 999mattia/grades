package controllers

import (
	"strconv"

	"github.com/999mattia/grades/db"
	"github.com/999mattia/grades/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateGrade(c *gin.Context) {
	moduleId, _ := strconv.Atoi(c.Param("id"))
	tokenUser, _ := c.Get("user")

	var body struct {
		Name  string
		Grade float32
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{"error": "Failed to read body"})
		return
	}

	var module models.Module
	db.DB.First(&module, moduleId)

	if module.ID == 0 {
		c.JSON(404, gin.H{"error": "Module not found"})
		return
	}

	if module.UserId != tokenUser.(models.User).ID {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	grade := models.Grade{ModuleId: uint(moduleId), UserId: tokenUser.(models.User).ID, Name: body.Name, Grade: body.Grade}

	result := db.DB.Create(&grade)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create grade"})
		return
	}

	c.JSON(200, gin.H{"message": "Grade created"})
}

func DeleteGrade(c *gin.Context) {
	id := c.Param("id")

	var grade models.Grade

	if err := db.DB.First(&grade, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"error": "Grade not found"})
			return
		}
	}

	tokenUser, _ := c.Get("user")

	if grade.UserId != tokenUser.(models.User).ID {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	db.DB.Delete(&grade)

	c.JSON(200, gin.H{"message": "Grade deleted"})
}

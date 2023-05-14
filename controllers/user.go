package controllers

import (
	"strconv"

	"github.com/999mattia/grades/db"
	"github.com/999mattia/grades/models"
	"github.com/gin-gonic/gin"
)

func GetUserById(c *gin.Context) {
	requestedId := c.Param("id")
	requestedIdInt, _ := strconv.Atoi(requestedId)
	tokenUser, _ := c.Get("user")

	if tokenUser.(models.User).ID != uint(requestedIdInt) {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	var user models.User
	db.DB.First(&user, requestedId)

	if user.ID == 0 {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	var modules []models.Module
	db.DB.Where("user_id = ?", user.ID).Find(&modules)

	var grades []models.Grade
	db.DB.Where("user_id = ?", user.ID).Find(&grades)

	c.JSON(200, gin.H{"user": user, "modules": modules, "grades": grades})
}

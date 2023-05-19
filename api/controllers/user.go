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

	var gradesDTO []models.GradeDTO
	var modulesDTO []models.ModuleDTO

	for _, grade := range grades {
		gradesDTO = append(gradesDTO, models.GradeDTO{Id: grade.ID, Name: grade.Name, Grade: grade.Grade, ModuleId: grade.ModuleId})
	}

	for _, module := range modules {
		var gradesForModule []models.GradeDTO

		for _, grade := range gradesDTO {
			if grade.ModuleId == module.ID {
				gradesForModule = append(gradesForModule, grade)
			}
		}

		modulesDTO = append(modulesDTO, models.ModuleDTO{Id: module.ID, Name: module.Name, Grades: gradesForModule, Semester: module.Semester})
	}

	userDTO := models.UserDTO{Id: user.ID, Name: user.Name, Modules: modulesDTO}

	c.JSON(200, userDTO)
}

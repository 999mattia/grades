package controllers

import (
	"net/http"
	"time"

	"github.com/999mattia/grades/db"
	"github.com/999mattia/grades/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var body struct {
		Name     string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{"error": "Failed to read body"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{Name: body.Name, Password: string(hash)}

	result := db.DB.Create(&user)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(200, gin.H{"message": "User created"})
}

func Login(c *gin.Context) {
	var body struct {
		Name     string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{"error": "Failed to read body"})
		return
	}

	var user models.User
	db.DB.First(&user, "name = ?", body.Name)

	if user.ID == 0 {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(401, gin.H{"error": "Wrong password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte("thisismysecret"))

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(200, gin.H{"token": tokenString})
}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Valid"})
}

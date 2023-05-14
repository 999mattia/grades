package middleware

import (
	"fmt"
	"time"

	"github.com/999mattia/grades/db"
	"github.com/999mattia/grades/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth(c *gin.Context) {
	if c.Request.Header["Authorization"] == nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "Missing Authorization header"})
		return
	}

	tokenString := c.Request.Header["Authorization"][0]

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("thisismysecret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(401, gin.H{"error": "Token expired"})
		}

		var user models.User
		db.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatusJSON(404, gin.H{"error": "User not found, invalid token"})
		}

		c.Set("user", user)

		c.Next()
	} else {
		c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
	}

}

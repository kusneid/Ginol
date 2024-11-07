package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"Ginol/user"
)

func main() {
	r := gin.Default()

	r.POST("/api/register", func(c *gin.Context) {
		var userAdded user.User
		if err := c.ShouldBindJSON(&userAdded); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		if err := userAdded.Crypt(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrypt password"})
			return
		}

		//fmt.Print(userAdded.Username, userAdded.Password) для теста
		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
	})

	r.POST("/api/login", func(c *gin.Context) {
		var credentials user.Credentials
		if err := c.ShouldBindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	})

	r.Run(":8080")
}

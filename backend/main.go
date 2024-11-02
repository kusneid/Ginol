package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Структура данных для пользователя и учетных данных

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	r := gin.Default()

	r.POST("/api/register", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
	})

	r.POST("/api/login", func(c *gin.Context) {
		var credentials Credentials
		if err := c.ShouldBindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	})

	r.Static("frontend/", "./frontend/build/static")
	r.NoRoute(func(c *gin.Context) {
		c.File("./frontend/build/index.html")
	})
	r.Run(":8080")
}

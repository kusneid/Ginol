package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/kusneid/Ginol/backend/user"
)

func main() {
	r := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	r.POST("/api/register", func(c *gin.Context) {
		var userAdded user.Credentials
		if err := c.ShouldBindJSON(&userAdded); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		log.Println("Registration api call handled")

		userAdded.RegistrationHandler()

		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
	})

	r.POST("/api/login", func(c *gin.Context) {
		var credentials user.Credentials
		if err := c.ShouldBindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		log.Println("Login api call handled")

		authResult := credentials.LoginHandler()
		log.Println("auth result:", authResult)

		if authResult {
			c.Redirect(http.StatusMovedPermanently, "/connection")
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		}

		r.GET("/connection", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "You are connected!"})
		})

	})

	r.Run(":8080")
}

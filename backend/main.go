package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/kusneid/Ginol/backend/routes"
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
		}

		log.Println("Registration API call handled")

		regResult, token := userAdded.RegistrationHandler()
		log.Println("Registration result:", regResult)

		c.JSON(http.StatusOK, gin.H{"bool": true, "username": userAdded.Username, "token": token})
	})

	r.POST("/api/login", func(c *gin.Context) {
		var credentials user.Credentials
		if err := c.ShouldBindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		log.Println("Login API call handled")
		authResult, token := credentials.LoginHandler()
		log.Println("Auth result:", authResult)
		if !authResult {

			c.JSON(http.StatusConflict, gin.H{"loginStatus": "false"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"bool": true, "username": credentials.Username, "token": token})
	})

	r.POST("/api/check-nickname", func(c *gin.Context) {

		var union user.Answer
		if err := c.ShouldBindJSON(&union); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		if union.FriendNickname == union.LoggedUser {
			log.Fatalln("can't connect same accounts sorry bro")
			return

		}
		log.Println("check nickname api handled")
		//fmt.Println("ERR:", union.FriendNickname)
		value, err := user.SendCheckRequest(union.FriendNickname)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		}
		if value {
			c.JSON(http.StatusOK, gin.H{"exists": true})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"exists": false})
		}

	})

	r.POST("/api/messages", user.CreateMessage)
	r.GET("/api/messages", user.GetMessage)
	r.GET("/ws", routes.HandleWebSocket)

	go routes.HandleMessages()

	r.Run(":8080")
}

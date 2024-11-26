package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/kusneid/Ginol/user"
	"github.com/kusneid/Ginol/routes"
)

func main() {
	r := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	r.POST("/api/register", func(c *gin.Context) {
		var userAdded user.User
		if err := c.ShouldBindJSON(&userAdded); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		log.Println("Registration api call handled")

		if err := userAdded.RegistrationHandler(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})

		user.AddUser(userAdded)
	})

	r.POST("/api/login", func(c *gin.Context) {
		var credentials user.Credentials
		if err := c.ShouldBindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}
		log.Println("Login api call handled")
			c.JSON(http.StatusConflict, gin.H{"loginStatus": "false"})
			return
		

		if err := credentials.LoginHandler(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	})

	var chatInst routes.ChatInstance
	r.POST("/api/chat-reg", func(c *gin.Context) {
    if err := c.ShouldBindJSON(&chatInst); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
        return
    }

    log.Printf("Chat registration: Username=%s, Friend=%s", chatInst.Username, chatInst.FriendUsername)

    c.JSON(http.StatusOK, gin.H{"message": "Chat registration successful"})
	})
	
	r.GET("/ws/chat", func(c *gin.Context){
		routes.HandleWebSocket(c, chatInst)
	})
	// r.POST("/api/messages", user.CreateMessage)
	// r.GET("/api/messages", user.GetMessage)

	// go routes.HandleMessages()

	r.Run(":8080")
}

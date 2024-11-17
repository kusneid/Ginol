package user

// отправка и прием сообщений

import (
	"time"
	"net/http"

	 "github.com/gin-gonic/gin"
)

type Message struct {
    Nickname  string `json:"nickname"`
    Text      string `json:"text"`
    Time      time.Time `json:"time"`
    SenderID  int    `json:"senderID"`
}

var messages []Message

func CreateMessage(c *gin.Context) {
	var newMessage Message

    if err := c.ShouldBindJSON(&newMessage); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"json file parsing error": err.Error()})
        return
    }

    newMessage.Time = time.Now()

    messages = append(messages, newMessage)

    c.JSON(http.StatusCreated, newMessage)

}

func GetMessage(c *gin.Context) {
	c.JSON(http.StatusOK, messages)
}
/*type Connection struct {
    User1 *User `json:"user1`
    User2 *User `json:"user2`
    Messages []Message `json:"messages"`

}

var Connect Connection

func CreateMessage(c *gin.Context) {
	var newMessage Message

    if err := c.ShouldBindJSON(&newMessage); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"json file parsing error": err.Error()})
        return
    }

    newMessage.Time = time.Now()

    Connect.Messages = append(Connect.Messages, newMessage)

    c.JSON(http.StatusCreated, newMessage)

}

func GetMessage(c *gin.Context) {
	c.JSON(http.StatusOK, Connect.Messages)
}*/
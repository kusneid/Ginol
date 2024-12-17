package routes

//реализация websocket подключения, создание маршрута, всякие проверки подключения итд

import (
	"log"
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Message struct {
    Sender  string `json:"sender"`
    Text      string `json:"text"`
    Time      time.Time `json:"time"`
    Target  string    `json:"target"`
}

// var clients = make(map[*websocket.Conn]bool)
// var broadcast = make(chan user.Message)

type ChatInstance struct{
  Username	string `json:"username"`   
  FriendUsername	string` json:"friend"`
}

var clients = make(map[string]*websocket.Conn) // активные ws подключения

var upgrader = websocket.Upgrader{ // 
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocket(c *gin.Context, chatInst ChatInstance) {        // установка WebSocket соединения
	// username := chatInst.Username
	// friend := chatInst.FriendUsername
    username := c.Query("username")
    friend := c.Query("friend")
    
    if username == "" || friend == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Missing username or friend"})
        return
    }
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Println("Error upgrading to WebSocket:", err)
        return
    }
    defer ws.Close()

    clients[username] = ws // сохраняем соединение
    log.Printf("WebSocket connection established for user: %s", username)

    for {
        var msg Message
        err := ws.ReadJSON(&msg)
        if err != nil {
            log.Printf("Error reading message from %s: %v", username, err)
            delete(clients, username)
            break
        }

        msg.Sender = username
        msg.Time = time.Now()
        msg.Target = friend

        if friendConn, ok := clients[friend]; ok {
            err = friendConn.WriteJSON(msg)
            if err != nil {
                log.Printf("Error sending message to %s: %v", friend, err)
                delete(clients, friend)
            }
        } else {
            log.Printf("Friend %s is not connected", friend)
        }
    }
}

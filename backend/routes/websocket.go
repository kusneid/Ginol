package routes

//реализация websocket подключения, создание маршрута, всякие проверки подключения итд

import (
    "net/http"
	"log"
	"time"
    "github.com/gorilla/websocket"
    "github.com/gin-gonic/gin"

    "github.com/kusneid/Ginol/user"

)

type ChatInstance struct{
  Username	string `json:"username"`      /*binding:"required"*/
  FriendUsername	string` json:"friend"`
}

var clients = make(map[string]*websocket.Conn) // активные ws подключения
var broadcast = make(chan user.Message) // канал для передачи сообщений

var upgrader = websocket.Upgrader{ // обновление с http до ws
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocket(c *gin.Context, chatInst ChatInstance) {        // Установка WebSocket соединения
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

    clients[username] = ws // Сохраняем соединение
    log.Printf("WebSocket connection established for user: %s", username)

    for {
        var msg user.Message
        err := ws.ReadJSON(&msg)
        if err != nil {
            log.Printf("Error reading message from %s: %v", username, err)
            delete(clients, username)
            break
        }

        // Добавляем отправителя и время к сообщению
        msg.Nickname = username
        msg.Time = time.Now()

        // Отправляем сообщение другу
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



/*func HandleMessages() { // рассылка сообщений
	for {
		msg := <- broadcast
		for client := range clients {
			if msg.Nickname == clients[client] {
        		continue // Пропустить отправителя
    		}
>>>>>>> Stashed changes
			log.Printf("Sending message to %s: %s", clients[client], msg.Text)
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Error sending message to %s: %v", clients[client], err)
				client.Close()
				delete(clients, client)
			}
		}
	}
<<<<<<< Updated upstream
}

/*func ConnectUser(c *gin.Context) {
	username := c.DefaultQuery("username", "")
	var currentUser user.User

	for _, user_buf := range user.UsersSlice {
		if user_buf.Username == username {
			currentUser = user_buf
		}
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()

	if user.Connect.User1 == nil {
		user.Connect.User1 = &currentUser
	} else if user.Connect.User2 == nil {
		user.Connect.User2 = &currentUser
	}

	go HandleMessages()

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User %s connected", username)})
}*/

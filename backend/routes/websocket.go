package routes

//реализация websocket подключения, создание маршрута, всякие проверки подключения итд

import (
	//"fmt"
    "net/http"
	"log"
	"time"
    "github.com/gorilla/websocket"
    "github.com/gin-gonic/gin"
	
    "github.com/kusneid/Ginol/user"
)

var clients = make(map[*websocket.Conn]string) // активные ws подключения
var broadcast = make(chan user.Message) // канал для передачи сообщений

var upgrader = websocket.Upgrader{ // обновление с http до ws
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocket(c *gin.Context) { // обработка нового ws подключения
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error upgrading WebSocket connection:", err)
		return
	}
	defer ws.Close()
	log.Printf("New connection established")	

	var msg user.Message // обработка входящего сообщения
	
	for {
		err := ws.ReadJSON(&msg)
		clients[ws] = msg.Nickname
		if err != nil {
			log.Printf("Error reading message from client %s: %v", msg.Nickname, err)
			delete(clients, ws)
			break
		}
		msg.Time = time.Now()
		broadcast <- msg
		log.Printf("Received message from %s: %s", msg.Nickname, msg.Text)
	}
}

func HandleMessages() { // рассылка сообщений
	for {
		msg := <- broadcast
		for client := range clients {
			// if msg.Nickname == clients[client] {
        	// 	continue // Пропустить отправителя
    		// }
			log.Printf("Sending message to %s: %s", clients[client], msg.Text)
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Error sending message to %s: %v", clients[client], err)
				client.Close()
				delete(clients, client)
			}
		}
	}
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


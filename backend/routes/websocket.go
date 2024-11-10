package routes

//реализация websocket подключения, создание маршрута, всякие проверки подключения итд

import (
    "net/http"
    "github.com/gorilla/websocket"
    "github.com/gin-gonic/gin"

    "github.com/kusneid/Ginol/user"
)

var clients = make(map[*websocket.Conn]bool) 
var broadcast = make(chan user.Message) 

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func HandleWebSocket(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	clients[ws] = true
    var msg user.Message
	for {
		err := ws.ReadJSON(&msg)
		if err != nil {
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}
}

func HandleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
	}
}
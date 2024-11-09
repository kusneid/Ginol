package routes

//реализация websocket подключения, создание маршрута, всякие проверки подключения итд

import (
    "net/http"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        // Позволяет подключение с любых источников для тестирования
        return true
    },
}
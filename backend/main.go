package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // главный инстанс клиента

	r.Static("frontend/assets", "./frontend/assets") //css and images

	r.LoadHTMLGlob("frontend/templates/*") // hrml templates

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil) // запуск страницы с выбором входа или регистрации
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil) // вход
	})

	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil) // регистрация
	})

	r.Run(":8080")
}

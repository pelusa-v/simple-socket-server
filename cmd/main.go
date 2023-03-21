package main

import (
	"socket-server/internal/chat"
	"socket-server/internal/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()
	go chat.Manager.ListenActions()
	// engine.LoadHTMLFiles("internal/templates/index.html")
	engine.LoadHTMLFiles("internal/templates/broadcast.html")
	engine.GET("/", controllers.IndexTemplateHandler)
	engine.GET("/ws", controllers.RegisterChatClientHandler)
	// engine.POST("/chat", nil)
	engine.GET("/ping", controllers.PingHandler)
	engine.GET("/wssimple", controllers.ChatbotHandler)
	engine.Run("localhost:8000")
}

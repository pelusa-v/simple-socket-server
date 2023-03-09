package main

import (
	"socket-server/internal/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.LoadHTMLFiles("internal/templates/index.html")
	engine.GET("/", controllers.IndexTemplateHandler)
	engine.GET("/ping", controllers.PingHandler)
	engine.GET("/wssimple", controllers.WebSocketHandler)
	engine.Run("0.0.0.0:8000")
}

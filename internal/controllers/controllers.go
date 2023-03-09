package controllers

import (
	"socket-server/internal/sockets"

	"github.com/gin-gonic/gin"
)

func PingHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "pong"})
}

func WebSocketHandler(ctx *gin.Context) {
	sockets.WebSocketHandler(ctx.Writer, ctx.Request)
}

func IndexTemplateHandler(ctx *gin.Context) {
	ctx.HTML(200, "index.html", nil)
}

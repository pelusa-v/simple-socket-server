package controllers

import (
	"socket-server/internal/sockets"

	"github.com/gin-gonic/gin"
)

func PingHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "pong"})
}

func ChatbotHandler(ctx *gin.Context) {
	sockets.ChatBotSocketHandler(ctx.Writer, ctx.Request)
}

func IndexTemplateHandler(ctx *gin.Context) {
	ctx.HTML(200, "broadcast.html", nil)
}

type ClientDTO struct {
	Name string
}

func RegisterChatClientHandler(ctx *gin.Context) {
	sockets.RegisterClientSocketHandler(ctx.Writer, ctx.Request)
}

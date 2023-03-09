package sockets

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Solve cross-domain problems
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("%v", err)
	}

	for {
		mtype, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Printf("Mensaje recibido vía socket: %s\n", string(msg))
		textResponse := ""
		if strings.ToLower(string(msg)) == "hola" {
			textResponse = "Hola, ¿Cómo estás?"
		} else if strings.ToLower(string(msg)) == "bien" || strings.ToLower(string(msg)) == "mal" {
			textResponse = "Gracias por comunicarte, ¿Qué deseas?"
		} else if strings.ToLower(string(msg)) == "productos" {
			time.Sleep(5 * time.Second)
			textResponse = "Productos:\n - Zanahoria\n - Queso\n - Tomate\n"
		} else {
			textResponse = "No podemos ayudarte con esta solicitud"
		}
		conn.WriteMessage(mtype, []byte(textResponse))
	}
}

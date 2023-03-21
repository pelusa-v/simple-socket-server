package chat

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type IClient interface {
	readws()
	writews()
}

type Client struct {
	UserID  string
	Socket  *websocket.Conn
	Message chan []byte
}

type ClientMessage struct {
	//Message Struct
	Sender string `json:"sender,omitempty"`
	// Recipient string `json:"recipient,omitempty"`
	Content string `json:"content,omitempty"`
	// ServerIP  string `json:"serverIp,omitempty"`
	// SenderIP  string `json:"senderIp,omitempty"`
}

func (c *Client) Read() {
	for {
		_, message, _ := c.Socket.ReadMessage()
		jsonMessage, _ := json.Marshal(&ClientMessage{Sender: c.UserID, Content: string(message)})
		Manager.Broadcast <- jsonMessage
	}
}

func (c *Client) Write() {
	for {
		select {
		case msg := <-c.Message:
			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}
}

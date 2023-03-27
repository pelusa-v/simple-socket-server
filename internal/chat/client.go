package chat

import (
	"encoding/json"
	"fmt"

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
	//Message Struct to catch message from client (for example browser)
	Sender      string `json:"sender,omitempty"`
	Destination string `json:"destination,omitempty"`
	Content     string `json:"content,omitempty"`
	// ServerIP  string `json:"serverIp,omitempty"`
	// SenderIP  string `json:"senderIp,omitempty"`
}

func (c *Client) ReadFromClient() {
	// Read message from socket connectd to client (browser) and broadcast to all clients
	for {
		clientMessage := ClientMessage{}
		_, messageBytes, _ := c.Socket.ReadMessage()
		json.Unmarshal(messageBytes, &clientMessage)
		clientMessage.Sender = c.UserID
		fmt.Printf("READ FROM CLIENT: %v\n", clientMessage)
		// jsonMessage, _ := json.Marshal(&ClientMessage{Sender: c.UserID, Content: string(message)})
		// Manager.Broadcast <- jsonMessage
		Manager.Sender <- &clientMessage
	}
}

func (c *Client) WriteToClient() {
	// When manager broadcast messages scanned by ReadFromClient function
	// (pass value by the message channel), this function send messages to all clietn sockets
	for {
		select {
		case msg := <-c.Message:
			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
		}
	}
}

package chat

import (
	"fmt"

	"github.com/goccy/go-json"
)

type ClientManager struct {
	ClientsStatus map[*Client]bool
	Register      chan *Client
	Unregister    chan *Client
	Broadcast     chan []byte
	Sender        chan *ClientMessage
}

var Manager = ClientManager{
	ClientsStatus: make(map[*Client]bool),
	Register:      make(chan *Client),
	Unregister:    make(chan *Client),
	Broadcast:     make(chan []byte),
	Sender:        make(chan *ClientMessage),
}

func (manager *ClientManager) ListenActions() {
	for {
		select {
		case channel := <-manager.Register:
			manager.ClientsStatus[channel] = true
			msg, _ := json.Marshal(&ClientMessage{Sender: channel.UserID, Content: "Nuevo usuario registrado"})
			manager.BroadcastMessage(msg)
		case channel := <-manager.Unregister:
			fmt.Printf("Se ha cerrado la conexión del siguiente cliente!\n")
			fmt.Printf("UserID: %v", (*channel).UserID)
			fmt.Printf("Message: %v", <-channel.Message)
		case channel := <-manager.Broadcast:
			manager.BroadcastMessage(channel)
		case channel := <-manager.Sender:
			manager.SendMessage(channel)
		}
	}
}

func (manager *ClientManager) BroadcastMessage(message []byte) {
	for client := range manager.ClientsStatus {
		client.Message <- message
	}
}

func (manager *ClientManager) SendMessage(clientMessage *ClientMessage) {
	for client, status := range manager.ClientsStatus {
		if client.UserID == clientMessage.Destination && status {
			client.Message <- []byte(clientMessage.Content)
		}
	}
}

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
}

var Manager = ClientManager{
	ClientsStatus: make(map[*Client]bool),
	Register:      make(chan *Client),
	Unregister:    make(chan *Client),
	Broadcast:     make(chan []byte),
}

func (manager *ClientManager) ListenActions() {
	for {
		select {
		case channel := <-manager.Register:
			manager.ClientsStatus[channel] = true
			msg, _ := json.Marshal(&ClientMessage{Sender: channel.UserID, Content: "Nuevo usuario registrado"})
			manager.BroadcastMessage(msg)
			// fmt.Printf("Se ha registrado un nuevo cliente!\n")
			// fmt.Println(*channel)
			// fmt.Printf("UserID: %v\n", (*channel).UserID)
			// fmt.Printf("Message: %v\n", <-channel.Message)
		case channel := <-manager.Unregister:
			fmt.Printf("Se ha cerrado la conexión del siguiente cliente!\n")
			fmt.Printf("UserID: %v", (*channel).UserID)
			fmt.Printf("Message: %v", <-channel.Message)
		case channel := <-manager.Broadcast:
			for client := range manager.ClientsStatus {
				client.Message <- channel
			}
		}
	}
}

func (manager *ClientManager) BroadcastMessage(message []byte) {
	for client := range manager.ClientsStatus {
		client.Message <- message
	}
}

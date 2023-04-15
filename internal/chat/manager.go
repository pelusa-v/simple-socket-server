package chat

import (
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
			if status := manager.ClientsStatus[channel]; status {
				close(channel.Message)
				delete(Manager.ClientsStatus, channel)
				msg, _ := json.Marshal(&ClientMessage{Sender: channel.UserID, Content: "Usuario desconectado"})
				manager.BroadcastMessage(msg)
			}
		case channel := <-manager.Broadcast:
			manager.BroadcastMessage(channel)
		case channelMessage := <-manager.Sender:
			manager.SendMessage(channelMessage)
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
			// client.Message <- []byte(clientMessage.Content)
			clientMessageBytes, _ := json.Marshal(clientMessage)
			client.Message <- clientMessageBytes
		}
	}
}

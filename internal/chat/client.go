package chat

import "github.com/gorilla/websocket"

type IClient interface {
	readws()
	writews()
}

type Client struct {
	UserID  int
	socket  *websocket.Conn
	message chan []byte
}

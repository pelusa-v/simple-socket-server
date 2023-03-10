package chat

type ClientManager struct {
	clientsStatus map[*Client]bool
	broadcast     chan []byte
	//Newly created long connection client
	register chan *Client
	//Newly canceled long connection client
	unregister chan *Client
}

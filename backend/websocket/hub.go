package ws

import (
	"encoding/json"
	"log"
	"sync"

	"golang.org/x/net/websocket"
)

type Client struct {
	Conn *websocket.Conn
}

var (
	clients = make(map[*Client]bool)
	mu      sync.Mutex
)

func HandleWS(ws *websocket.Conn) {
	client := &Client{Conn: ws}

	mu.Lock()
	clients[client] = true
	mu.Unlock()

	log.Println("🔌 Desktop client connected")

	defer func() {
		mu.Lock()
		delete(clients, client)
		mu.Unlock()
		ws.Close()
		log.Println("❌ Desktop client disconnected")
	}()

	// Keep connection alive
	for {
		var msg string
		if err := websocket.Message.Receive(ws, &msg); err != nil {
			break
		}
	}
}

func Broadcast(payload any) {
	data, _ := json.Marshal(payload)

	mu.Lock()
	defer mu.Unlock()

	for client := range clients {
		websocket.Message.Send(client.Conn, string(data))
	}
}

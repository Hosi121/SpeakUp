package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

const port = "3001"

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins
	},
}

type client struct {
	conn *websocket.Conn
	mu   sync.Mutex
}

var clients = make(map[*client]bool)
var clientsMu sync.Mutex

func main() {
	http.HandleFunc("/", handleWebSocket)
	log.Printf("WebSocket server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	log.Println("-- WebSocket connected --")

	c := &client{conn: conn}
	clientsMu.Lock()
	clients[c] = true
	clientsMu.Unlock()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		broadcast(message, c)
	}

	clientsMu.Lock()
	delete(clients, c)
	clientsMu.Unlock()
}

func broadcast(message []byte, sender *client) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	for c := range clients {
		if c == sender {
			log.Println("- skip sender -")
			continue
		}
		c.mu.Lock()
		err := c.conn.WriteMessage(websocket.TextMessage, message)
		c.mu.Unlock()
		if err != nil {
			log.Println("Error broadcasting message:", err)
		}
	}
}

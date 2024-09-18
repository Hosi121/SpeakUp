package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 注意: 本番環境ではオリジンをチェックすべきです
	},
}

type Client struct {
	conn *websocket.Conn
	send chan []byte
}

type Message struct {
	Type      string          `json:"type"`
	OfferId   int             `json:"offerId,omitempty"`
	Offer     json.RawMessage `json:"offer,omitempty"`
	Answer    json.RawMessage `json:"answer,omitempty"`
	SDP       string          `json:"sdp,omitempty"`
	Candidate json.RawMessage `json:"candidate,omitempty"`
}

var (
	clients    = make(map[*Client]bool)
	broadcast  = make(chan []byte)
	register   = make(chan *Client)
	unregister = make(chan *Client)
	mutex      = &sync.Mutex{}
)

func main() {
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()

	log.Println("Server starting on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{conn: conn, send: make(chan []byte, 256)}
	register <- client

	go readPump(client)
	go writePump(client)
}

func readPump(client *Client) {
	defer func() {
		unregister <- client
		client.conn.Close()
	}()

	for {
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		handleMessage(message, client)
	}
}

func writePump(client *Client) {
	defer client.conn.Close()

	for {
		select {
		case message, ok := <-client.send:
			if !ok {
				client.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := client.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}

func handleMessage(message []byte, sender *Client) {
	var msg Message
	err := json.Unmarshal(message, &msg)
	if err != nil {
		log.Println("Error unmarshaling message:", err)
		return
	}

	switch msg.Type {
	case "offer", "answer", "ice_candidate":
		// オファー、アンサー、ICE candidateは他のすべてのクライアントに転送
		mutex.Lock()
		for client := range clients {
			if client != sender {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(clients, client)
				}
			}
		}
		mutex.Unlock()
	case "end_call":
		// 通話終了メッセージは他のすべてのクライアントに転送
		mutex.Lock()
		for client := range clients {
			if client != sender {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(clients, client)
				}
			}
		}
		mutex.Unlock()
	default:
		log.Println("Unknown message type:", msg.Type)
	}
}

func handleMessages() {
	for {
		select {
		case client := <-register:
			mutex.Lock()
			clients[client] = true
			mutex.Unlock()
		case client := <-unregister:
			mutex.Lock()
			if _, ok := clients[client]; ok {
				delete(clients, client)
				close(client.send)
			}
			mutex.Unlock()
		case message := <-broadcast:
			mutex.Lock()
			for client := range clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(clients, client)
				}
			}
			mutex.Unlock()
		}
	}
}

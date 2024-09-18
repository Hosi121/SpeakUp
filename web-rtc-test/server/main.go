package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 本番環境では適切なオリジン確認を行ってください
	},
}

type Message struct {
	Type      string          `json:"type"`
	Offer     json.RawMessage `json:"offer,omitempty"`
	Answer    json.RawMessage `json:"answer,omitempty"`
	Candidate json.RawMessage `json:"candidate,omitempty"`
}

type UserInfoMessage struct {
	HashedId int `json:"hashedId"`
}

// mock
var matchings map[int]int = map[int]int{
	1: 2,
	2: 1,
}

var wsToId map[*websocket.Conn]int = map[*websocket.Conn]int{}
var idToWs map[int]*websocket.Conn = map[int]*websocket.Conn{}

var clients = make(map[*websocket.Conn]bool)

func main() {
	http.HandleFunc("/ws", handleConnections)

	log.Println("Server starting on :8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	clients[ws] = true

	var msg UserInfoMessage
	ws.ReadJSON(&msg)
	wsToId[ws] = msg.HashedId
	fmt.Printf("connect: hashedId=%d\n", msg.HashedId)

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}

		to := matchings[wsToId[ws]]
		sendMessage(msg, to)
	}
}

func broadcastMessage(msg Message, sender *websocket.Conn) {
	for client := range clients {
		if client != sender {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func sendMessage(msg Message, to int) {
	client := idToWs[to]
	err := client.WriteJSON(msg)
	if err != nil {
		log.Printf("error: %v", err)
		client.Close()
		delete(clients, client)
	}
}

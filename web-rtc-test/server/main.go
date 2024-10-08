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

	port := "8083"
	log.Println("Server starting on :" + port)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleConnections")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	clients[ws] = true

	var msg UserInfoMessage
	ws.ReadJSON(&msg)
	fmt.Printf("msg=%v\n", msg)
	wsToId[ws] = msg.HashedId
	idToWs[msg.HashedId] = ws
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
		fmt.Printf("%d->%d\n", wsToId[ws], to)
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
	fmt.Printf("%v\n", client)
	err := client.WriteJSON(msg)
	if err != nil {
		log.Printf("error: %v", err)
		client.Close()
		delete(clients, client)
	}
}

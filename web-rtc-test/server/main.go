package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"goserver2/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// WebSocketのアップグレード用
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 本番環境では適切なオリジン確認を行ってください
	},
}

// メッセージ用の構造体
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
var matchings = map[int]int{
	1: 2,
	2: 1,
}

var wsToId = map[*websocket.Conn]int{}
var idToWs = map[int]*websocket.Conn{}

var clients = make(map[*websocket.Conn]bool)

func main() {
	router := gin.Default()

	router.Use(middlewares.JWTAuthMiddleware())
	// WebSocket接続用のエンドポイント
	router.GET("/ws", handleConnections)

	log.Println("Server starting on :8082")
	err := router.Run(":8082")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// WebSocket接続処理
func handleConnections(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer ws.Close()

	clients[ws] = true

	var msg UserInfoMessage
	if err := ws.ReadJSON(&msg); err != nil {
		log.Printf("error reading user info: %v", err)
		return
	}
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

// クライアントにメッセージを送信
func sendMessage(msg Message, to int) {
	client := idToWs[to]
	if client == nil {
		log.Printf("error: no client with id %d", to)
		return
	}

	err := client.WriteJSON(msg)
	if err != nil {
		log.Printf("error: %v", err)
		client.Close()
		delete(clients, client)
	}
}

package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Hosi121/SpeakUp/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// WebSocketのアップグレード用
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// メッセージ用の構造体
type Message struct {
	Type      string          `json:"type"`
	Offer     json.RawMessage `json:"offer,omitempty"`
	Answer    json.RawMessage `json:"answer,omitempty"`
	Candidate json.RawMessage `json:"candidate,omitempty"`
	Token     string          `json:"token,omitempty"`
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

// WebSocket接続処理
func SignalingController(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer ws.Close()

	clients[ws] = true

	count := 0
	for {
		count++
		fmt.Printf("count=%d\n", count)
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}

		if msg.Type == "Authorization" {
			token := msg.Token[7:]
			userIDStr, err := utils.ValidateJWT(token)
			var response Message
			if err == nil {
				response = Message{Type: "AuthResult", Token: "Authentication successful"}
			} else {
				response = Message{Type: "AuthResult", Token: "Authentication failed"}
			}
			userId, _ := strconv.Atoi(userIDStr)
			wsToId[ws] = userId
			idToWs[userId] = ws
			fmt.Printf("connect: hashedId=%d\n", userId)

			if err := ws.WriteJSON(response); err != nil {
				fmt.Printf("Error sending response: %v\n", err)
				return
			}
		} else {
			to := matchings[wsToId[ws]]
			sendMessage(msg, to)
		}
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

func handleHoge(c *gin.Context) {
	c.Writer.Write([]byte("hoge"))
}

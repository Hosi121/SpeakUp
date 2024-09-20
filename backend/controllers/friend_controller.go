package controllers

import (
	"log"
	"net/http"

	"github.com/Hosi121/SpeakUp/ent"
	"github.com/gin-gonic/gin"
)

type Friend struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type FriendRecord struct {
	UserID       int    `json:"user_id"`
	TargetUserID int    `json:"target_user_id"`
	Status       string `json:"status"`
}

// モックデータ
var friends = []Friend{
	{ID: 1, Username: "Alice", Avatar: "https://example.com/avatar1.png"},
	{ID: 2, Username: "Bob", Avatar: "https://example.com/avatar2.png"},
	{ID: 3, Username: "Charlie", Avatar: "https://example.com/avatar3.png"},
}

func GetFriendList() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"friends": friends})
	}
}

func GetFriendByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		friendname := c.Param("friendname")

		for _, friend := range friends {
			if friend.Username == friendname {
				c.JSON(http.StatusOK, friend)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Friend not found"})
	}
}

func AddFriend() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("AddFriend handler called")

		// JWT から取得したユーザーIDを使用
		userIDInterface, exists := c.Get("user_id")
		if !exists {
			log.Println("User ID not found in context")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			return
		}
		userID, ok := userIDInterface.(int)
		if !ok {
			log.Printf("Invalid user ID type: %T", userIDInterface)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
			return
		}

		log.Printf("Authenticated user ID: %d", userID)

		// ターゲットユーザーIDをリクエストボディから取得
		var req struct {
			TargetUserID int `json:"target_user_id" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Printf("Failed to bind JSON: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		log.Printf("Target user ID: %d", req.TargetUserID)

		// データベースクライアントを取得
		clientInterface, exists := c.Get("db")
		if !exists {
			log.Println("Database client not found in context")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database client is not available"})
			return
		}
		client, ok := clientInterface.(*ent.Client)
		if !ok {
			log.Println("Invalid database client type")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid database client"})
			return
		}

		// 新しいフレンド関係を作成
		friend, err := client.FRIENDS.Create().
			SetUserID(userID).
			SetTargetUserID(req.TargetUserID).
			SetStatus("FRIEND").
			Save(c.Request.Context())

		if err != nil {
			log.Printf("Failed to add friend: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add friend: " + err.Error()})
			return
		}

		log.Printf("Friend added successfully: %+v", friend)

		c.JSON(http.StatusOK, gin.H{
			"message": "Friend added successfully",
			"friend":  friend,
		})
	}
}

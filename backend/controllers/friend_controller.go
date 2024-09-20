package controllers

import (
	"net/http"
	"strconv"

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
		// ユーザーIDとターゲットユーザーIDを取得
		userID, err := strconv.Atoi(c.PostForm("user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		targetUserID, err := strconv.Atoi(c.PostForm("target_user_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid target user ID"})
			return
		}

		client := ent.FromContext(c)

		// 新しいフレンド関係を作成
		friend, err := client.FRIENDS.Create().
			SetUserID(userID).
			SetTargetUserID(targetUserID).
			SetStatus("FRIEND"). // 強制的に"FRIEND"で登録
			Save(c.Request.Context())

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add friend"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Friend added successfully",
			"friend":  friend,
		})
	}
}

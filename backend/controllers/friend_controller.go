package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Friend struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
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

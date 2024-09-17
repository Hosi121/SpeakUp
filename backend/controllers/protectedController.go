package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProtectedEndpoint(c *gin.Context) {
	// ミドルウェアで設定した user_id を取得
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	// エンドポイントの処理
	c.JSON(http.StatusOK, gin.H{
		"message": "This is a protected endpoint",
		"user_id": userID,
	})
}

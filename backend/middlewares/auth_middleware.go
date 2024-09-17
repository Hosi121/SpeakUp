package middleware

import (
	"net/http"
	"strings"

	"github.com/Hosi121/SpeakUp/utils"

	"github.com/gin-gonic/gin"
)

// JWT認証ミドルウェア
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authorizationヘッダーからトークンを取得
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// "Bearer " プレフィックスを取り除く
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// トークンの検証
		userID, err := utils.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// ユーザーIDをコンテキストに設定
		c.Set("user_id", userID)

		// 次のハンドラーへ進む
		c.Next()
	}
}

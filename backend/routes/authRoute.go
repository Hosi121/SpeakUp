package routes

import (
	"github.com/Hosi121/SpeakUp/controllers"
	"github.com/gin-gonic/gin"
)

func SupabaseAuthRoutes(r *gin.Engine) {
	// サインアップとサインインルートの定義
	r.POST("/signup", controllers.SignUp)
	r.POST("/signin", controllers.SignIn)
}

func ProtectedRoutes(r *gin.RouterGroup) {
	// 認証が必要なルートの定義
	r.GET("/protected_endpoint", controllers.ProtectedEndpoint)
	// 他の認証が必要なエンドポイントをここに追加
}

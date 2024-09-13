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

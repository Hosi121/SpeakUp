package routes

import (
	"github.com/Hosi121/SpeakUp/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/auth", controllers.Auth())
}

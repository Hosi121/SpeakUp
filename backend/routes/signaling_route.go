package routes

import (
	"github.com/Hosi121/SpeakUp/controllers"
	"github.com/gin-gonic/gin"
)

func SignalingRoutes(router *gin.RouterGroup) {
	router.GET("/ws", controllers.SignalingController)
}

package routes

import (
	"github.com/Hosi121/SpeakUp/controllers"
	"github.com/Hosi121/SpeakUp/ent"
	"github.com/gin-gonic/gin"
)

func ProtectedRoutes(router *gin.RouterGroup, client *ent.Client) {
	// Existing routes
	router.GET("/user/info", controllers.GetUserInfo(client))

	// Add this line
	router.PUT("/user/update", controllers.UpdateUserInfo(client))
}

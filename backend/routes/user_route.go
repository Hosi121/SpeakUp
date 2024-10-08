package routes

import (
	"github.com/Hosi121/SpeakUp/controllers"
	"github.com/Hosi121/SpeakUp/ent"
	"github.com/gin-gonic/gin"
)

func ProtectedRoutes(router *gin.RouterGroup, client *ent.Client) {
	// Existing routes
	router.GET("/user/info", controllers.GetUserInfo(client))
	router.PUT("/user/update", controllers.UpdateUserInfo(client))
	router.PUT("/user/avatar", controllers.UpdateAvatar(client))
	router.GET("/users/:userId/avatar", controllers.GetUserAvatar(client))
	router.GET("/users/search", controllers.SearchUsers(client))
	router.GET("/users/search/id/:id", controllers.SearchUserByID(client))

	// Memo routes
	router.GET("/memo", controllers.GetMemo(client))
	router.PUT("/memo", controllers.UpdateMemo(client))

	router.GET("/friend/me", controllers.GetFriendList())
	router.GET("/friend/:friendname", controllers.GetFriendByName())
	router.POST("/friend/register", controllers.AddFriend())
	router.POST("/events", controllers.CreateEvent(client))
	router.GET("/events", controllers.GetEvents(client))

}

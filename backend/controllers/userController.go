package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Hosi121/SpeakUp/ent"
	"github.com/Hosi121/SpeakUp/ent/users"
	"github.com/gin-gonic/gin"
)

// UserResponse defines the user data to be returned
type UserResponse struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	AvatarURL string    `json:"avatar_url"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetUserInfo handles the GET /user/info endpoint
func GetUserInfo(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve user_id from context
		userIDInterface, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
			return
		}

		userID, ok := userIDInterface.(int)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID is invalid"})
			return
		}

		// Query the user from the database
		user, err := client.USERS.
			Query().
			Where(users.IDEQ(userID)).
			Only(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query user", "details": err.Error()})
			return
		}

		// Construct the response
		response := UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			AvatarURL: "https://example.com/avatar.jpg",
			Role:      "user",
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}

		// Return the user information as JSON
		c.JSON(http.StatusOK, response)
	}
}

package controllers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
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

type UpdateUserRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
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
			AvatarURL: user.AvatarURL,
			Role:      "user",
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}

		// Return the user information as JSON
		c.JSON(http.StatusOK, response)
	}
}

// UpdateUserInfo handles the PUT /user/update endpoint
func UpdateUserInfo(client *ent.Client) gin.HandlerFunc {
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

		var req UpdateUserRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// Build the update query
		updateQuery := client.USERS.UpdateOneID(userID)
		if req.Username != "" {
			updateQuery.SetUsername(req.Username)
		}
		if req.Email != "" {
			updateQuery.SetEmail(req.Email)
		}

		// Execute the update
		_, err := updateQuery.Save(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user", "details": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
	}
}

// Define maximum upload size (e.g., 2MB)
const MaxUploadSize = 2 * 1024 * 1024

// UpdateAvatar handles the PUT /user/avatar endpoint for uploading user avatars
func UpdateAvatar(client *ent.Client) gin.HandlerFunc {
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

		// Set maximum upload size
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxUploadSize)

		// Parse the uploaded file
		file, header, err := c.Request.FormFile("avatar")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload file", "details": err.Error()})
			return
		}
		defer file.Close()

		// Ensure upload directory exists
		uploadDir := "./backend/upload"
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
			return
		}

		// Create unique filename
		filename := fmt.Sprintf("avatar_%d%s", userID, filepath.Ext(header.Filename))
		filepath := filepath.Join(uploadDir, filename)

		// Save the file to the upload directory
		out, err := os.Create(filepath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file", "details": err.Error()})
			return
		}
		defer out.Close()

		if _, err := file.Seek(0, 0); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reset file pointer"})
			return
		}

		if _, err := out.ReadFrom(file); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file data", "details": err.Error()})
			return
		}

		// Update user's avatar URL in the database
		avatarURL := "/upload/" + filename
		_, err = client.USERS.UpdateOneID(userID).SetAvatarURL(avatarURL).Save(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user avatar URL", "details": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":    "Avatar uploaded successfully",
			"avatar_url": "http://localhost:8081/upload/" + filename, // 画像のURLを返す
		})
	}
}
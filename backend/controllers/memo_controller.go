package controllers

import (
	"net/http"
	"strconv"

	"github.com/Hosi121/SpeakUp/ent"
	"github.com/Hosi121/SpeakUp/ent/memos"
	"github.com/gin-gonic/gin"
)

type MemoResponse struct {
	Memo1 string `json:"memo1"`
	Memo2 string `json:"memo2"`
}

// GetMemo handles the GET /memo endpoint
func GetMemo(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve user_id from context
		userIDInterface, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			return
		}

		var userID int
		switch v := userIDInterface.(type) {
		case int:
			userID = v
		case string:
			var err error
			userID, err = strconv.Atoi(v)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID conversion error"})
				return
			}
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
			return
		}

		ctx := c.Request.Context()

		// Query the memo from the database
		m, err := client.MEMOS.
			Query().
			Where(memos.UserIDEQ(userID)).
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				// Return empty memos if not found
				c.JSON(http.StatusOK, MemoResponse{Memo1: "", Memo2: ""})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query memo", "details": err.Error()})
			return
		}

		// Return the memo content
		c.JSON(http.StatusOK, MemoResponse{Memo1: m.Memo1, Memo2: m.Memo2})
	}
}

// UpdateMemo handles the PUT /memo endpoint
func UpdateMemo(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve user_id from context
		userIDInterface, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			return
		}

		var userID int
		switch v := userIDInterface.(type) {
		case int:
			userID = v
		case string:
			var err error
			userID, err = strconv.Atoi(v)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID conversion error"})
				return
			}
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
			return
		}

		var req MemoResponse
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		ctx := c.Request.Context()

		// Check if memo exists
		m, err := client.MEMOS.
			Query().
			Where(memos.UserIDEQ(userID)).
			Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				// Create a new memo
				_, err = client.MEMOS.
					Create().
					SetUserID(userID).
					SetMemo1(req.Memo1).
					SetMemo2(req.Memo2).
					Save(ctx)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create memo", "details": err.Error()})
					return
				}
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query memo", "details": err.Error()})
				return
			}
		} else {
			// Update existing memo
			_, err = m.Update().
				SetMemo1(req.Memo1).
				SetMemo2(req.Memo2).
				Save(ctx)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update memo", "details": err.Error()})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"message": "Memo updated successfully"})
	}
}

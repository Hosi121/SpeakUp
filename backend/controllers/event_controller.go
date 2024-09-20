package controllers

import (
	"net/http"
	"time"

	"github.com/Hosi121/SpeakUp/ent" // あなたのプロジェクトの実際のパスに置き換えてください
	"github.com/gin-gonic/gin"
)

type EventInput struct {
	DateTime string   `json:"dateTime"`
	Theme    string   `json:"theme"`
	Topics   []string `json:"topics"`
}

func CreateEvent(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input EventInput

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 複数のフォーマットを試行
		layouts := []string{
			time.RFC3339,
			"2006-01-02T15:04:05Z07:00",
			"2006-01-02T15:04:05",
			"2006-01-02 15:04:05",
		}

		var startTime time.Time
		var err error
		for _, layout := range layouts {
			startTime, err = time.Parse(layout, input.DateTime)
			if err == nil {
				break
			}
		}

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date time format"})
			return
		}

		endTime := startTime.Add(30 * time.Minute)

		// トランザクションを開始
		tx, err := client.Tx(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
			return
		}

		// AI_THEMESを作成
		theme, err := tx.AI_THEMES.Create().
			SetThemeText(input.Theme).
			SetTopic1(input.Topics[0]).
			SetTopic2(input.Topics[1]).
			SetTopic3(input.Topics[2]).
			Save(c)

		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create theme"})
			return
		}

		// EVENTを作成
		event, err := tx.EVENTS.Create().
			SetEventStart(startTime).
			SetEventEnd(endTime).
			SetThemeID(theme.ID).
			Save(c)

		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
			return
		}

		// トランザクションをコミット
		if err := tx.Commit(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Event created successfully",
			"event":   event,
			"theme":   theme,
		})
	}
}

package controllers

import (
	"net/http"
	"time"

	"github.com/Hosi121/SpeakUp/ent"
	"github.com/Hosi121/SpeakUp/ent/events"
	"github.com/gin-gonic/gin"
)

type EventInput struct {
	EventStart string   `json:"event_start"`
	DateTime   string   `json:"dateTime"`
	Theme      string   `json:"theme"`
	Topics     []string `json:"topics"`
}

func CreateEvent(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input EventInput

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		startValue := input.EventStart
		if startValue == "" {
			startValue = input.DateTime
		}
		if startValue == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "event_start is required"})
			return
		}

		// 複数のフォーマットを試行
		layouts := []string{
			time.RFC3339,
			time.RFC3339Nano,
			"2006-01-02T15:04:05Z07:00",
			"2006-01-02T15:04:05",
			"2006-01-02 15:04:05",
		}

		var startTime time.Time
		var err error
		for _, layout := range layouts {
			startTime, err = time.Parse(layout, startValue)
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

		var topic1, topic2, topic3 string
		if len(input.Topics) > 0 {
			topic1 = input.Topics[0]
		}
		if len(input.Topics) > 1 {
			topic2 = input.Topics[1]
		}
		if len(input.Topics) > 2 {
			topic3 = input.Topics[2]
		}

		// AI_THEMESを作成
		theme, err := tx.AI_THEMES.Create().
			SetThemeText(input.Theme).
			SetTopic1(topic1).
			SetTopic2(topic2).
			SetTopic3(topic3).
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
			"id":          event.ID,
			"event_start": event.EventStart.Format(time.RFC3339),
			"event_end":   event.EventEnd.Format(time.RFC3339),
			"theme_id":    theme.ID,
			"theme": gin.H{
				"theme_text": theme.ThemeText,
				"topic1":     theme.Topic1,
				"topic2":     theme.Topic2,
				"topic3":     theme.Topic3,
			},
		})
	}
}

func GetEvents(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// イベントを取得し、関連するテーマ情報も含める
		events, err := client.EVENTS.Query().
			WithUses().                              // テーマ情報を含める
			Order(ent.Desc(events.FieldEventStart)). // 開始時間の降順でソート
			All(c)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch events"})
			return
		}

		// レスポンス用のスライスを準備
		var response []gin.H

		for _, event := range events {
			theme := event.Edges.Uses
			if theme == nil {
				// テーマが見つからない場合はスキップ
				continue
			}

			response = append(response, gin.H{
				"id":          event.ID,
				"event_start": event.EventStart.Format(time.RFC3339),
				"event_end":   event.EventEnd.Format(time.RFC3339),
				"theme_id":    event.ThemeID,
				"theme": gin.H{
					"theme_text": theme.ThemeText,
					"topic1":     theme.Topic1,
					"topic2":     theme.Topic2,
					"topic3":     theme.Topic3,
				},
			})
		}

		c.JSON(http.StatusOK, response)
	}
}

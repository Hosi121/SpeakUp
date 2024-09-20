package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func createEvent(c *gin.Context) {
	var input struct {
		DateTime string   `json:"dateTime"`
		Theme    string   `json:"theme"`
		Topics   []string `json:"topics"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startTime, err := time.Parse(time.RFC3339, input.DateTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date time format"})
		return
	}

	endTime := startTime.Add(30 * time.Minute)

	// Create AI_THEMES
	theme, err := client.AI_THEMES.Create().
		SetThemeText(input.Theme).
		SetTopic1(input.Topics[0]).
		SetTopic2(input.Topics[1]).
		SetTopic3(input.Topics[2]).
		Save(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create theme"})
		return
	}

	// Create EVENT
	event, err := client.EVENTS.Create().
		SetEventStart(startTime).
		SetEventEnd(endTime).
		SetThemeID(theme.ID).
		Save(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event created successfully",
		"event":   event,
		"theme":   theme,
	})
}

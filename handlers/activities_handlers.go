package handlers

import (
	db "course-chart/config"
	"course-chart/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetActivities(c *gin.Context) {
	var activities []models.Activity
	err := db.Conn.Find(&activities).Where("activities.custom = FALSE").Error
	if err != nil {
		RenderError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Activities found",
		"data":    activities,
	})
}

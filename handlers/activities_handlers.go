package handlers

import (
	db "course-chart/config"
	"course-chart/models"

	"github.com/gin-gonic/gin"
)

func GetActivities(c *gin.Context) {
	var activities []models.Activity
	err := db.Conn.Find(&activities).Where("activities.custom = FALSE").Error
	if err != nil {
		renderNotFound(c, err)
		return
	}

	renderFound(c, activities, "Activities found")
}

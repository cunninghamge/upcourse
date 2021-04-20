package handlers

import (
	db "course-chart/config"
	"course-chart/models"

	"github.com/gin-gonic/gin"
)

func GetActivities(c *gin.Context) {
	var activities []models.Activity
	err := db.Conn.Where("activities.custom = FALSE").Find(&activities).Error
	if err != nil {
		renderError(c, err)
		return
	}

	renderFound(c, activities, "Activities found")
}

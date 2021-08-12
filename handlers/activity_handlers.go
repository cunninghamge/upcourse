package handlers

import (
	"github.com/gin-gonic/gin"

	db "upcourse/config"
	"upcourse/models"
)

func GetActivities(c *gin.Context) {
	var activities []models.Activity
	err := db.Conn.Where("activities.custom = FALSE").Find(&activities).Error
	if err != nil {
		renderError(c, err)
		return
	}

	var serializedActivities []SerializedResource
	for _, a := range activities {
		serializedActivities = append(serializedActivities, SerializeActivity(a))
	}
	renderFoundRecords(c, serializedActivities)
}

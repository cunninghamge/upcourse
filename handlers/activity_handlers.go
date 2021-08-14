package handlers

import (
	"github.com/gin-gonic/gin"

	db "upcourse/config"
	"upcourse/models"
)

func GetActivities(c *gin.Context) {
	var activities []*models.Activity
	tx := db.Conn.Where("activities.custom = FALSE").Find(&activities)
	if tx.Error != nil {
		renderError(c, tx.Error)
		return
	}

	renderFoundRecords(c, activities)
}

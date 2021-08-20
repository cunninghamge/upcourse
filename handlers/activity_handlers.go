package handlers

import (
	"github.com/gin-gonic/gin"

	"upcourse/models"
)

func GetActivities(c *gin.Context) {
	activities, err := models.GetActivities()
	if err != nil {
		renderError(c, err)
		return
	}

	renderFoundRecords(c, activities)
}

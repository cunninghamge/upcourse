package handlers

import (
	db "course-chart/config"
	"course-chart/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetModule(c *gin.Context) {
	var module models.Module
	err := db.Conn.Preload("ModuleActivities.Activity").First(&module, c.Param("id")).Error

	if err != nil {
		RenderError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Module found",
		"data":    module,
	})
}

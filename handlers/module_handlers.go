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

func CreateModule(c *gin.Context) {
	var input models.Module

	if bindErr := c.ShouldBindJSON(&input); bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  bindErr.Error(),
		})
		return
	}

	err := db.Conn.Create(&input).Error
	if err != nil {
		RenderPostError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Module created successfully",
	})
}

func RenderPostError(c *gin.Context, err error) {
	c.JSON(http.StatusServiceUnavailable, gin.H{
		"status": http.StatusServiceUnavailable,
		"error":  "Unable to create record",
	})
}

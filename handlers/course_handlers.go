package handlers

import (
	db "course-chart/config"
	"course-chart/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCourse(c *gin.Context) {
	var course models.Course
	db.Conn.Preload("Modules.ModuleActivities.Activity").First(&course, c.Param("id"))

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Course found",
		"data":    course,
	})
}

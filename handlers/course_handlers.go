package handlers

import (
	db "course-chart/config"
	"course-chart/models"

	"github.com/gin-gonic/gin"
)

func GetCourse(c *gin.Context) {
	var course models.Course
	db.Conn.First(&course, 1)

	c.String(200, course.Name)
}

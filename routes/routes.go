package routes

import (
	db "course-chart/config"

	"github.com/gin-gonic/gin"
)

type Course struct {
	Id          int
	Name        string
	Institution string
	CreditHours int
	Length      int
	CreatedAt   string
	UpdatedAt   string
}

func GetRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/courses/1", GetCourse)
	return r
}

func GetCourse(c *gin.Context) {
	var course Course
	db.Conn.First(&course, 1)

	c.String(200, course.Name)
}

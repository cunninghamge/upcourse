package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func Routes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.GET("/courses/1", func(c *gin.Context) {

		var course Course
		db.First(&course, 1)

		c.String(200, course.Name)
	})
	return r
}

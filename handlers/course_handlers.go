package handlers

import (
	db "course-chart/config"
	"course-chart/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderError(c *gin.Context, err error) {
	log.Printf("Error retrieving course from database.\nReason: %v", err)
	c.JSON(http.StatusNotFound, gin.H{
		"status": http.StatusNotFound,
		"errors": "Course not found",
	})
}

func GetCourse(c *gin.Context) {
	var course models.Course
	err := db.Conn.First(&course, c.Param("id")).Error

	if err != nil {
		RenderError(c, err)
		return
	}

	var activityTotals []models.ActivityTotals
	db.Conn.Model(&models.Activity{}).
		Joins("JOIN module_activities ON module_activities.activity_id=activities.id").
		Joins("JOIN modules ON modules.id=module_activities.module_id").
		Select("activities.name, activities.id, modules.id AS module_id, modules.name AS module_name, sum(multiplier * module_activities.input) AS minutes").
		Group("activities.id, modules.id, modules.name").
		Where("modules.course_id = ?", c.Param("id")).
		Scan(&activityTotals)

	completeResponse := struct {
		Course         models.Course           `json:"course"`
		ActivityTotals []models.ActivityTotals `json:"activityTotals"`
	}{
		Course:         course,
		ActivityTotals: activityTotals,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Course found",
		"data":    completeResponse,
	})
}

func GetCourses(c *gin.Context) {
	var courses []models.Course
	err := db.Conn.Find(&courses).Error

	if err != nil {
		RenderError(c, err)
		return
	}

	var courseList []models.CourseIdentifier

	for _, course := range courses {
		courseList = append(courseList, models.CourseIdentifier{
			Id:   course.Id,
			Name: course.Name,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Courses found",
		"data":    courseList,
	})
}

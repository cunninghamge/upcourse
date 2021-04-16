package handlers

import (
	db "course-chart/config"
	"course-chart/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderError(c *gin.Context, err error) {
	log.Printf("Error retrieving record from database.\nReason: %v", err)
	c.JSON(http.StatusNotFound, gin.H{
		"status": http.StatusNotFound,
		"errors": "Record not found",
	})
}

func GetCourse(c *gin.Context) {
	var course models.Course
	err := db.Conn.Preload("Modules.ModuleActivities.Activity").First(&course, c.Param("id")).Error

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
	err := db.Conn.Preload("Modules").Select("courses.id, courses.name").Find(&courses).Error

	if err != nil {
		RenderError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Courses found",
		"data":    courses,
	})
}

func CreateCourse(c *gin.Context) {
	var input models.Course

	if bindErr := c.ShouldBindJSON(&input); bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  bindErr.Error(),
		})
		return
	}

	err := db.Conn.Create(&input).Error
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": http.StatusServiceUnavailable,
			"error":  "Unable to create record",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Course created successfully",
	})
}

func UpdateCourse(c *gin.Context) {
	var input models.UpdatableCourse

	if bindErr := c.ShouldBindJSON(&input); bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  bindErr.Error(),
		})
		return
	}

	err := db.Conn.First(&models.Course{}, c.Param("id")).Error
	if err != nil {
		RenderError(c, err)
		return
	}

	course := models.Course{
		Name:        input.Name,
		Institution: input.Institution,
		CreditHours: input.CreditHours,
		Goal:        input.Goal,
	}

	err = db.Conn.Model(&models.Course{}).Where("id = ?", c.Param("id")).Updates(&course).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Course updated successfully",
	})
}

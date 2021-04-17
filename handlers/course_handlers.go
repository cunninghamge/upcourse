package handlers

import (
	db "course-chart/config"
	"course-chart/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetCourse(c *gin.Context) {
	var course models.Course
	err := db.Conn.Preload("Modules.ModuleActivities.Activity").
		First(&course, c.Param("id")).Error

	if err != nil {
		renderNotFound(c, err)
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

	renderFound(c, completeResponse, "Course found")
}

func GetCourses(c *gin.Context) {
	var courses []models.Course
	err := db.Conn.Preload("Modules", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name, course_id")
	}).Select("courses.id, courses.name").Find(&courses).Error

	if err != nil {
		renderNotFound(c, err)
		return
	}

	renderFound(c, courses, "Courses found")
}

func CreateCourse(c *gin.Context) {
	var input models.Course
	if bindErr := c.ShouldBindJSON(&input); bindErr != nil {
		renderBindError(c, bindErr)
		return
	}

	err := db.Conn.Create(&input).Error
	if err != nil {
		renderError(c, err)
		return
	}

	renderCreated(c, "Course created successfully")
}

func UpdateCourse(c *gin.Context) {
	err := db.Conn.First(&models.Course{}, c.Param("id")).Error
	if err != nil {
		renderNotFound(c, err)
		return
	}

	var input models.UpdatableCourse
	if bindErr := c.ShouldBindJSON(&input); bindErr != nil {
		renderBindError(c, bindErr)
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
		renderError(c, err)
		return
	}

	renderSuccess(c, "Course updated successfully")
}

func DeleteCourse(c *gin.Context) {
	err := db.Conn.First(&models.Course{}, c.Param("id")).Error
	if err != nil {
		renderNotFound(c, err)
		return
	}

	err = db.Conn.Model(&models.Course{}).
		Delete(&models.Course{}, c.Param("id")).Error
	if err != nil {
		renderError(c, err)
		return
	}

	renderSuccess(c, "Course deleted successfully")
}

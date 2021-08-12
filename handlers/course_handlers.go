package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	db "upcourse/config"
	"upcourse/models"
)

func GetCourse(c *gin.Context) {
	var course models.Course
	err := db.Conn.Preload("Modules.ModuleActivities.Activity").
		First(&course, c.Param("id")).Error
	if err != nil {
		renderError(c, err)
		return
	}

	serializedCourse := SerializeCourse(course)
	renderFoundRecords(c, serializedCourse)
}

func GetCourses(c *gin.Context) {
	var courses []models.Course
	err := db.Conn.Preload("Modules", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name, course_id")
	}).Select("courses.id, courses.name").Find(&courses).Error
	if err != nil {
		renderError(c, err)
		return
	}

	var serializedCourses []SerializedResource
	for _, c := range courses {
		serializedCourses = append(serializedCourses, SerializeCourse(c))
	}
	renderFoundRecords(c, serializedCourses)
}

func CreateCourse(c *gin.Context) {
	var input models.Course
	if err := c.ShouldBindJSON(&input); err != nil {
		renderError(c, err)
		return
	}

	errs := validateFields(input)
	if len(errs) > 0 {
		renderErrors(c, errs)
		return
	}

	err := db.Conn.Create(&input).Error
	if err != nil {
		renderError(c, err)
		return
	}

	renderSuccess(c, http.StatusCreated)
}

func UpdateCourse(c *gin.Context) {
	err := db.Conn.First(&models.Course{}, c.Param("id")).Error
	if err != nil {
		renderError(c, err)
		return
	}

	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		renderError(c, err)
		return
	}

	err = db.Conn.Model(&models.Course{}).Where("id = ?", c.Param("id")).Updates(&course).Error
	if err != nil {
		renderError(c, err)
		return
	}

	renderSuccess(c, http.StatusOK)
}

func DeleteCourse(c *gin.Context) {
	err := db.Conn.Model(&models.Course{}).
		Delete(&models.Course{}, c.Param("id")).Error
	if err != nil {
		renderError(c, err)
		return
	}

	renderSuccess(c, http.StatusOK)
}

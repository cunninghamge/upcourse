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
	tx := db.Conn.Preload("Modules.ModuleActivities.Activity").
		First(&course, c.Param("id"))
	if tx.Error != nil {
		renderError(c, tx.Error)
		return
	}

	renderFoundRecords(c, &course)
}

func GetCourses(c *gin.Context) {
	var courses []*models.Course
	tx := db.Conn.Preload("Modules", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name, number, course_id")
	}).Select("courses.id, courses.name").Find(&courses)
	if tx.Error != nil {
		renderError(c, tx.Error)
		return
	}

	renderFoundRecords(c, courses)
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

	tx := db.Conn.Create(&input)
	if tx.Error != nil {
		renderError(c, tx.Error)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func UpdateCourse(c *gin.Context) {
	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		renderError(c, err)
		return
	}

	tx := db.Conn.Model(&models.Course{}).First(&models.Course{}, c.Param("id")).Updates(&course)
	if tx.Error != nil {
		renderError(c, tx.Error)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func DeleteCourse(c *gin.Context) {
	tx := db.Conn.Model(&models.Course{}).
		Delete(&models.Course{}, c.Param("id"))
	if tx.Error != nil {
		renderError(c, tx.Error)
		return
	}

	c.JSON(http.StatusOK, nil)
}

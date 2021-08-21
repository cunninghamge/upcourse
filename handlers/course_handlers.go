package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"upcourse/models"
)

const courseSchema = "./schemas/course_schema.json"

func GetCourse(c *gin.Context) {
	course, err := models.GetCourse(c.Param("id"))
	if err != nil {
		renderError(c, err)
		return
	}

	renderFoundRecords(c, course)
}

func GetCourses(c *gin.Context) {
	courses, err := models.GetCourseList()
	if err != nil {
		renderError(c, err)
		return
	}

	renderFoundRecords(c, courses)
}

func CreateCourse(c *gin.Context) {
	jsonData, errs := models.Validate(c, courseSchema)
	if errs != nil {
		renderErrors(c, errs)
		return
	}

	var course models.Course
	if err := json.Unmarshal(jsonData, &course); err != nil {
		renderError(c, err)
		return
	}

	if err := models.CreateCourse(&course); err != nil {
		renderError(c, err)
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

	if err := models.UpdateCourse(course, c.Param("id")); err != nil {
		renderError(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func DeleteCourse(c *gin.Context) {
	if err := models.DeleteCourse(c.Param("id")); err != nil {
		renderError(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

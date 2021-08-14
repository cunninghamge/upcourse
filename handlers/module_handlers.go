package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	db "upcourse/config"
	"upcourse/models"
)

func GetModule(c *gin.Context) {
	var module models.Module
	tx := db.Conn.Preload("ModuleActivities.Activity").First(&module, c.Param("id"))
	if tx.Error != nil {
		renderError(c, tx.Error)
		return
	}

	renderFoundRecords(c, &module)
}

func CreateModule(c *gin.Context) {
	var input models.Module
	if err := c.ShouldBindJSON(&input); err != nil {
		renderError(c, err)
		return
	}

	errs := validateFields(input)
	if len(errs) > 0 {
		renderErrors(c, errs)
		return
	}

	courseId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		renderError(c, errors.New(ErrBadRequest))
		return
	}
	input.CourseId = courseId

	tx := db.Conn.Create(&input)
	if tx.Error != nil {
		renderError(c, tx.Error)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func UpdateModule(c *gin.Context) {
	var module models.Module
	if err := c.ShouldBindJSON(&module); err != nil {
		renderError(c, err)
		return
	}

	tx := db.Conn.Model(&models.Module{}).First(&models.Module{}, c.Param("id")).Updates(&module)
	if tx.Error != nil {
		renderError(c, tx.Error)
		return
	}

	moduleId, _ := strconv.Atoi(c.Param("id"))
	for _, ma := range module.ModuleActivities {
		ma.ModuleId = moduleId
		tx = db.Conn.Where(models.ModuleActivity{ModuleId: moduleId, ActivityId: ma.ActivityId}).
			FirstOrCreate(&models.ModuleActivity{}).
			Updates(&ma)
		if tx.Error != nil {
			renderError(c, tx.Error)
			return
		}
	}

	c.JSON(http.StatusOK, nil)
}

func DeleteModule(c *gin.Context) {
	tx := db.Conn.Delete(&models.Module{}, c.Param("id"))
	if tx.Error != nil {
		renderError(c, tx.Error)
		return
	}

	c.JSON(http.StatusOK, nil)
}

package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	db "upcourse/config"
	"upcourse/models"
)

func GetModule(c *gin.Context) {
	var module models.Module
	err := db.Conn.Preload("ModuleActivities.Activity").First(&module, c.Param("id")).Error

	if err != nil {
		renderError(c, err)
		return
	}

	serializedModule := SerializeModule(module)
	renderFoundRecords(c, serializedModule)
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
		renderError(c, err)
		return
	}
	input.CourseId = courseId

	err = db.Conn.Create(&input).Error
	if err != nil {
		renderError(c, err)
		return
	}

	renderSuccess(c, http.StatusCreated)
}

func UpdateModule(c *gin.Context) {
	err := db.Conn.First(&models.Module{}, c.Param("id")).Error
	if err != nil {
		renderError(c, err)
		return
	}

	var module models.Module
	if err := c.ShouldBindJSON(&module); err != nil {
		renderError(c, err)
		return
	}

	err = db.Conn.Model(&models.Module{}).Where("id = ?", c.Param("id")).Updates(&module).Error
	if err != nil {
		renderError(c, err)
		return
	}

	moduleId, _ := strconv.Atoi(c.Param("id"))
	for _, ma := range module.ModuleActivities {
		ma.ModuleId = moduleId
		err := db.Conn.Where(models.ModuleActivity{ModuleId: moduleId, ActivityId: ma.ActivityId}).
			FirstOrCreate(&models.ModuleActivity{}).
			Updates(&ma).Error
		if err != nil {
			renderError(c, err)
			return
		}
	}

	renderSuccess(c, http.StatusOK)
}

func DeleteModule(c *gin.Context) {
	err := db.Conn.Delete(&models.Module{}, c.Param("id")).Error
	if err != nil {
		renderError(c, err)
		return
	}

	renderSuccess(c, http.StatusOK)
}

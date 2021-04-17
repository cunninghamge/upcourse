package handlers

import (
	db "course-chart/config"
	"course-chart/models"

	"github.com/gin-gonic/gin"
)

func GetModule(c *gin.Context) {
	var module models.Module
	err := db.Conn.Preload("ModuleActivities.Activity").First(&module, c.Param("id")).Error

	if err != nil {
		renderNotFound(c, err)
		return
	}

	renderFound(c, module, "Module found")
}

func CreateModule(c *gin.Context) {
	var input models.Module

	if bindErr := c.ShouldBindJSON(&input); bindErr != nil {
		renderBindError(c, bindErr)
		return
	}

	err := db.Conn.Create(&input).Error
	if err != nil {
		renderError(c, err)
		return
	}

	renderCreated(c, "Module created successfully")
}

func UpdateModule(c *gin.Context) {
	err := db.Conn.First(&models.Module{}, c.Param("id")).Error
	if err != nil {
		renderNotFound(c, err)
		return
	}

	var input models.UpdatableModule
	if bindErr := c.ShouldBindJSON(&input); bindErr != nil {
		renderBindError(c, bindErr)
		return
	}

	module := models.Module{
		Name:             input.Name,
		Number:           input.Number,
		ModuleActivities: input.ModuleActivities,
	}

	err = db.Conn.Model(&models.Module{}).Where("id = ?", c.Param("id")).Updates(&module).Error
	if err != nil {
		renderError(c, err)
		return
	}

	for _, modActivity := range module.ModuleActivities {
		err = db.Conn.Model(&models.ModuleActivity{}).Where("id = ?", modActivity.ID).Updates(&modActivity).Error
		if err != nil {
			renderError(c, err)
			return
		}
	}

	renderSuccess(c, "Module updated successfully")
}

func DeleteModule(c *gin.Context) {
	err := db.Conn.First(&models.Module{}, c.Param("id")).Error
	if err != nil {
		renderNotFound(c, err)
		return
	}

	err = db.Conn.Delete(&models.Module{}, c.Param("id")).Error
	if err != nil {
		renderError(c, err)
		return
	}

	renderSuccess(c, "Module deleted successfully")
}

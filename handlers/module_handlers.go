package handlers

import (
	"net/http"
	"strconv"
	db "upcourse/config"
	"upcourse/models"

	"github.com/gin-gonic/gin"
)

func GetModule(c *gin.Context) {
	var module models.Module
	err := db.Conn.Preload("ModuleActivities.Activity").First(&module, c.Param("id")).Error

	if err != nil {
		renderError(c, err)
		return
	}

	renderFoundRecords(c, module)
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

	err := db.Conn.Create(&input).Error
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

	var existingActivityIds []int
	err = db.Conn.Model(&models.ModuleActivity{}).Where("module_id = ?", c.Param("id")).Select("activity_id").Scan(&existingActivityIds).Error
	if err != nil {
		renderError(c, err)
		return
	}

	for _, modActivity := range module.ModuleActivities {
		modActivity.ModuleId, _ = strconv.Atoi(c.Param("id"))
		if contains(existingActivityIds, modActivity.ActivityId) == true {
			err = db.Conn.Model(&models.ModuleActivity{}).
				Where("module_id = ? AND activity_id = ?", modActivity.ModuleId, modActivity.ActivityId).
				Updates(&modActivity).Error
			if err != nil {
				renderError(c, err)
				return
			}
		} else {
			err = db.Conn.Model(&models.ModuleActivity{}).Create(&modActivity).Error
			if err != nil {
				renderError(c, err)
				return
			}
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

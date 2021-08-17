package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"upcourse/models"
)

const moduleSchema = "./schemas/module_schema.json"

func GetModule(c *gin.Context) {
	module, err := models.GetModule(c.Param("id"))
	if err != nil {
		renderErrors(c, err)
		return
	}

	renderFoundRecords(c, module)
}

func CreateModule(c *gin.Context) {
	courseId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		renderErrors(c, errors.New("invalid request"))
		return
	}

	jsonData, errs := Validate(c, moduleSchema)
	if errs != nil {
		renderErrors(c, errs...)
		return
	}

	module := models.Module{CourseId: courseId}
	if err := json.Unmarshal(jsonData, &module); err != nil {
		renderErrors(c, err)
		return
	}

	if err := models.CreateModule(&module); err != nil {
		renderErrors(c, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func UpdateModule(c *gin.Context) {
	var module models.Module
	if err := c.ShouldBindJSON(&module); err != nil {
		renderErrors(c, err)
		return
	}

	if err := models.UpdateModule(&module, c.Param("id")); err != nil {
		renderErrors(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func DeleteModule(c *gin.Context) {
	if err := models.DeleteModule(c.Param("id")); err != nil {
		renderErrors(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

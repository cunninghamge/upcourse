package main

import (
	"bytes"
	"course-chart/config"
	"course-chart/models"
	"course-chart/routes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPOSTModules(t *testing.T) {
	gin.SetMode(gin.TestMode)
	config.Connect()

	t.Run("it posts a new module with associated activiies", func(t *testing.T) {
		var moduleCount int64
		config.Conn.Model(models.Module{}).Count(&moduleCount)
		var activityCount int64
		config.Conn.Model(models.ModuleActivity{}).Count(&activityCount)

		newModule := `{
			"name": "Module 9",
			"number": 9,
			"courseId": 9999,
			"moduleActivities":[
				{
					"input": 30,
					"notes": "A note",
					"activityId": 1
				},
				{
					"input": 8,
					"notes": null,
					"activityId": 2
				},
				{
					"input": 180,
					"notes": "",
					"activityId": 11
				}
			]
		}`

		router := routes.GetRoutes()
		request, _ := http.NewRequest("POST", "/modules", bytes.NewBufferString(newModule))
		response := httptest.NewRecorder()

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(response, request)

		assert.Equal(t, 201, response.Code)

		body, _ := ioutil.ReadAll(response.Body)
		postModuleResponse := MarshaledPostModuleResponse{}
		err := json.Unmarshal([]byte(body), &postModuleResponse)

		if err != nil {
			t.Errorf("Error marshaling JSON response\nError: %v", err)
		}

		assertResponseValue(t, postModuleResponse.Message, "Module created successfully", "Message")

		var newModuleCount int64
		config.Conn.Model(models.Module{}).Count(&newModuleCount)

		var newActivityCount int64
		config.Conn.Model(models.ModuleActivity{}).Count(&newActivityCount)

		if newModuleCount != (moduleCount + 1) {
			t.Errorf("did not create a new module record")
		}

		if newActivityCount != (activityCount + 3) {
			t.Errorf("did not create a new module record")
		}
	})
}

type MarshaledPostModuleResponse struct {
	Message string
	Status  int
}

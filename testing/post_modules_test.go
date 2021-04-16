package testing

import (
	"bytes"
	"course-chart/config"
	"course-chart/models"
	"course-chart/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPOSTModules(t *testing.T) {
	newSimpleCourse()
	defer teardown()

	t.Run("it posts a new module with associated activiies", func(t *testing.T) {
		var moduleCount int64
		config.Conn.Model(models.Module{}).Count(&moduleCount)
		var activityCount int64
		config.Conn.Model(models.ModuleActivity{}).Count(&activityCount)

		newModule := `{
			"name": "Module 9",
			"number": 9,
			"courseId": 1,
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

		response := newPostModuleRequest(newModule)

		assert.Equal(t, 201, response.Code)

		parsedResponse := UnmarshalSuccess(t, response.Body)

		assertResponseValue(t, parsedResponse.Message, "Module created successfully", "Message")

		var newModuleCount int64
		config.Conn.Model(models.Module{}).Count(&newModuleCount)

		var newActivityCount int64
		config.Conn.Model(models.ModuleActivity{}).Count(&newActivityCount)

		if newModuleCount != (moduleCount + 1) {
			t.Errorf("module count did not change")
		}

		if newActivityCount != (activityCount + 3) {
			t.Errorf("activity count did not change")
		}
	})
}

func newPostModuleRequest(json string) *httptest.ResponseRecorder {
	router := routes.GetRoutes()
	request, _ := http.NewRequest("POST", "/modules", bytes.NewBufferString(json))
	response := httptest.NewRecorder()

	request.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(response, request)

	return response
}

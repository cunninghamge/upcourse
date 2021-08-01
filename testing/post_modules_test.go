package testing

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"upcourse/config"
	"upcourse/models"
	"upcourse/server"

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

	t.Run("it returns an error if a required field is missing", func(t *testing.T) {
		var moduleCount int64
		config.Conn.Model(models.Module{}).Count(&moduleCount)
		var activityCount int64
		config.Conn.Model(models.ModuleActivity{}).Count(&activityCount)

		newModule := `{
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

		assert.Equal(t, 400, response.Code)

		parsedResponse := UnmarshalErrors(t, response.Body)

		expected := []string{"Name is required"}
		if !reflect.DeepEqual(parsedResponse.Errors, expected) {
			t.Errorf("got %v, wanted %v for field Error messages", parsedResponse.Errors, expected)
		}

		var newModuleCount int64
		config.Conn.Model(models.Module{}).Count(&newModuleCount)

		var newActivityCount int64
		config.Conn.Model(models.ModuleActivity{}).Count(&newActivityCount)

		if newModuleCount != moduleCount {
			t.Errorf("module count changed but should not have")
		}

		if newActivityCount != activityCount {
			t.Errorf("module activity count changed but should not have")
		}
	})

	t.Run("it returns an error if database is unavailable", func(t *testing.T) {
		db, _ := config.Conn.DB()
		db.Close()

		newModule := `{
			"name":"Module 9",
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

		assert.Equal(t, 503, response.Code)

		config.Connect()
	})

	t.Run("it returns an error if no body is sent", func(t *testing.T) {
		response := newPostModuleRequest("")

		assert.Equal(t, 400, response.Code)
	})
}

func newPostModuleRequest(json string) *httptest.ResponseRecorder {
	router := server.AppRouter()
	request, _ := http.NewRequest("POST", "/modules", bytes.NewBufferString(json))
	response := httptest.NewRecorder()

	request.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(response, request)

	return response
}

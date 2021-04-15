package main

import (
	"course-chart/config"
	"course-chart/models"
	"course-chart/routes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGETCourse(t *testing.T) {
	t.Run("returns the name of a course", func(t *testing.T) {
		router := routes.GetRoutes()
		request, _ := http.NewRequest("GET", fmt.Sprintf("/courses/%d", 9999), nil)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, 200, response.Code)

		body, _ := ioutil.ReadAll(response.Body)
		marshaledCourse := MarshaledCourse{}
		err := json.Unmarshal([]byte(body), &marshaledCourse)
		if err != nil {
			t.Errorf("Error marshaling JSON response\nError: %v", err)
		}

		var course models.Course
		db := config.Conn
		db.Preload("Modules.ModuleActivities.Activity").First(&course, 9999)

		if reflect.DeepEqual(marshaledCourse.Data.Course, models.Course{}) {
			t.Errorf("response does not contain an id property")
		}

		assertResponseValue(t, nestedCourse.Message, "Course found", "Response message")
		assertResponseValue(t, nestedCourse.Data.Course.ID, course.ID, "Id")
		assertResponseValue(t, nestedCourse.Data.Course.Name, course.Name, "Name")
		assertResponseValue(t, nestedCourse.Data.Course.CreditHours, course.CreditHours, "CreditHours")
		assertResponseValue(t, nestedCourse.Data.Course.Length, course.Length, "Length")
		firstResponseModule := nestedCourse.Data.Course.Modules[0]
		firstDBModule := course.Modules[0]
		assertResponseValue(t, firstResponseModule.ID, firstDBModule.ID, "Module Id")
		assertResponseValue(t, firstResponseModule.Name, firstDBModule.Name, "Module Name")
		assertResponseValue(t, firstResponseModule.Number, firstDBModule.Number, "Module Number")
		firstResponseModActivity := firstResponseModule.ModuleActivities[0]
		firstDBModActivity := firstDBModule.ModuleActivities[0]
		assertResponseValue(t, firstResponseModActivity.Input, firstDBModActivity.Input, "Module Activity Input")
		assertResponseValue(t, firstResponseModActivity.Notes, firstDBModActivity.Notes, "Module Activity Notes")
		firstResponseActivity := firstResponseModActivity.Activity
		firstDBActivity := firstDBModActivity.Activity
		assertResponseValue(t, firstResponseActivity.ID, firstDBActivity.ID, "Activity Id")
		assertResponseValue(t, firstResponseActivity.Description, firstDBActivity.Description, "Activity Description")
		assertResponseValue(t, firstResponseActivity.Metric, firstDBActivity.Metric, "Activity Metric")
		assertResponseValue(t, firstResponseActivity.Multiplier, firstDBActivity.Multiplier, "Activity Multiplier")
	})

	t.Run("returns a message if the course is not found", func(t *testing.T) {
		router := routes.GetRoutes()
		request, _ := http.NewRequest("GET", fmt.Sprintf("/courses/%d", 999999), nil)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, 404, response.Code)

		body, _ := ioutil.ReadAll(response.Body)
		marshaledResponse := MarshaledError{}
		err := json.Unmarshal([]byte(body), &marshaledResponse)

		if err != nil {
			t.Errorf("Error marshaling JSON response\nError: %v", err)
		}

		assertResponseValue(t, marshaledResponse.Errors, "Record not found", "Response message")
	})
}

type MarshaledCourse struct {
	Data struct {
		Course         models.Course
		ActivityTotals []models.ActivityTotals
	}
	Message string
	Status  int
}

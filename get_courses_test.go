package main

import (
	"course-chart/config"
	"course-chart/models"
	"course-chart/routes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGETCourses(t *testing.T) {
	gin.SetMode(gin.TestMode)
	config.Connect()

	t.Run("returns a list of courses", func(t *testing.T) {
		router := routes.GetRoutes()
		request, _ := http.NewRequest("GET", "/courses", nil)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, 200, response.Code)

		body, _ := ioutil.ReadAll(response.Body)
		courseList := MarshaledCourses{}
		err := json.Unmarshal([]byte(body), &courseList)
		if err != nil {
			t.Errorf("Error marshaling JSON response\nError: %v", err)
		}

		var courses []models.Course
		db := config.Conn
		db.Find(&courses)

		if reflect.DeepEqual(courseList.Data[0], models.Course{}) {
			t.Errorf("response does not contain an id property")
		}

		assertResponseValue(t, courseList.Message, "Courses found", "Response message")
		firstCourse := courseList.Data[0]
		assertResponseValue(t, firstCourse.Id, courses[0].Id, "Id")
		assertResponseValue(t, firstCourse.Name, courses[0].Name, "Name")
	})
}

type MarshaledCourses struct {
	Data    []models.Course
	Message string
	Status  int
}

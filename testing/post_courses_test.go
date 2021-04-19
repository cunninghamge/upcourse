package testing

import (
	"bytes"
	"course-chart/config"
	"course-chart/models"
	"course-chart/routes"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostCourses(t *testing.T) {
	defer teardown()

	t.Run("it posts a new course", func(t *testing.T) {
		var courseCount int64
		config.Conn.Model(models.Course{}).Count(&courseCount)

		newCourse := `{
			"name": "Nursing 101",
			"institution": "Tampa Bay Nurses United University",
			"creditHours": 3,
			"length": 16,
			"goal": "8-10 hours"
			}`

		response := newPostCourseRequest(newCourse)

		assert.Equal(t, 201, response.Code)

		parsedResponse := UnmarshalSuccess(t, response.Body)

		assertResponseValue(t, parsedResponse.Message, "Course created successfully", "Message")

		var newCount int64
		config.Conn.Model(models.Course{}).Count(&newCount)

		if newCount != (courseCount + 1) {
			t.Errorf("course count did not change")
		}
	})

	t.Run("it returns an error if a required field is missing", func(t *testing.T) {
		var courseCount int64
		config.Conn.Model(models.Course{}).Count(&courseCount)

		newCourse := `{
			"creditHours": 3,
			"length": 16,
			"goal": "8-10 hours"
			}`

		response := newPostCourseRequest(newCourse)
		log.Print(response.Body.String())
		assert.Equal(t, 400, response.Code)

		parsedResponse := UnmarshalErrors(t, response.Body)

		expected := []string{"Name is required", "Institution is required"}
		if !reflect.DeepEqual(parsedResponse.Errors, expected) {
			t.Errorf("got %v, wanted %v for field Error messages", parsedResponse.Errors, expected)
		}

		var newCount int64
		config.Conn.Model(models.Course{}).Count(&newCount)

		if newCount != courseCount {
			t.Errorf("course count changed but should not have")
		}
	})
}

func newPostCourseRequest(json string) *httptest.ResponseRecorder {
	router := routes.GetRoutes()
	request, _ := http.NewRequest("POST", "/courses", bytes.NewBufferString(json))
	response := httptest.NewRecorder()

	request.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(response, request)

	return response
}

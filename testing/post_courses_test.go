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

	t.Run("it returns an error if database is unavailable", func(t *testing.T) {
		db, _ := config.Conn.DB()
		db.Close()

		newCourse := `{
			"name": "Nursing 101",
			"institution": "Tampa Bay Nurses United University",
			"creditHours": 3,
			"length": 16,
			"goal": "8-10 hours"
			}`
		response := newPostCourseRequest(newCourse)

		assert.Equal(t, 503, response.Code)

		config.Connect()
	})

	t.Run("it returns an error if no body is sent", func(t *testing.T) {
		response := newPostCourseRequest("")

		assert.Equal(t, 400, response.Code)
	})
}

func newPostCourseRequest(json string) *httptest.ResponseRecorder {
	router := server.AppRouter()
	request, _ := http.NewRequest("POST", "/courses", bytes.NewBufferString(json))
	response := httptest.NewRecorder()

	request.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(response, request)

	return response
}

package testing

import (
	"bytes"
	"course-chart/config"
	"course-chart/models"
	"course-chart/routes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPATCHCourse(t *testing.T) {
	course := newSimpleCourse()
	defer teardown()

	t.Run("it updates an existing course", func(t *testing.T) {
		var courseCount int64
		config.Conn.Model(models.Course{}).Count(&courseCount)

		newCourseInfo := `{
			"institution": "Tampa Bay Nurses United University",
			"creditHours": 3,
			"length": 16,
			"goal": "8-10 hours"
			}`

		response := newPatchCourseRequest(newCourseInfo, course.ID)

		assert.Equal(t, 200, response.Code)

		parsedResponse := UnmarshalSuccess(t, response.Body)

		assertResponseValue(t, parsedResponse.Message, "Course updated successfully", "Message")

		var newCount int64
		config.Conn.Model(models.Course{}).Count(&newCount)

		if newCount != courseCount {
			t.Errorf("patch request should not have changed course count")
		}

		var updatedCourse models.Course
		config.Conn.First(&updatedCourse, course.ID)

		if updatedCourse.Institution != "Tampa Bay Nurses United University" {
			t.Errorf("did not update the course record")
		}

		if updatedCourse.Name != "Foundations of Nursing" {
			t.Errorf("updated a field that shouldn't have been updated")
		}
	})

	t.Run("it returns an error if database is unavailable", func(t *testing.T) {
		db, _ := config.Conn.DB()
		db.Close()

		newCourseInfo := `{
			"institution": "Tampa Bay Nurses United University",
			"creditHours": 3,
			"length": 16,
			"goal": "8-10 hours"
			}`
		response := newPatchCourseRequest(newCourseInfo, course.ID)

		assert.Equal(t, 404, response.Code)

		config.Connect()
	})

	t.Run("it returns an error if no body is sent", func(t *testing.T) {
		response := newPatchCourseRequest("", course.ID)

		assert.Equal(t, 400, response.Code)
	})
}

func newPatchCourseRequest(json string, courseId int) *httptest.ResponseRecorder {
	router := routes.GetRoutes()
	request, _ := http.NewRequest("PATCH", fmt.Sprintf("/courses/%d", courseId), bytes.NewBufferString(json))
	response := httptest.NewRecorder()

	request.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(response, request)

	return response
}

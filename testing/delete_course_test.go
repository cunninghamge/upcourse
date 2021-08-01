package testing

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"upcourse/config"
	"upcourse/models"
	"upcourse/server"

	"github.com/stretchr/testify/assert"
)

func TestDELETEcourse(t *testing.T) {
	mockCourse := newFullCourse()
	defer teardown()

	t.Run("Deletes a course and its children", func(t *testing.T) {
		var courseCount int64
		config.Conn.Model(models.Course{}).Count(&courseCount)

		response := newDeleteCourseRequest(mockCourse.ID)

		assert.Equal(t, 200, response.Code)

		parsedResponse := UnmarshalSuccess(t, response.Body)

		assertResponseValue(t, parsedResponse.Message, "Course deleted successfully", "Message")

		var newCourseCount int64
		config.Conn.Model(models.Course{}).Count(&newCourseCount)

		if courseCount == newCourseCount {
			t.Errorf("Did not delete course")
		}

		err := config.Conn.First(&models.Course{}, mockCourse.ID).Error
		if err == nil {
			t.Errorf("Also did not delete course")
		}

		err = config.Conn.First(&models.Module{}, mockCourse.Modules[0].ID).Error
		if err == nil {
			t.Errorf("Did not delete associated modules")
		}

		err = config.Conn.First(&models.ModuleActivity{}, mockCourse.Modules[0].ModuleActivities[0].ID).Error
		if err == nil {
			t.Errorf("Did not delete associated module activities")
		}
	})

	t.Run("returns an error if database is unavailable", func(t *testing.T) {
		db, _ := config.Conn.DB()
		db.Close()
		response := newDeleteCourseRequest(mockCourse.ID)

		assert.Equal(t, 503, response.Code)

		config.Connect()
	})
}

func newDeleteCourseRequest(courseId int) *httptest.ResponseRecorder {
	router := server.AppRouter()
	request, _ := http.NewRequest("DELETE", fmt.Sprintf("/courses/%d", courseId), nil)
	response := httptest.NewRecorder()

	request.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(response, request)

	return response
}

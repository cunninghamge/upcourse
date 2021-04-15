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

func TestPATCHCourse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	config.Connect()

	t.Run("it updates an existing course", func(t *testing.T) {
		var courseCount int64
		config.Conn.Model(models.Course{}).Count(&courseCount)

		newCourseInfo := `{
			"institution": "Tampa Bay Nurses United University",
			"creditHours": 3,
			"length": 16,
			"goal": "8-10 hours"
			}`

		router := routes.GetRoutes()
		request, _ := http.NewRequest("PATCH", "/courses/9999", bytes.NewBufferString(newCourseInfo))
		response := httptest.NewRecorder()

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(response, request)

		assert.Equal(t, 200, response.Code)

		body, _ := ioutil.ReadAll(response.Body)
		patchCourseResponse := MarshaledPatchCourseResponse{}
		err := json.Unmarshal([]byte(body), &patchCourseResponse)

		if err != nil {
			t.Errorf("Error marshaling JSON response\nError: %v", err)
		}

		assertResponseValue(t, patchCourseResponse.Message, "Course updated successfully", "Message")

		var newCount int64
		config.Conn.Model(models.Course{}).Count(&newCount)

		if newCount != (courseCount) {
			t.Errorf("did not update the course record")
		}
		var course models.Course

		config.Conn.First(&course, 9999)

		if course.Institution != "Tampa Bay Nurses United University" {
			t.Errorf("did not update the course record")
		}

		if course.Name != "Foundations of Nursing" {
			t.Errorf("updated a field that shouldn't have been updated")
		}
	})
}

type MarshaledPatchCourseResponse struct {
	Message string
	Status  int
}

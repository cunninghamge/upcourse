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

func TestPostCourses(t *testing.T) {
	gin.SetMode(gin.TestMode)
	config.Connect()

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

		router := routes.GetRoutes()
		request, _ := http.NewRequest("POST", "/courses", bytes.NewBufferString(newCourse))
		response := httptest.NewRecorder()

		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(response, request)

		assert.Equal(t, 201, response.Code)

		body, _ := ioutil.ReadAll(response.Body)
		postCourseResponse := MarshaledPostCourseResponse{}
		err := json.Unmarshal([]byte(body), &postCourseResponse)

		if err != nil {
			t.Errorf("Error marshaling JSON response\nError: %v", err)
		}

		assertResponseValue(t, postCourseResponse.Message, "Course created successfully", "Message")

		var newCount int64
		config.Conn.Model(models.Course{}).Count(&newCount)

		if newCount != (courseCount + 1) {
			t.Errorf("did not create a new course record")
		}
	})
}

type MarshaledPostCourseResponse struct {
	Message string
	Status  int
}

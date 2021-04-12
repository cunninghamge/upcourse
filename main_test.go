package main

import (
	"course-chart/routes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"course-chart/config"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGETCourses(t *testing.T) {
	gin.SetMode(gin.TestMode)
	config.Connect()

	t.Run("returns the name of a course", func(t *testing.T) {
		course := &routes.Course{
			Id:   1,
			Name: "Foundations of Nursing",
		}

		router := routes.GetRoutes()
		request, _ := http.NewRequest("GET", fmt.Sprintf("/courses/%d", course.Id), nil)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, course.Name, response.Body.String())
	})
}

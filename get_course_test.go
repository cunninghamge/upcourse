package main

import (
	"course-chart/config"
	"course-chart/models"
	"course-chart/routes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGETCourses(t *testing.T) {
	gin.SetMode(gin.TestMode)
	config.Connect()

	t.Run("returns the name of a course", func(t *testing.T) {
		// course := &models.Course{
		// 	Id:   1,
		// 	Name: "Foundations of Nursing",
		// }

		router := routes.GetRoutes()
		request, _ := http.NewRequest("GET", fmt.Sprintf("/courses/%d", 1), nil)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, 200, response.Code)

		body, _ := ioutil.ReadAll(response.Body)
		// var course models.Course

		type Response struct {
			data    models.Course `json:"data"`
			message string        `json:"message"`
			status  int           `json:"status"`
		}
		var nestedCourse = new(Response)
		// var course = new(models.Course)
		// var result map[string]interface{}

		err := json.Unmarshal([]byte(body), &nestedCourse)
		// something := mapstructure.Decode(result["data"], &course)

		log.Print(nestedCourse)

		// err := json.NewDecoder(response.Body).Decode(&course)

		if err != nil {
			t.Errorf("JSON response does not map to course type\nError: %v", err)
		}
	})
}

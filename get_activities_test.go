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

func TestGetActivities(t *testing.T) {
	gin.SetMode(gin.TestMode)
	config.Connect()

	t.Run("returns a list of the default activities", func(t *testing.T) {
		router := routes.GetRoutes()
		request, _ := http.NewRequest("GET", "/activities", nil)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, 200, response.Code)

		body, _ := ioutil.ReadAll(response.Body)
		activityList := MarshaledActivities{}
		err := json.Unmarshal([]byte(body), &activityList)
		if err != nil {
			t.Errorf("Error marshaling JSON response\nError: %v", err)
		}

		var activities []models.Activity
		config.Conn.Find(&activities)

		if reflect.DeepEqual(activityList.Data[0], models.Activity{}) {
			t.Errorf("response does not contain an id property")
		}

		assertResponseValue(t, activityList.Message, "Activities found", "Response message")
		assertResponseValue(t, activityList.Data[0].ID, activities[0].ID, "Id")
		assertResponseValue(t, activityList.Data[0].Name, activities[0].Name, "Name")
		assertResponseValue(t, activityList.Data[0].Description, activities[0].Description, "Description")
		assertResponseValue(t, activityList.Data[0].Metric, activities[0].Metric, "Metric")
		assertResponseValue(t, activityList.Data[0].Multiplier, activities[0].Multiplier, "Multiplier")
	})
}

type MarshaledActivities struct {
	Data    []models.Activity
	Message string
	Status  int
}

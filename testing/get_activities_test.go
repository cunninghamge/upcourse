package testing

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"upcourse/config"
	"upcourse/models"
	"upcourse/server"

	"github.com/stretchr/testify/assert"
)

func TestGetActivities(t *testing.T) {
	activities := coreActivities()

	t.Run("returns a list of the default activities", func(t *testing.T) {
		response := newGetActivitiesRequest()

		assert.Equal(t, 200, response.Code)

		parsedResponse := unmarshalGETActivities(t, response.Body)

		assertResponseValue(t, parsedResponse.Message, "Activities found", "Response message")
		assertResponseValue(t, parsedResponse.Data[0].ID, activities[0].ID, "Id")
		assertResponseValue(t, parsedResponse.Data[0].Name, activities[0].Name, "Name")
		assertResponseValue(t, parsedResponse.Data[0].Description, activities[0].Description, "Description")
		assertResponseValue(t, parsedResponse.Data[0].Metric, activities[0].Metric, "Metric")
		assertResponseValue(t, parsedResponse.Data[0].Multiplier, activities[0].Multiplier, "Multiplier")
	})

	t.Run("does not include custom activities", func(t *testing.T) {
		config.Conn.Create(&models.Activity{Custom: true})

		response := newGetActivitiesRequest()

		assert.Equal(t, 200, response.Code)

		parsedResponse := unmarshalGETActivities(t, response.Body)

		assertResponseValue(t, len(parsedResponse.Data), len(activities), "number of activities")
	})

	t.Run("returns an error if database is unavailable", func(t *testing.T) {
		db, _ := config.Conn.DB()
		db.Close()
		response := newGetActivitiesRequest()

		assert.Equal(t, 503, response.Code)

		config.Connect()
	})
}

func newGetActivitiesRequest() *httptest.ResponseRecorder {
	router := server.AppRouter()
	request, _ := http.NewRequest("GET", "/activities", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	return response
}

func unmarshalGETActivities(t *testing.T, response io.Reader) getActivitiesResponse {
	t.Helper()
	body, _ := ioutil.ReadAll(response)

	activityList := getActivitiesResponse{}
	err := json.Unmarshal([]byte(body), &activityList)
	if err != nil {
		t.Errorf("Error marshaling JSON response\nError: %v", err)
	}

	if reflect.DeepEqual(activityList.Data[0], models.Activity{}) {
		t.Errorf("response does not contain an id property")
	}

	return activityList
}

type getActivitiesResponse struct {
	Data    []models.Activity
	Message string
	Status  int
}

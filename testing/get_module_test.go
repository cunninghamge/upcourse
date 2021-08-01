package testing

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"upcourse/models"
	"upcourse/server"

	"github.com/stretchr/testify/assert"
)

func TestGETModule(t *testing.T) {
	mockModule := newModule()
	defer teardown()

	t.Run("returns a module", func(t *testing.T) {
		response := newGetModuleRequest(mockModule.ID)

		assert.Equal(t, 200, response.Code)

		parsedResponse := unmarshalGETModule(t, response.Body)

		assertResponseValue(t, parsedResponse.Message, "Module found", "Response message")
		assertResponseValue(t, parsedResponse.Data.ID, mockModule.ID, "Id")
		assertResponseValue(t, parsedResponse.Data.Name, mockModule.Name, "Name")
		assertResponseValue(t, parsedResponse.Data.Number, mockModule.Number, "Number")
		assertResponseValue(t, parsedResponse.Data.CourseId, mockModule.CourseId, "CourseId")
		firstResponseModActivity := parsedResponse.Data.ModuleActivities[0]
		firstMockModActivity := mockModule.ModuleActivities[0]
		assertResponseValue(t, firstResponseModActivity.Input, firstMockModActivity.Input, "Module Activity Input")
		assertResponseValue(t, firstResponseModActivity.Notes, firstMockModActivity.Notes, "Module Activity Notes")
		firstResponseActivity := firstResponseModActivity.Activity
		firstMockActivity := firstMockModActivity.Activity
		assertResponseValue(t, firstResponseActivity.ID, firstMockActivity.ID, "Activity Id")
		assertResponseValue(t, firstResponseActivity.Description, firstMockActivity.Description, "Activity Description")
		assertResponseValue(t, firstResponseActivity.Metric, firstMockActivity.Metric, "Activity Metric")
		assertResponseValue(t, firstResponseActivity.Multiplier, firstMockActivity.Multiplier, "Activity Multiplier")
	})

	t.Run("returns a message if the module is not found", func(t *testing.T) {
		response := newGetModuleRequest(mockModule.ID + 1)

		assert.Equal(t, 404, response.Code)

		parsedResponse := UnmarshalError(t, response.Body)

		assertResponseValue(t, parsedResponse.Errors, "Record not found", "Response message")
	})
}

func newGetModuleRequest(moduleId int) *httptest.ResponseRecorder {
	router := server.AppRouter()
	request, _ := http.NewRequest("GET", fmt.Sprintf("/modules/%d", moduleId), nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	return response
}

func unmarshalGETModule(t *testing.T, response io.Reader) getModuleResponse {
	t.Helper()
	body, _ := ioutil.ReadAll(response)

	responseModule := getModuleResponse{}
	err := json.Unmarshal([]byte(body), &responseModule)
	if err != nil {
		t.Errorf("Error marshaling JSON response\nError: %v", err)
	}

	if reflect.DeepEqual(responseModule.Data, models.Module{}) {
		t.Errorf("response does not contain an id property")
	}

	return responseModule
}

type getModuleResponse struct {
	Data    models.Module
	Message string
	Status  int
}

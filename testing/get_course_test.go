package testing

import (
	"course-chart/models"
	"course-chart/routes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGETCourse(t *testing.T) {
	mockCourse := newMockCourse()
	defer teardown()

	t.Run("returns the name of a course", func(t *testing.T) {
		response := newGetCourseRequest(mockCourse.ID)

		assert.Equal(t, 200, response.Code)

		parsedResponse := unmarshalGETCourse(t, response.Body)

		assertResponseValue(t, parsedResponse.Message, "Course found", "Response message")
		responseCourse := parsedResponse.Data.Course
		assertResponseValue(t, responseCourse.ID, mockCourse.ID, "Id")
		assertResponseValue(t, responseCourse.Name, mockCourse.Name, "Name")
		assertResponseValue(t, responseCourse.CreditHours, mockCourse.CreditHours, "CreditHours")
		assertResponseValue(t, responseCourse.Length, mockCourse.Length, "Length")
		firstResponseModule := responseCourse.Modules[0]
		firstMockModule := mockCourse.Modules[0]
		assertResponseValue(t, firstResponseModule.ID, firstMockModule.ID, "Module Id")
		assertResponseValue(t, firstResponseModule.Name, firstMockModule.Name, "Module Name")
		assertResponseValue(t, firstResponseModule.Number, firstMockModule.Number, "Module Number")
		firstResponseModActivity := firstResponseModule.ModuleActivities[0]
		firstMockModActivity := firstMockModule.ModuleActivities[0]
		assertResponseValue(t, firstResponseModActivity.Input, firstMockModActivity.Input, "Module Activity Input")
		assertResponseValue(t, firstResponseModActivity.Notes, firstMockModActivity.Notes, "Module Activity Notes")
		firstResponseActivity := firstResponseModActivity.Activity
		firstMockActivity := firstMockModActivity.Activity
		assertResponseValue(t, firstResponseActivity.ID, firstMockActivity.ID, "Activity Id")
		assertResponseValue(t, firstResponseActivity.Description, firstMockActivity.Description, "Activity Description")
		assertResponseValue(t, firstResponseActivity.Metric, firstMockActivity.Metric, "Activity Metric")
		assertResponseValue(t, firstResponseActivity.Multiplier, firstMockActivity.Multiplier, "Activity Multiplier")
	})

	t.Run("returns a message if the course is not found", func(t *testing.T) {
		response := newGetCourseRequest(mockCourse.ID + 1)

		assert.Equal(t, 404, response.Code)

		parsedResponse := UnmarshalError(t, response.Body)

		assertResponseValue(t, parsedResponse.Errors, "Record not found", "Response message")
	})
}

func newGetCourseRequest(courseId int) *httptest.ResponseRecorder {
	router := routes.GetRoutes()

	request, _ := http.NewRequest("GET", fmt.Sprintf("/courses/%d", courseId), nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	return response
}

func unmarshalGETCourse(t *testing.T, response io.Reader) getCourseResponse {
	t.Helper()
	body, _ := ioutil.ReadAll(response)

	responseCourse := getCourseResponse{}
	err := json.Unmarshal([]byte(body), &responseCourse)
	if err != nil {
		t.Errorf("Error marshaling JSON response\nError: %v", err)
	}

	if reflect.DeepEqual(responseCourse.Data.Course, models.Course{}) {
		t.Errorf("response does not contain an id property")
	}

	return responseCourse
}

type getCourseResponse struct {
	Data struct {
		Course         models.Course
		ActivityTotals []models.ActivityTotals
	}
	Message string
	Status  int
}

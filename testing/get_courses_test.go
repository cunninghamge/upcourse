package testing

import (
	"course-chart/models"
	"course-chart/routes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGETCourses(t *testing.T) {
	mockCourses := newCourseList()
	defer teardown()

	t.Run("returns a list of courses", func(t *testing.T) {
		response := newGetCoursesRequest()

		assert.Equal(t, 200, response.Code)

		parsedResponse := unmarshalGETCourses(t, response.Body)

		assertResponseValue(t, parsedResponse.Message, "Courses found", "Response message")
		firstCourse := parsedResponse.Data[0]
		assertResponseValue(t, firstCourse.ID, mockCourses[0].ID, "Id")
		assertResponseValue(t, firstCourse.Name, mockCourses[0].Name, "Name")
	})
}

func newGetCoursesRequest() *httptest.ResponseRecorder {
	router := routes.GetRoutes()

	request, _ := http.NewRequest("GET", "/courses", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	return response
}

func unmarshalGETCourses(t *testing.T, response io.Reader) getCoursesResponse {
	t.Helper()
	body, _ := ioutil.ReadAll(response)

	responseCourses := getCoursesResponse{}
	err := json.Unmarshal([]byte(body), &responseCourses)
	if err != nil {
		t.Errorf("Error marshaling JSON response\nError: %v", err)
	}

	if reflect.DeepEqual(responseCourses.Data[0], models.CourseIdentifier{}) {
		t.Errorf("response does not contain an id property")
	}

	return responseCourses
}

type getCoursesResponse struct {
	Data    []models.Course
	Message string
	Status  int
}

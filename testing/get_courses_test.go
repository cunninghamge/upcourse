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
	"upcourse/routes"

	"github.com/stretchr/testify/assert"
)

func TestGETCourses(t *testing.T) {
	t.Run("returns a list of courses", func(t *testing.T) {
		mockCourses := newCourseList()
		defer teardown()

		response := newGetCoursesRequest()

		assert.Equal(t, 200, response.Code)

		parsedResponse := unmarshalGETCourses(t, response.Body)

		if reflect.DeepEqual(parsedResponse.Data[0], models.Course{}) {
			t.Errorf("response does not contain an id property")
		}

		assertResponseValue(t, parsedResponse.Message, "Courses found", "Response message")
		firstCourse := parsedResponse.Data[0]
		assertResponseValue(t, firstCourse.ID, mockCourses[0].ID, "Id")
		assertResponseValue(t, firstCourse.Name, mockCourses[0].Name, "Name")
		firstModule := firstCourse.Modules[0]
		assertResponseValue(t, firstModule.Name, mockCourses[0].Modules[0].Name, "Module Name")
	})

	t.Run("retuns an array even if only one course is found", func(t *testing.T) {
		newSimpleCourse()
		defer teardown()

		response := newGetCoursesRequest()

		assert.Equal(t, 200, response.Code)

		parsedResponse := unmarshalGETCourses(t, response.Body)

		if reflect.DeepEqual(parsedResponse.Data[0], models.Course{}) {
			t.Errorf("response does not contain an id property")
		}

		assertResponseValue(t, len(parsedResponse.Data), 1, "Number of results")
	})

	t.Run("retuns an array even if no courses are found", func(t *testing.T) {
		response := newGetCoursesRequest()

		assert.Equal(t, 200, response.Code)

		parsedResponse := unmarshalGETCourses(t, response.Body)

		assertResponseValue(t, len(parsedResponse.Data), 0, "Number of results")
	})

	t.Run("returns an error if database is unavailable", func(t *testing.T) {
		db, _ := config.Conn.DB()
		db.Close()
		response := newGetCoursesRequest()

		assert.Equal(t, 503, response.Code)

		config.Connect()
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

	return responseCourses
}

type getCoursesResponse struct {
	Data    []models.Course
	Message string
	Status  int
}

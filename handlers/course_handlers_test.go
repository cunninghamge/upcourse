package handlers

import (
	"fmt"
	"net/http"
	"testing"

	db "upcourse/database"
	"upcourse/models"

	"github.com/Pallinder/go-randomdata"
	"github.com/gin-gonic/gin"
	"github.com/google/jsonapi"
)

func TestCourses(t *testing.T) {
	defer teardown()

	testCases := map[string]struct {
		funcToTest      func(*gin.Context)
		params          map[string]string
		body            string
		statusCode      int
		model           interface{}
		many            bool
		getIdForCleanup bool
	}{
		"GetCourse": {
			funcToTest: GetCourse,
			params:     map[string]string{"id": fmt.Sprint(mockCourse().ID)},
			statusCode: http.StatusOK,
			model:      new(models.Course),
		},
		"GetCourses": {
			funcToTest: GetCourses,
			statusCode: http.StatusOK,
			model:      new(models.Course),
			many:       true,
		},
		"CreateCourse": {
			funcToTest: CreateCourse,
			body: `{
				"name": "Create Course Handler Test",
				"institution": "Test University",
				"creditHours": 3,
				"length": 16,
				"goal": "8-10 hours"
			}`,
			statusCode:      http.StatusCreated,
			getIdForCleanup: true,
		},
		"UpdateCourse": {
			funcToTest: UpdateCourse,
			params:     map[string]string{"id": fmt.Sprint(mockCourse().ID)},
			body: `{
				"name": "Update Course Handler Test",
				"institution": "Test University"
			}`,
			statusCode: http.StatusOK,
		},
		"DeleteCourse": {
			funcToTest: DeleteCourse,
			params:     map[string]string{"id": fmt.Sprint(mockCourse().ID)},
			statusCode: http.StatusOK,
		},
	}

	for name, tc := range testCases {
		t.Run(name+" success", func(t *testing.T) {
			w := newRequest(tc.funcToTest, tc.params, tc.body)

			assertStatusCode(t, w.Code, tc.statusCode)
			if tc.model != nil {
				unmarshalPayload(t, w.Body, tc.model, tc.many)
			}

			if tc.getIdForCleanup {
				var course models.Course
				jsonapi.UnmarshalPayload(w.Body, &course)
				courseIds = append(courseIds, course.ID)
			}
		})

		t.Run(name+" failure", func(t *testing.T) {
			forceError()
			defer clearError()

			w := newRequest(tc.funcToTest, tc.params, tc.body)

			assertStatusCode(t, w.Code, http.StatusInternalServerError)
			unmarshalPayload(t, w.Body, new(error), many)
		})
	}

	t.Run("CreateCourse validation failure", func(t *testing.T) {
		incompleteCourseInfo := `{
			"name": "Nursing 101",
			"institution": "Tampa Bay Nurses United University",
			"creditHours": 3,
			"length": 16
		}`
		w := newRequest(CreateCourse, nil, incompleteCourseInfo)

		assertStatusCode(t, w.Code, http.StatusBadRequest)
		unmarshalPayload(t, w.Body, new(error), many)
	})

	t.Run("UpdateCourse json unmarshaling failure", func(t *testing.T) {
		w := newRequest(UpdateCourse, nil, "")

		assertStatusCode(t, w.Code, http.StatusBadRequest)
		unmarshalPayload(t, w.Body, new(error), many)
	})
}

func mockCourse() *models.Course {
	course := models.Course{
		Name:        "Handlers Test Course",
		Institution: "Test University",
		CreditHours: randomdata.Number(5),
		Length:      randomdata.Number(16),
		Goal:        "8-10 hours",
		Modules: []*models.Module{
			{Number: 1},
			{Number: 2},
		},
	}

	db.Conn.Create(&course)
	courseIds = append(courseIds, course.ID)

	return &course
}

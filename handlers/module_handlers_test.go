package handlers

import (
	"fmt"
	"net/http"
	"testing"

	db "upcourse/database"
	"upcourse/models"

	"github.com/Pallinder/go-randomdata"
	"github.com/gin-gonic/gin"
)

func TestModules(t *testing.T) {
	courseId := mockCourseId()
	defer teardown()

	testCases := map[string]struct {
		funcToTest func(*gin.Context)
		params     map[string]string
		body       string
		statusCode int
		model      interface{}
		many       bool
	}{
		"GetModule": {
			funcToTest: GetModule,
			params:     map[string]string{"id": fmt.Sprint(mockModule(courseId, 1).ID)},
			statusCode: http.StatusOK,
			model:      new(models.Module),
		},
		"CreateModule": {
			funcToTest: CreateModule,
			params:     map[string]string{"id": fmt.Sprint(course.ID)},
			body: `{
				"name": "Module 9",
				"number": 9,
				"moduleActivities":[
					{
						"input": 30,
						"notes": "A note",
						"activityId": 1
					},
					{
						"input": 8,
						"activityId": 2
					},
					{
						"input": 180,
						"notes": "",
						"activityId": 11
					}
				]
			}`,
			statusCode: http.StatusCreated,
		},
		"UpdateModule": {
			funcToTest: UpdateModule,
			params:     map[string]string{"id": fmt.Sprint(mockModule(courseId, 2).ID)},
			body: `{
				"name": "new module name",
				"moduleActivities":[
					{
						"input": 30,
						"notes": "A new note",
						"activityId": 1
					}
				]
			}`,
			statusCode: http.StatusOK,
		},
		"DeleteModule": {
			funcToTest: DeleteModule,
			params:     map[string]string{"id": fmt.Sprint(mockModule(courseId, 3).ID)},
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
		})

		t.Run(name+" failure", func(t *testing.T) {
			forceError()
			defer clearError()

			w := newRequest(tc.funcToTest, tc.params, tc.body)

			assertStatusCode(t, w.Code, http.StatusInternalServerError)

			unmarshalPayload(t, w.Body, new(error), many)
		})

	}

	t.Run("CreateModule missing course failure", func(t *testing.T) {
		moduleInfo := `{
			"name": "Module 9",
			"number": 9,
			"moduleActivities":[
				{
					"input": 180,
					"activityId": 11
				}
			]
		}`
		params := map[string]string{"id": "-"}
		w := newRequest(CreateModule, params, moduleInfo)

		assertStatusCode(t, w.Code, http.StatusBadRequest)

		unmarshalPayload(t, w.Body, new(error), many)
	})

	t.Run("CreateModule missing info failure", func(t *testing.T) {
		moduleInfo := `{
			"name": "Module 9",
			"moduleActivities":[
				{
					"input": 180,
					"activityId": 11
				}
			]
		}`
		params := map[string]string{"id": fmt.Sprint(courseId)}
		w := newRequest(CreateModule, params, moduleInfo)

		assertStatusCode(t, w.Code, http.StatusBadRequest)

		unmarshalPayload(t, w.Body, new(error), many)
	})

	t.Run("UpdateModule json unmarshal failure", func(t *testing.T) {
		w := newRequest(UpdateModule, nil, "")

		assertStatusCode(t, w.Code, http.StatusBadRequest)

		unmarshalPayload(t, w.Body, new(error), many)
	})
}

func mockCourseId() int {
	course := models.Course{
		Name:        "Handlers Test Course",
		Institution: "Test University",
		CreditHours: randomdata.Number(5),
		Length:      randomdata.Number(16),
		Goal:        "8-10 hours",
	}

	db.Conn.Create(&course)

	return course.ID
}

func mockModule(courseId, moduleNumber int) *models.Module {
	module := models.Module{
		Name:     "Handlers Test Module",
		Number:   moduleNumber,
		CourseId: courseId,
		ModuleActivities: []*models.ModuleActivity{
			{Input: 10, ActivityId: 1},
			{Input: 20, ActivityId: 2},
		},
	}

	db.Conn.Create(&module)

	return &module
}

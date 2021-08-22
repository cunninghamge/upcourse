package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	db "upcourse/database"
	"upcourse/models"

	"github.com/gin-gonic/gin"
)

func defaultActivities() []*models.Activity {
	var activities []*models.Activity
	db.Conn.Select("id, name, description, metric, multiplier").Where("custom=false").Find(&activities)

	return activities
}

func TestRenderFoundRecords(t *testing.T) {
	defer teardown()

	testCases := map[string]struct {
		records interface{}
		model   interface{}
		many    bool
	}{
		"course": {
			records: mockCourse(),
			model:   &models.Course{},
			many:    false,
		},
		"courses": {
			records: []*models.Course{
				mockCourse(),
				mockCourse(),
			},
			model: &models.Course{},
			many:  true,
		},
		"module": {
			records: mockModule(mockCourseId(), 1),
			model:   &models.Module{},
			many:    false,
		},
		"activities": {
			records: defaultActivities(),
			model:   &models.Activity{},
			many:    true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			renderFoundRecords(ctx, tc.records)

			assertStatusCode(t, w.Code, http.StatusOK)
			unmarshalPayload(t, w.Body, tc.model, tc.many)
		})
	}

	t.Run("it returns jsonapi errors", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		renderFoundRecords(ctx, models.Course{})
		assertStatusCode(t, w.Code, http.StatusInternalServerError)
	})
}

func TestRenderError(t *testing.T) {
	for message, code := range errCodes {
		t.Run(message, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			renderErrors(ctx, errors.New(message))

			assertStatusCode(t, w.Code, code)
			unmarshalPayload(t, w.Body, new(error), many)
		})
	}

	t.Run("default error", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		renderErrors(ctx, errors.New(testDBErr))

		assertStatusCode(t, w.Code, http.StatusInternalServerError)
		unmarshalPayload(t, w.Body, new(error), many)
	})

	t.Run("render multiple errors", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		errs := []error{
			errors.New("first error"),
			errors.New("second error"),
		}

		renderErrors(ctx, errs...)

		assertStatusCode(t, w.Code, http.StatusBadRequest)
		unmarshalPayload(t, w.Body, new(error), many)
	})
}

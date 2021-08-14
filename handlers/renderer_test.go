package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	testHelpers "upcourse/internal/helpers"
	"upcourse/internal/mocks"
)

func TestRenderFoundRecords(t *testing.T) {
	testCases := map[string]interface{}{
		"course":     mocks.FullCourse(),
		"courses":    mocks.CourseList(),
		"module":     mocks.Module(),
		"activities": mocks.DefaultActivities(),
	}

	for name, model := range testCases {
		t.Run(name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx := mocks.NewMockContext(w, nil)

			renderFoundRecords(ctx, model)

			testHelpers.AssertStatusCode(t, w.Code, http.StatusOK)

			errs := testHelpers.UnmarshalErrors(t, w)
			if len(errs) != 0 {
				t.Errorf("got unexpected errors: %s", errs)
			}
		})
	}
}

func TestRenderError(t *testing.T) {
	testCases := map[string]int{
		ErrNotFound:             http.StatusNotFound,
		ErrBadRequest:           http.StatusBadRequest,
		testHelpers.DatabaseErr: http.StatusInternalServerError,
	}

	for name, code := range testCases {
		t.Run(name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx := mocks.NewMockContext(w, nil)

			renderError(ctx, errors.New(name))

			testHelpers.AssertStatusCode(t, w.Code, code)

			errs := testHelpers.UnmarshalErrors(t, w)
			if errs[0] != name {
				t.Errorf("failed to return errors")
			}
		})
	}
}

func TestRenderErrors(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := mocks.NewMockContext(w, nil)

	mockErrors := []error{
		errors.New("first error"),
		errors.New("second error"),
	}
	renderErrors(ctx, mockErrors)

	if w.Code != http.StatusBadRequest {
		t.Errorf("got %d want %d", w.Code, http.StatusBadRequest)
	}

	errs := testHelpers.UnmarshalErrors(t, w)
	if errs[0] != "first error" || errs[1] != "second error" {
		t.Errorf("failed to return errors")
	}
}

package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"upcourse/internal/mocks"
)

func TestRenderSuccess(t *testing.T) {
	testCases := map[string]int{
		"StatusOK":      http.StatusOK,
		"StatusCreated": http.StatusCreated,
	}

	for name, code := range testCases {
		t.Run(name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx := mocks.NewMockContext(w, nil)

			renderSuccess(ctx, code)

			if w.Code != code {
				t.Errorf("got %d want %d", w.Code, code)
			}
		})
	}
}

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

			if w.Code != http.StatusOK {
				t.Errorf("got %d want %d", w.Code, http.StatusOK)
			}

			response := unmarshalResponse(t, w.Body)

			data := reflect.ValueOf(response.Data)
			if data.Kind() != reflect.Map && data.Kind() != reflect.Slice {
				t.Errorf("failed to return a record, got %s", data)
			}

			errors := response.Errors
			if len(errors) > 0 {
				t.Errorf("unexpected errors: %s", errors)
			}
		})
	}
}

func TestRenderError(t *testing.T) {
	testCases := map[string]int{
		"record not found": http.StatusNotFound,
		"other error":      http.StatusBadRequest,
	}

	for name, code := range testCases {
		t.Run(name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx := mocks.NewMockContext(w, nil)

			renderError(ctx, errors.New(name))

			if w.Code != code {
				t.Errorf("got %d want %d", w.Code, code)
			}

			response := unmarshalResponse(t, w.Body)

			errors := response.Errors
			if errors[0] != name {
				t.Errorf("failed to return errors")
			}

			data := reflect.ValueOf(response.Data)
			if data.Kind() == reflect.Map {
				t.Errorf("unexpected data: %s", data)
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

	response := unmarshalResponse(t, w.Body)

	errors := response.Errors
	if errors[0] != "first error" || errors[1] != "second error" {
		t.Errorf("failed to return errors")
	}

	data := reflect.ValueOf(response.Data)
	if data.Kind() == reflect.Map && data.Kind() == reflect.Slice {
		t.Errorf("unexpected data: %s", data)
	}
}

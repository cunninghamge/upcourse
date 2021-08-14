package helpers

import (
	"encoding/json"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/google/jsonapi"
)

func UnmarshalPayload(t *testing.T, w *httptest.ResponseRecorder, model interface{}) interface{} {
	t.Helper()

	err := jsonapi.UnmarshalPayload(w.Body, model)
	if err != nil {
		t.Errorf("error unmarshaling json response: %v", err)
	}
	return model
}

func UnmarshalManyPayload(t *testing.T, w *httptest.ResponseRecorder, ty interface{}) []interface{} {
	t.Helper()

	response, err := jsonapi.UnmarshalManyPayload(w.Body, reflect.TypeOf(ty))
	if err != nil {
		t.Errorf("error unmarshaling json response: %v", err)
	}
	return response
}

func UnmarshalErrors(t *testing.T, w *httptest.ResponseRecorder) []string {
	t.Helper()

	var response struct{ Errors []string }
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("error unmarshaling json response: %v", err)
	}

	return response.Errors
}

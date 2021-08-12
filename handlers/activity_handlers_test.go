package handlers

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"upcourse/config"
	"upcourse/internal/mocks"
	"upcourse/models"
)

func TestGetActivities(t *testing.T) {
	activities := mocks.DefaultActivities()

	t.Run("returns a list of the default activities", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{})

		GetActivities(ctx)

		if w.Code != 200 {
			t.Errorf("expected response code to be 200, got %d", w.Code)
		}

		var response struct {
			Data []SerializedResource
		}
		if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
			t.Errorf("error unmarshaling json response: %v", err)
		}

		for i, activity := range response.Data {
			assertResponseValue(t, activity.Type, "activity", "activity type")
			assertResponseValue(t, activity.ID, activities[i].ID, "Id")
		}
	})

	t.Run("does not include custom activities", func(t *testing.T) {
		config.Conn.Create(&models.Activity{Custom: true})
		defer teardown()

		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{})

		GetActivities(ctx)

		var response struct {
			Data []SerializedResource
		}
		if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
			t.Errorf("error unmarshaling json response: %v", err)
		}

		assertResponseValue(t, len(response.Data), len(activities), "number of activities")
	})

	t.Run("returns database errors if they occur", func(t *testing.T) {
		err := "some database error"
		config.Conn.Error = errors.New(err)
		defer func() {
			config.Conn.Error = nil
		}()

		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{})

		GetActivities(ctx)

		response := unmarshalResponse(t, w.Body)
		if response.Errors[0] != err {
			t.Errorf("got %s want %s for error message", response.Errors[0], err)
		}
	})
}

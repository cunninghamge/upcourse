package handlers

import (
	"net/http"
	"reflect"
	"testing"

	"upcourse/config"
	testHelpers "upcourse/internal/helpers"
	"upcourse/internal/mocks"
	"upcourse/models"
)

func TestGetActivities(t *testing.T) {
	activities := mocks.DefaultActivities()

	t.Run("returns a list of the default activities", func(t *testing.T) {
		w := testHelpers.NewRequest(nil, "", GetActivities)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusOK)

		response := testHelpers.UnmarshalManyPayload(t, w, new(models.Activity))
		for i := 0; i < len(activities); i++ {
			got, ok := response[i].(*models.Activity)
			if !ok {
				t.Errorf("error casting response element as activity")
			}
			want := activities[i]
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v for response activities[%d]", got, want, i)
			}
		}
	})

	t.Run("does not include custom activities", func(t *testing.T) {
		config.Conn.Create(&models.Activity{Custom: true})
		defer testHelpers.Teardown()

		w := testHelpers.NewRequest(nil, "", GetActivities)

		response := testHelpers.UnmarshalManyPayload(t, w, new(models.Activity))
		if len(response) != len(activities) {
			t.Errorf("got %d want %d for number of results", len(response), len(activities))
		}
	})

	t.Run("returns database errors if they occur", func(t *testing.T) {
		testHelpers.ForceError()
		defer testHelpers.ClearError()

		w := testHelpers.NewRequest(nil, "", GetActivities)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusInternalServerError)

		response := testHelpers.UnmarshalErrors(t, w)
		if response[0] != testHelpers.DatabaseErr {
			t.Errorf("got %s want %s for error message", response[0], testHelpers.DatabaseErr)
		}
	})
}

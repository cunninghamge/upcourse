package handlers

import (
	"net/http"
	"testing"

	"upcourse/models"
)

func TestGetActivities(t *testing.T) {
	t.Run("when GetActivities succeeds", func(t *testing.T) {
		w := newRequest(GetActivities, nil, "")

		assertStatusCode(t, w.Code, http.StatusOK)

		unmarshalPayload(t, w.Body, new(models.Activity), many)
	})

	t.Run("when GetActivities fails", func(t *testing.T) {
		forceError()
		defer clearError()

		w := newRequest(GetActivities, nil, "")

		assertStatusCode(t, w.Code, http.StatusInternalServerError)

		unmarshalPayload(t, w.Body, new(error), many)
	})
}

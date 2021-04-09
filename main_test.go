package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETCourses(t *testing.T) {
	t.Run("returns the name of a course", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/courses/1", nil)
		response := httptest.NewRecorder()

		Server(response, request)

		got := response.Body.String()
		want := "Nursing 101"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGETCourses(t *testing.T) {
	t.Run("returns the name of a course", func(t *testing.T) {
		router := setupRouter()
		request, _ := http.NewRequest("GET", "/courses/1", nil)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "Nursing 101", response.Body.String())
	})
}

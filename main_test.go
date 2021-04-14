package main

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"

	"course-chart/config"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	config.Connect()
	code := m.Run()
	os.Exit(code)
}

func assertResponseValue(t *testing.T, got, want interface{}, field string) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v for field %q", got, want, field)
	}
}

type MarshaledError struct {
	Status int
	Errors string
}

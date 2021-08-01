package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/gin-gonic/gin"

	"upcourse/config"
	"upcourse/models"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	config.Connect()
	code := m.Run()
	os.Exit(code)
}

func teardown() {
	config.Conn.Unscoped().Where("1=1").Delete(&models.ModuleActivity{})
	config.Conn.Unscoped().Where("1=1").Delete(&models.Module{})
	config.Conn.Unscoped().Where("1=1").Delete(&models.Course{})
	config.Conn.Unscoped().Where("custom=true").Delete(&models.Activity{})
}

func assertResponseValue(t *testing.T, got, want interface{}, field string) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v for field %q", got, want, field)
	}
}

func unmarshalSuccessResponse(t *testing.T, r io.Reader) SuccessResponse {
	t.Helper()

	body, _ := ioutil.ReadAll(r)
	successResponse := SuccessResponse{}
	err := json.Unmarshal([]byte(body), &successResponse)

	if err != nil {
		t.Errorf("Error marshaling JSON response\nError: %v", err)
	}

	return successResponse
}

type SuccessResponse struct {
	Status  int
	Message string
}

func unmarshalErrorResponse(t *testing.T, r io.Reader) ErrorResponse {
	t.Helper()

	body, _ := ioutil.ReadAll(r)
	responseError := ErrorResponse{}
	err := json.Unmarshal([]byte(body), &responseError)
	if err != nil {
		t.Errorf("Error marshaling JSON response\nError: %v", err)
	}

	return responseError
}

type ErrorResponse struct {
	Status int
	Errors string
}

func unmarshalMultipleErrorResponse(t *testing.T, r io.Reader) MultipleErrorResponse {
	t.Helper()

	body, _ := ioutil.ReadAll(r)
	responseErrors := MultipleErrorResponse{}
	err := json.Unmarshal([]byte(body), &responseErrors)
	if err != nil {
		t.Errorf("Error marshaling JSON response\nError: %v", err)
	}

	return responseErrors
}

type MultipleErrorResponse struct {
	Status int
	Errors []string
}

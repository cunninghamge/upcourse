package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
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
	config.Conn.Where("1=1").Delete(&models.ModuleActivity{})
	config.Conn.Where("1=1").Delete(&models.Module{})
	config.Conn.Where("1=1").Delete(&models.Course{})
	config.Conn.Where("custom=true").Delete(&models.Activity{})
}

func assertResponseValue(t *testing.T, got, want interface{}, field string) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v for field %q", got, want, field)
	}
}

type Response struct {
	Data   interface{}
	Errors []string
}

func unmarshalResponse(t *testing.T, r io.Reader) Response {
	t.Helper()

	body, _ := ioutil.ReadAll(r)
	response := Response{}
	err := json.Unmarshal([]byte(body), &response)

	if err != nil {
		t.Errorf("Error marshaling JSON response\nError: %v", err)
	}

	return response
}

package handlers

import (
	"log"
	"os"
	"testing"

	"github.com/gin-gonic/gin"

	"upcourse/config"
	testHelpers "upcourse/internal/helpers"
)

func TestMain(m *testing.M) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("recovered from error: %v", err)
			testHelpers.Teardown()
		}
	}()

	gin.SetMode(gin.TestMode)
	if err := config.Connect(); err != nil {
		log.Fatal("could not connect to test database")
	}

	code := m.Run()
	os.Exit(code)
}

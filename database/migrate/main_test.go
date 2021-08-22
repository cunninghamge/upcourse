package main

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
	db "upcourse/database"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	if err := db.Connect(); err != nil {
		log.Fatal("could not connect to test database")
	}

	code := m.Run()

	os.Exit(code)
}

func TestPackageMain(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)

	main()

	messages := []string{
		"Completed automigration of database models",
		"Migration complete",
	}
	for _, message := range messages {
		if !strings.Contains(buf.String(), message) {
			t.Errorf("expected but did not receive output message: %s", message)
		}
	}
}

package models

import (
	"errors"
	"log"
	"os"
	"testing"
	db "upcourse/database"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	if err := db.Connect(); err != nil {
		log.Fatal("could not connect to the database")
	}

	code := m.Run()
	os.Exit(code)
}

const testDBErr = "some database error"

func forceError() {
	db.Conn.Error = errors.New(testDBErr)
}

func clearError() {
	db.Conn.Error = nil
}

func teardown() {
	db.Conn.Where("custom=true").Delete(&Activity{})
	db.Conn.Where("1=1").Delete(&Course{})
}

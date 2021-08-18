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

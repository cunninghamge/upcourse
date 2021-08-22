package main

import (
	"testing"
	db "upcourse/database"
	"upcourse/models"

	"github.com/gin-gonic/gin"
)

func TestSeedTestMode(t *testing.T) {
	var count int64
	db.Conn.Model(&models.Activity{}).Count(&count)

	if count != 14 {
		t.Errorf("expected 14 activities, got %d", count)
	}

	tx := db.Conn.Where(&sampleCourse).First(&models.Course{})
	if tx.Error == nil {
		t.Error("expected not to find sample course in test database but did")
	}
}

func TestSeedDebugMode(t *testing.T) {
	t.Skip() //can't run debug mode in CircleCI
	gin.SetMode(gin.DebugMode)
	db.Connect()
	defer func() {
		gin.SetMode(gin.TestMode)
		db.Connect()
	}()

	var count int64
	db.Conn.Model(&models.Activity{}).Count(&count)

	if count != 14 {
		t.Errorf("expected 14 activities, got %d", count)
	}

	var course models.Course
	if err := db.Conn.First(&course, 1).Error; err != nil {
		t.Error("failed to find sample course in database")
	}
}

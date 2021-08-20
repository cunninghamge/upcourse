package main

import (
	"fmt"
	"testing"
	db "upcourse/database"
	"upcourse/models"

	"github.com/gin-gonic/gin"
)

func TestSeedTestMode(t *testing.T) {
	db.Conn.Where("1=1").Delete(&models.Activity{})

	err := seed()
	fmt.Printf("err: %v\n", err)
	var count int64
	db.Conn.Model(&models.Activity{}).Count(&count)

	if count != 14 {
		t.Errorf("expected 14 activities, got %d", count)
	}

	var courseCount int64
	db.Conn.Model(&models.Course{}).Count(&courseCount)

	if courseCount != 0 {
		t.Errorf("expected 0 courses, got %d", courseCount)
	}
}

func TestSeedDebugMode(t *testing.T) {
	gin.SetMode(gin.DebugMode)

	seed()

	var count int64
	db.Conn.Model(&models.Activity{}).Count(&count)

	if count != 14 {
		t.Errorf("expected 14 activities, got %d", count)
	}

	var course models.Course
	if err := db.Conn.First(&course, 1).Error; err != nil {
		t.Error("failed to find sample course in database")
	}

	gin.SetMode(gin.TestMode)
	db.Connect()
}

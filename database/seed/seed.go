package main

import (
	"log"

	db "upcourse/database"
	"upcourse/models"

	"github.com/gin-gonic/gin"
)

func seed() error {
	if err := db.Connect(); err != nil {
		return err
	}

	if err := createDefaultActivities(); err != nil {
		return err
	}
	log.Println("Created default activities")

	if gin.Mode() != gin.TestMode {
		if err := createSampleCourse(); err != nil {
			return err
		}
		log.Println("Created sample course")
	}

	return nil
}

func createDefaultActivities() error {
	for _, activity := range defaultActivities {
		if err := db.Conn.FirstOrCreate(&models.Activity{}, activity).Error; err != nil {
			return err
		}
	}
	return nil
}

func createSampleCourse() error {
	if err := db.Conn.First(&models.Course{}, 1).Error; err == nil {
		if err := db.Conn.FirstOrCreate(&sampleCourse).Error; err != nil {
			return err
		}
	}
	return nil
}

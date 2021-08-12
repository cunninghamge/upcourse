package main

import (
	"log"
	"upcourse/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GormMigration struct {
	ID      int
	Version int64
	db      *gorm.DB `gorm:"-"`
}

func (m *GormMigration) execute() {
	if err := m.autoMigrate(); err != nil {
		log.Fatalf("Migration error: %v", err)
	}
	log.Println("Completed automigration of database models")

	if err := m.setConstraints(); err != nil {
		log.Fatalf("Migration error: %v", err)
	}
	log.Println("Completed database indexes")

	if err := m.createDefaultActivities(); err != nil {
		log.Fatalf("Error creating default activities: %v", err)
	}
	log.Println("Created default activities")

	if gin.Mode() != gin.TestMode {
		if err := m.createSampleCourse(); err != nil {
			log.Fatalf("Error creating sample course: %v", err)
		}
		log.Println("Created sample course")
	}

	if err := m.db.Create(&m).Error; err != nil {
		log.Fatalf("Error updating gorm_migration table: %v", err)
	}
}

func (m GormMigration) autoMigrate() error {
	return m.db.AutoMigrate(&models.Course{}, &models.Module{}, &models.ModuleActivity{}, &models.Activity{}, &GormMigration{})
}

func (m GormMigration) createDefaultActivities() error {
	for _, activity := range defaultActivities {
		if err := m.db.FirstOrCreate(&models.Activity{}, activity).Error; err != nil {
			return err
		}
	}
	return nil
}

func (m GormMigration) createSampleCourse() error {
	if err := m.db.First(&models.Course{}, 1).Error; err != nil {
		if err := m.db.Create(&sampleCourse).Error; err != nil {
			return err
		}
	}
	return nil
}

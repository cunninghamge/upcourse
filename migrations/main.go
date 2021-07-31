package main

import (
	"log"
	"upcourse/config"
	"upcourse/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type Migration struct {
	db *gorm.DB
}

func (m Migration) execute() error {
	if err := m.autoMigrate(); err != nil {
		log.Fatalf("Migration error: %v", err)
	}
	log.Println("Completed automigration of database models")

	if err := m.createDefaultActivities(); err != nil {
		log.Fatalf("Error creating default activities: %v", err)
	}
	log.Println("Created default activities")

	if gin.Mode() != gin.TestMode {
		if err := m.createSampleCourse(); err != nil {
			log.Fatalf("Error creating sample course")
		}
		log.Println("Created sample course")
	}
	return nil
}

func (m Migration) autoMigrate() error {
	return m.db.AutoMigrate(&models.Course{}, &models.Module{}, &models.ModuleActivity{}, &models.Activity{})
}

func (m Migration) createDefaultActivities() error {
	for _, activity := range defaultActivities {
		if err := m.db.FirstOrCreate(&models.Activity{}, activity).Error; err != nil {
			return err
		}
	}
	return nil
}

func (m Migration) createSampleCourse() error {
	if err := m.db.FirstOrCreate(&models.Course{}, sampleCourse).Error; err != nil {
		return err
	}
	return nil
}

func main() {
	godotenv.Load()

	if err := config.Connect(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	migration := Migration{config.Conn}

	if err := migration.execute(); err != nil {
		log.Fatalf("Error completing migration: %v", err)
	}
	log.Print("Migration complete")
}

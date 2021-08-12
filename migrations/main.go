package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/gorm"

	"upcourse/config"
	"upcourse/models"
)

// TODO: versioning
type Migration struct {
	db *gorm.DB
}

func (m Migration) execute() error {
	if err := m.autoMigrate(); err != nil {
		log.Fatalf("Migration error: %v", err)
	}
	log.Println("Completed automigration of database models")

	if err := m.createIndexes(); err != nil {
		log.Fatalf("Migration error: %v", err)
	}
	log.Println("Completed database indexes")

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
	if err := m.db.First(&models.Course{}, 1).Error; err != nil {
		m.db.Exec(`DELETE FROM courses WHERE id = 1;`)
		if err := m.db.Create(&sampleCourse).Error; err != nil {
			return err
		}
	}
	return nil
}

func (m Migration) createIndexes() error {
	if !m.db.Migrator().HasIndex(&models.ModuleActivity{}, "index_module_activities_on_activities_modules") {
		err := m.db.Exec(`CREATE UNIQUE INDEX index_module_activities_on_activities_modules ON module_activities (module_id, activity_id);`).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	if err := config.Connect(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	migration := Migration{config.Conn}

	if err := migration.execute(); err != nil {
		log.Fatalf("Error completing migration: %v", err)
	}
	log.Print("Migration complete")
}

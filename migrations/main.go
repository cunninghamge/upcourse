package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/gorm"

	"upcourse/config"
	"upcourse/models"
)

type GormMigration struct {
	ID      int
	Version int64
	db      *gorm.DB `gorm:"-"`
}

func (m *GormMigration) execute() error {
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
			log.Fatalf("Error creating sample course: %v", err)
		}
		log.Println("Created sample course")
	}

	if err := m.db.Create(&m).Error; err != nil {
		log.Fatalf("Error updating gorm_migration table: %v", err)
	}
	return nil
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

func (m GormMigration) createIndexes() error {
	if !m.db.Migrator().HasIndex(&models.ModuleActivity{}, "index_module_activities_on_activities_modules") {
		err := m.db.Exec(`CREATE UNIQUE INDEX index_module_activities_on_activities_modules ON module_activities (module_id, activity_id);`).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	if err := config.Connect(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	migration := GormMigration{db: config.Conn, Version: time.Now().Unix()}

	if err := migration.execute(); err != nil {
		log.Fatalf("Error completing migration: %v", err)
	}
	log.Printf("Completed migration to version %d", migration.ID)
}

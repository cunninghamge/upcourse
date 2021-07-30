package main

import (
	"upcourse/config"
	"upcourse/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()
	mode := gin.Mode()

	switch mode {
	case "release":
		migrate("release")
	case "test":
		migrate("test")
	default:
		migrate("test")
		migrate("default")
	}
}

func migrate(mode string) {
	// establish connection
	var gormDB *gorm.DB
	switch mode {
	case "release":
		gormDB = config.DBConnectRelease()
	case "test":
		gormDB = config.DBConnect("upcourse_test")
	default:
		gormDB = config.DBConnect("upcourse")
	}
	// run automigrate
	gormDB.AutoMigrate(&models.Course{}, &models.Module{}, &models.ModuleActivity{}, &models.Activity{})
	// set ON DELETE constraint to CASCADE
	set_constraints(gormDB)
	// set triggers and seed activities if none exist
	err := gormDB.First(&models.Activity{}, 1).Error
	if err != nil {
		set_triggers(gormDB)
		seed_activities(gormDB)
	}
	// seed default course
	err = gormDB.First(&models.Course{}, 1).Error
	if err != nil {
		seed_full_course(gormDB)
	}
}

package main

import (
	"log"
	db "upcourse/database"
	"upcourse/models"
)

func run() error {
	if err := db.Connect(); err != nil {
		return err
	}

	if err := autoMigrate(); err != nil {
		return err
	}
	log.Println("Completed automigration of database models")

	ran, err := createIndexes()
	if err != nil {
		return err
	} else if ran {
		log.Println("Completed creation of database indexes")
	}

	return nil
}

func autoMigrate() error {
	return db.Conn.AutoMigrate(&models.Course{}, &models.Module{}, &models.ModuleActivity{}, &models.Activity{})
}

func createIndexes() (bool, error) {
	var ran bool
	if !db.Conn.Migrator().HasIndex(&models.ModuleActivity{}, "index_module_activities_on_activities_modules") {
		ran = true
		err := db.Conn.Exec(`CREATE UNIQUE INDEX index_module_activities_on_activities_modules ON module_activities (module_id, activity_id);`).Error
		if err != nil {
			return ran, err
		}
	}
	if !db.Conn.Migrator().HasIndex(&models.Module{}, "index_modules_on_courses_number") {
		ran = true
		err := db.Conn.Exec(`CREATE UNIQUE INDEX index_modules_on_courses_number ON modules (course_id, number);`).Error
		if err != nil {
			return ran, err
		}
	}
	return false, nil
}

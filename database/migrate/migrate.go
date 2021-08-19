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

	if err := setConstraints(); err != nil {
		return err
	}
	log.Println("Completed creation of database indexes")

	return nil
}

func autoMigrate() error {
	return db.Conn.AutoMigrate(&models.Course{}, &models.Module{}, &models.ModuleActivity{}, &models.Activity{})
}

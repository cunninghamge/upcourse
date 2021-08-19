package main

import (
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Panicf("Error completing migration: %v", err)
	}

	log.Print("Migration complete")
}

// [x] remove versioning
// [x] extract as much as possible to migration.go
// [x] extract seeds to separate executable
// [x] change seeds to use first or create
// [ ] update circle & heroku config to run seeds
// [ ] write tests based on what the state of the database should be
// [ ] try using gorm.Model

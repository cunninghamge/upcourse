package main

import (
	"log"
	"time"

	_ "github.com/lib/pq"

	db "upcourse/database"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	migration := GormMigration{
		db:      db.Conn,
		Version: time.Now().Unix(),
	}

	migration.execute()

	log.Printf("Completed migration to version %d", migration.ID)
}

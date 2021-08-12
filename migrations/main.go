package main

import (
	"log"
	"time"

	_ "github.com/lib/pq"

	"upcourse/config"
)

func main() {
	if err := config.Connect(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	migration := GormMigration{
		db:      config.Conn,
		Version: time.Now().Unix(),
	}

	migration.execute()

	log.Printf("Completed migration to version %d", migration.ID)
}

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

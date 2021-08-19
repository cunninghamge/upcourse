package main

import "log"

func main() {
	if err := seed(); err != nil {
		log.Panicf("Error completing migration: %v", err)
	}

	log.Print("Migration complete")
}

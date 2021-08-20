package main

import "log"

func main() {
	if err := seed(); err != nil {
		log.Panicf("Error seeding database: %v", err)
	}

	log.Print("Seeding complete")
}

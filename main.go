package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	database "upcourse/config"
	"upcourse/routes"
)

func main() {
	godotenv.Load()

	port := ":" + os.Getenv("PORT")

	err := database.Connect()
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	router := routes.GetRoutes()

	log.Fatal(router.Run(port))
}

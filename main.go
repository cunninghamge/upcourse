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

	database.Connect()

	router := routes.GetRoutes()

	log.Fatal(router.Run(port))
}

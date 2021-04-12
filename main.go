package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	database "course-chart/config"
	"course-chart/routes"
)

func main() {
	godotenv.Load()

	port := ":" + os.Getenv("PORT")

	db := database.Connect()

	router := routes.Routes(db)

	log.Fatal(router.Run(port))
}

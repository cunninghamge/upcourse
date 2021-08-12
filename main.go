package main

import (
	"log"

	"upcourse/server"
)

func main() {
	server := server.NewServer()
	log.Fatal(server.Engine.Run(":" + server.Port))
}

package server

import (
	"log"
	"os"

	db "upcourse/database"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Port   string
	Engine *gin.Engine
}

func NewServer() Server {
	if err := db.Connect(); err != nil {
		log.Panicf("Error connecting to database: %q", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	server := Server{
		Port:   port,
		Engine: AppRouter(),
	}

	return server
}

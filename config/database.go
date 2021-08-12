package config

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func Connect() error {
	mode := gin.Mode()

	var dsn string
	switch mode {
	case "release":
		dsn = os.Getenv("DATABASE_URL") + "?sslmode=require"
	default:
		dsn = baseDSN(mode)
	}

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	Conn = gormDB
	log.Printf("Connected to database")
	return nil
}

func baseDSN(mode string) string {
	const (
		host = "localhost"
		port = 5432
		user = "postgres"
	)

	var dbName string
	switch mode {
	case "test":
		dbName = "upcourse_test"
	default:
		dbName = "upcourse"
	}

	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbName)
}

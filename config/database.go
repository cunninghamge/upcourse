package config

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const baseDSN = "host=localhost port=5432 user=postgres sslmode=disable dbname=upcourse"

var Conn *gorm.DB

func Connect() error {
	mode := gin.Mode()
	var dsn string

	switch mode {
	case "release":
		dsn = os.Getenv("DATABASE_URL") + "?sslmode=require"
	case "test":
		dsn = baseDSN + "_test"
	default:
		dsn = baseDSN
	}

	if _, err := DBConnect(dsn); err != nil {
		return err
	}

	log.Printf("Connected to database")
	return nil
}

func DBConnect(dsn string) (*gorm.DB, error) {
	Conn, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "postgres",
		DSN:        dsn,
	}), &gorm.Config{})

	return Conn, err
}

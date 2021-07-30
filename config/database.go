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

const (
	host       = "localhost"
	port       = 5432
	user       = "postgres"
	driverName = "postgres"
)

var Conn *gorm.DB

func Connect() error {
	mode := gin.Mode()
	var err error

	switch mode {
	case "release":
		Conn, err = DBConnectRelease()
	case "test":
		Conn, err = DBConnect("upcourse_test")
	default:
		Conn, err = DBConnect("upcourse")
	}
	if err != nil {
		return err
	}

	log.Printf("Connected to database")
	return nil
}

func DBConnectRelease() (*gorm.DB, error) {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: driverName,
		DSN:        os.Getenv("DATABASE_URL") + "?sslmode=require",
	}), &gorm.Config{})

	return gormDB, err
}

func DBConnect(dbname string) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: driverName,
		DSN:        psqlInfo,
	}), &gorm.Config{})

	return gormDB, err
}

package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func Connect() *gorm.DB {
	mode := gin.Mode()
	var gormDB *gorm.DB

	switch mode {
	case "release":
		gormDB = DBConnectRelease()
	case "test":
		gormDB = DBConnect("upcourse_test")
	default:
		gormDB = DBConnect("upcourse")
	}

	log.Printf("Connected to database")
	Conn = gormDB
	return gormDB
}

func DBConnectRelease() *gorm.DB {
	sqlDB, err := sql.Open("postgres", os.Getenv("DATABASE_URL")+"?sslmode=require")

	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	return gormDB
}

func DBConnect(dbname string) *gorm.DB {
	const (
		host = "localhost"
		port = 5432
		user = "postgres"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)

	sqlDB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	return gormDB
}

package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func Connect() *gorm.DB {
	mode := gin.Mode()

	switch mode {
	case "release":
		return DBConnectRelease()
	case "test":
		return DBConnectTest()
	default:
		return DBConnectDefault()
	}
}

func DBConnectRelease() *gorm.DB {
	sqlDB, err := sql.Open("postgres", os.Getenv("DATABASE_URL")+"?sslmode=require")

	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	Conn = gormDB

	return gormDB
}

func DBConnectDefault() *gorm.DB {
	const (
		host   = "localhost"
		port   = 5432
		user   = "postgres"
		dbname = "course_chart"
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

	Conn = gormDB

	return gormDB
}

func DBConnectTest() *gorm.DB {
	const (
		host   = "localhost"
		port   = 5432
		user   = "postgres"
		dbname = "course_chart_test"
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

	Conn = gormDB

	return gormDB
}

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

func Connect() *gorm.DB {
	mode := gin.Mode()

	switch mode {
	case "release":
		return DBConnectRelease()
	default:
		return SQLDBDefault()
	}
}

func DBConnectRelease() *gorm.DB {
	sqlDB, err := sql.Open("postgres", os.Getenv("DATABASE_URL")+"?sslmode=require")

	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	return gormDB
}

func SQLDBDefault() *gorm.DB {
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

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	return gormDB
}

package config

import (
	"log"
	"net/url"
	"os"

	"github.com/go-pg/pg"
)

func Connect() *pg.DB {
	// mode := gin.Mode()

	var pgOptions *pg.Options

	// switch mode {
	// case "release":
	pgOptions = PGOptionsRelease()
	// default:
	// 	pgOptions = PGOptionsDefault()
	// }

	db := pg.Connect(pgOptions)

	if db == nil {
		log.Printf("Could not connect to database")
		os.Exit(100)
	}

	log.Printf("Connected to database")

	return db
}

func PGOptionsRelease() *pg.Options {
	parsedUrl, err := url.Parse(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	pgOptions := &pg.Options{
		User:     parsedUrl.User.Username(),
		Database: parsedUrl.Path[1:] + "?sslmode=require",
		Addr:     parsedUrl.Host,
	}

	if password, ok := parsedUrl.User.Password(); ok {
		pgOptions.Password = password
	}

	return pgOptions
}

func PGOptionsDefault() *pg.Options {
	return &pg.Options{
		User:     os.Getenv("POSTGRES_USER"),
		Database: os.Getenv("POSTGRES_NAME"),
		Addr:     os.Getenv("POSTGRES_ADDRESS"),
	}
}

func PGOptionsTest() *pg.Options {
	return &pg.Options{
		User:     os.Getenv("POSTGRES_USER"),
		Database: "course_chart_test",
	}
}

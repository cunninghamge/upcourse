package config

import (
	"log"
	"net/url"
	"os"

	"github.com/go-pg/pg"
)

func Connect() *pg.DB {
	mode := os.Getenv("GIN_MODE")

	var pgOptions *pg.Options

	switch mode {
	case "release":
		pgOptions = pgOptionsRelease()
	default:
		pgOptions = pgOptionsDefault()
	}

	db := pg.Connect(pgOptions)

	if db == nil {
		log.Printf("Could not connect to database")
		os.Exit(100)
	}

	log.Printf("Connected to database")

	return db
}

func pgOptionsRelease() *pg.Options {
	parsedUrl, err := url.Parse(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	pgOptions := &pg.Options{
		User:     parsedUrl.User.Username(),
		Database: parsedUrl.Path[1:],
		Addr:     parsedUrl.Host,
	}

	if password, ok := parsedUrl.User.Password(); ok {
		pgOptions.Password = password
	}

	return pgOptions
}

func pgOptionsDefault() *pg.Options {
	pgOptions := &pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Addr:     os.Getenv("DB_ADDRESS"),
		Database: os.Getenv("DB_NAME"),
	}

	return pgOptions
}

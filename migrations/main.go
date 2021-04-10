package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"
	"github.com/joho/godotenv"
)

func main() {
	flag.Usage = usage
	flag.Parse()

	// migrate("course_chart")
	migrate("course_chart_test")
}

func migrate(db string) {
	godotenv.Load()

	database := pg.Connect(&pg.Options{
		User:     os.Getenv("POSTGRES_USER"),
		Database: db,
	})
	defer database.Close()

	oldVersion, newVersion, err := migrations.Run(database, flag.Args()...)
	if err != nil {
		exitf(err.Error())
	}
	if newVersion != oldVersion {
		fmt.Printf("migrated %q from version %d to %d\n", db, oldVersion, newVersion)
	} else {
		fmt.Printf("%q version is %d\n", db, oldVersion)
	}
}

func usage() {
	flag.PrintDefaults()
	os.Exit(2)
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}

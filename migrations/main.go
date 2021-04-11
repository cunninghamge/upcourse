package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"
	"github.com/joho/godotenv"

	"course-chart/config"
)

func main() {
	flag.Usage = usage
	flag.Parse()

	godotenv.Load()
	mode := gin.Mode()

	switch mode {
	case "test":
		migrate("test")
	case "release":
		migrate("release")
	default:
		migrate("test")
		migrate("default")
	}
}

func migrate(db string) {
	var pgOptions *pg.Options
	switch db {
	case "test":
		pgOptions = config.PGOptionsTest()
	case "release":
		pgOptions = config.PGOptionsRelease()
	default:
		pgOptions = config.PGOptionsDefault()
	}

	database := pg.Connect(pgOptions)
	defer database.Close()

	oldVersion, newVersion, err := migrations.Run(database, flag.Args()...)
	if err != nil {
		exitf(err.Error())
	}
	if newVersion != oldVersion {
		fmt.Printf("migrated %q from version %d to %d\n", pgOptions.Database, oldVersion, newVersion)
	} else {
		fmt.Printf("%q version is %d\n", pgOptions.Database, oldVersion)
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

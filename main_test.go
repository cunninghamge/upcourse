package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-pg/pg"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"course-chart/config"
)

func TestGETCourses(t *testing.T) {
	db := dbConnect()

	t.Run("returns the name of a course", func(t *testing.T) {
		course := &Course{
			Id:   1,
			Name: "Nursing 101",
		}

		db.Insert(course)
		defer db.Delete(course)

		router := setupRouter(db)
		request, _ := http.NewRequest("GET", fmt.Sprintf("/courses/%d", course.Id), nil)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, course.Name, response.Body.String())
	})
}

func dbConnect() *pg.DB {
	godotenv.Load()
	db := pg.Connect(config.PGOptionsTest())

	if db == nil {
		log.Printf("Could not connect to database")
		os.Exit(100)
	}

	log.Printf("Connected to database")

	return db
}

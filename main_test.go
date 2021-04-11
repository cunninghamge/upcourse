package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGETCourses(t *testing.T) {
	db := dbConnect()

	t.Run("returns the name of a course", func(t *testing.T) {
		course := &Course{
			Id:   1,
			Name: "Foundations of Nursing",
		}

		router := setupRouter(db)
		request, _ := http.NewRequest("GET", fmt.Sprintf("/courses/%d", course.Id), nil)
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, course.Name, response.Body.String())
	})
}

func dbConnect() *gorm.DB {
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

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	return gormDB
}

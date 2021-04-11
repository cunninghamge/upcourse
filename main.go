package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Course struct {
	Id          int
	Name        string
	Institution string
	CreditHours int
	Length      int
	CreatedAt   string
	UpdatedAt   string
}

func setupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()
	r.GET("/courses/1", func(c *gin.Context) {

		// course := &Course{Id: 1}
		_, err := db.Query("SELECT * FROM courses WHERE id = 1")
		if err != nil {
			panic(err)
		}

		c.String(200, "connected to db with no error")
	})
	return r
}

func main() {
	godotenv.Load()

	port := ":" + os.Getenv("PORT")

	// db := database.Connect()
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL")+"?sslmode=require")
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	router := setupRouter(db)

	log.Fatal(router.Run(port))
}

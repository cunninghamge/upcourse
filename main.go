package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/gorm"

	database "course-chart/config"
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

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.GET("/courses/1", func(c *gin.Context) {

		var course Course
		result := db.First(&course, 1)

		log.Print(result)
		log.Print(course)

		c.String(200, course.Name)
	})
	return r
}

func main() {
	godotenv.Load()

	port := ":" + os.Getenv("PORT")

	db := database.Connect()

	router := setupRouter(db)

	log.Fatal(router.Run(port))
}

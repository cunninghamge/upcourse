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

		// course := &Course{Id: 1}
		// _, err := db.Query("SELECT * FROM courses WHERE id = 1")
		// if err != nil {
		// 	panic(err)
		// }

		c.String(200, "still working")
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

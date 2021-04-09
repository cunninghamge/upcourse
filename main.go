package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
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

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/courses/1", func(c *gin.Context) {
		db := pg.Connect(&pg.Options{
			User:     "postgres",
			Database: "course_chart",
		})
		defer db.Close()

		course := &Course{Id: 1}
		err := db.Model(course).WherePK().Select()
		if err != nil {
			panic(err)
		}

		c.String(200, course.CreatedAt)
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}

package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/courses/1", func(c *gin.Context) {
		c.String(200, "Nursing 101")
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}

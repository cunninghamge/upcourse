package routes

import (
	"course-chart/handlers"

	"github.com/gin-gonic/gin"
)

func GetRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/courses/:id", handlers.GetCourse)
	return router
}

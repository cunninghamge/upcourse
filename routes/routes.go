package routes

import (
	"course-chart/handlers"

	"github.com/gin-gonic/gin"
)

func GetRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/courses/1", handlers.GetCourse)
	return router
}

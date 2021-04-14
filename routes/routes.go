package routes

import (
	"course-chart/handlers"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func GetRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/courses", handlers.GetCourses)
	router.GET("/courses/:id", handlers.GetCourse)
	router.GET("/modules/:id", handlers.GetModule)
	return router
}

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
	router.POST("/courses", handlers.CreateCourse)
	router.GET("/courses/:id", handlers.GetCourse)
	router.GET("/modules/:id", handlers.GetModule)
	router.GET("/activities", handlers.GetActivities)

	return router
}

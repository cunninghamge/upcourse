package routes

import (
	"upcourse/handlers"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func GetRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/courses", handlers.GetCourses)
	router.GET("/courses/:id", handlers.GetCourse)
	router.POST("/courses", handlers.CreateCourse)
	router.PATCH("/courses/:id", handlers.UpdateCourse)
	router.DELETE("/courses/:id", handlers.DeleteCourse)
	router.POST("/modules", handlers.CreateModule)
	router.GET("/modules/:id", handlers.GetModule)
	router.PATCH("/modules/:id", handlers.UpdateModule)
	router.DELETE("/modules/:id", handlers.DeleteModule)
	router.GET("/activities", handlers.GetActivities)

	return router
}

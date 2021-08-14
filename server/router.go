package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"upcourse/handlers"
)

func AppRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	v1 := router.Group("/v1")
	{
		v1.GET("/courses", handlers.GetCourses)
		v1.GET("/courses/:id", handlers.GetCourse)
		v1.POST("/courses", handlers.CreateCourse)
		v1.PATCH("/courses/:id", handlers.UpdateCourse)
		v1.DELETE("/courses/:id", handlers.DeleteCourse)
		v1.POST("/courses/:id/modules", handlers.CreateModule)
		v1.GET("/modules/:id", handlers.GetModule)
		v1.PATCH("/modules/:id", handlers.UpdateModule)
		v1.DELETE("/modules/:id", handlers.DeleteModule)
		v1.GET("/activities", handlers.GetActivities)
	}

	return router
}

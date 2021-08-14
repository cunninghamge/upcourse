package server

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAppRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	expectedRoutes := gin.RoutesInfo{
		{
			Method:  "GET",
			Path:    "/v1/courses",
			Handler: "upcourse/handlers.GetCourses",
		}, {
			Method:  "GET",
			Path:    "/v1/courses/:id",
			Handler: "upcourse/handlers.GetCourse",
		}, {
			Method:  "POST",
			Path:    "/v1/courses",
			Handler: "upcourse/handlers.CreateCourse",
		}, {
			Method:  "PATCH",
			Path:    "/v1/courses/:id",
			Handler: "upcourse/handlers.UpdateCourse",
		}, {
			Method:  "DELETE",
			Path:    "/v1/courses/:id",
			Handler: "upcourse/handlers.DeleteCourse",
		}, {
			Method:  "POST",
			Path:    "/v1/courses/:id/modules",
			Handler: "upcourse/handlers.CreateModule",
		}, {
			Method:  "GET",
			Path:    "/v1/modules/:id",
			Handler: "upcourse/handlers.GetModule",
		}, {
			Method:  "PATCH",
			Path:    "/v1/modules/:id",
			Handler: "upcourse/handlers.UpdateModule",
		}, {
			Method:  "DELETE",
			Path:    "/v1/modules/:id",
			Handler: "upcourse/handlers.DeleteModule",
		}, {
			Method:  "GET",
			Path:    "/v1/activities",
			Handler: "upcourse/handlers.GetActivities",
		},
	}

	router := AppRouter()
	routes := router.Routes()

	for _, er := range expectedRoutes {
		routeIndex := -1
		for i, r := range routes {
			if er.Method == r.Method && er.Path == r.Path && er.Handler == r.Handler {
				routeIndex = i
			}
		}

		if routeIndex >= 0 {
			routes[routeIndex] = routes[len(routes)-1]
			routes = routes[:len(routes)-1]
		} else {
			t.Errorf("missing route: expected to find route %s %s but did not", er.Method, er.Path)
		}
	}

	if len(routes) > 0 {
		var extraRoutes []string
		for _, r := range routes {
			extraRoutes = append(extraRoutes, fmt.Sprintf("%s %s", r.Method, r.Path))
		}
		t.Errorf("unexpected route(s):\n%s", strings.Join(extraRoutes, "\n"))
	}
}

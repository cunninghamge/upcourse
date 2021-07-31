package routes

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	expectedRoutes := gin.RoutesInfo{
		{
			Method:  "GET",
			Path:    "/courses",
			Handler: "upcourse/handlers.GetCourses",
		}, {
			Method:  "GET",
			Path:    "/courses/:id",
			Handler: "upcourse/handlers.GetCourse",
		}, {
			Method:  "POST",
			Path:    "/courses",
			Handler: "upcourse/handlers.CreateCourse",
		}, {
			Method:  "PATCH",
			Path:    "/courses/:id",
			Handler: "upcourse/handlers.UpdateCourse",
		}, {
			Method:  "DELETE",
			Path:    "/courses/:id",
			Handler: "upcourse/handlers.DeleteCourse",
		}, {
			Method:  "POST",
			Path:    "/modules",
			Handler: "upcourse/handlers.CreateModule",
		}, {
			Method:  "GET",
			Path:    "/modules/:id",
			Handler: "upcourse/handlers.GetModule",
		}, {
			Method:  "PATCH",
			Path:    "/modules/:id",
			Handler: "upcourse/handlers.UpdateModule",
		}, {
			Method:  "DELETE",
			Path:    "/modules/:id",
			Handler: "upcourse/handlers.DeleteModule",
		}, {
			Method:  "GET",
			Path:    "/activities",
			Handler: "upcourse/handlers.GetActivities",
		},
	}

	router := GetRoutes()
	routes := router.Routes()

	for _, expectedRoute := range expectedRoutes {
		routePresent, idx := routePresent(t, expectedRoute, routes)
		if routePresent {
			routes[idx] = routes[len(routes)-1]
			routes = routes[:len(routes)-1]
		}
		if !routePresent {
			t.Errorf("missing route: expected to find route %s %s but did not", expectedRoute.Method, expectedRoute.Path)
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

func routePresent(t *testing.T, er gin.RouteInfo, routes gin.RoutesInfo) (bool, int) {
	for i, r := range routes {
		if er.Method == r.Method && er.Path == r.Path && er.Handler == r.Handler {
			return true, i
		}
	}
	return false, -1
}
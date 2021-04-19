package testing

import (
	"course-chart/config"
	"course-chart/handlers"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetActiviesHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	config.Connect()

	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	handlers.GetActivities(context)

	if writer.Code != 200 {
		t.Errorf("got status code %d, wanted 200", writer.Code)
	}
}

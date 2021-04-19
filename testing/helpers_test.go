package testing

import (
	"course-chart/handlers"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRenderSuccess(t *testing.T) {
	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	handlers.RenderSuccess(context, "this function works")

	if writer.Code != 200 {
		t.Errorf("got status code %d, wanted 200", writer.Code)
	}

	expected := `{"message":"this function works","status":200}`
	if writer.Body.String() != expected {
		t.Errorf("got %q, wanted %q", writer.Body.String(), expected)
	}
}

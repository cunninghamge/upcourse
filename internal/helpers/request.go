package helpers

import (
	"net/http/httptest"
	"upcourse/internal/mocks"

	"github.com/gin-gonic/gin"
)

func NewRequest(params map[string]string, body string, fn func(*gin.Context)) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	ctx := mocks.NewMockContext(w, params)
	if body != "" {
		mocks.SetRequestBody(ctx, body)
	}

	fn(ctx)

	return w
}

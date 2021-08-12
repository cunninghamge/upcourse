package mocks

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
)

func NewMockContext(w *httptest.ResponseRecorder, params map[string]string) *gin.Context {
	ctx, _ := gin.CreateTestContext(w)
	for k, v := range params {
		ctx.Params = append(ctx.Params, gin.Param{Key: k, Value: v})
	}

	return ctx
}

func SetRequestBody(c *gin.Context, json string) *gin.Context {
	reader := strings.NewReader(json)
	c.Request = &http.Request{
		Header: make(http.Header),
		Body:   io.NopCloser(reader),
	}

	return c
}

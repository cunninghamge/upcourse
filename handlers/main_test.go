package handlers

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/jsonapi"

	db "upcourse/database"
	testHelpers "upcourse/internal/helpers"
	"upcourse/models"
)

func TestMain(m *testing.M) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("recovered from error: %v", err)
			testHelpers.Teardown()
		}
	}()

	gin.SetMode(gin.TestMode)
	if err := db.Connect(); err != nil {
		log.Fatal("could not connect to test database")
	}

	code := m.Run()
	os.Exit(code)
}

const testDBErr = "some database error"

func forceError() {
	db.Conn.Error = errors.New(testDBErr)
}

func clearError() {
	db.Conn.Error = nil
}

func teardown() {
	db.Conn.Where("custom=true").Delete(&models.Activity{})
}

func newRequest(fn func(*gin.Context), params map[string]string, body string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	for k, v := range params {
		ctx.Params = append(ctx.Params, gin.Param{Key: k, Value: v})
	}
	if body != "" {
		ctx.Request = &http.Request{
			Header: make(http.Header),
			Body:   io.NopCloser(strings.NewReader(body)),
		}
	}
	fn(ctx)
	return w
}

const many = true

func unmarshalPayload(t *testing.T, body *bytes.Buffer, model interface{}, many bool) {
	t.Helper()

	var err error
	if many {
		_, err = jsonapi.UnmarshalManyPayload(body, reflect.TypeOf(model))
	} else {
		err = jsonapi.UnmarshalPayload(body, model)
	}

	if err != nil {
		t.Errorf("error unmarshaling json response: %v", err)
	}
}

func assertStatusCode(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got status code %d wanted %d", got, want)
	}
}

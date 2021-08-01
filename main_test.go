package main

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(t *testing.T) {
	t.Run("can't connect to release mode in test", func(t *testing.T) {
		gin.SetMode(gin.ReleaseMode)

		var err interface{}
		defer func() {
			err = recover()
		}()

		main()

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

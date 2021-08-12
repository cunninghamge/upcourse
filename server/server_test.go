package server

import (
	"os"
	"testing"

	"upcourse/config"

	"github.com/gin-gonic/gin"
)

func TestNewServer(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("it creates a server with a specified port", func(t *testing.T) {
		port := "4040"
		os.Setenv("PORT", port)
		server := NewServer()

		if server.Port != port {
			t.Errorf("got %s want %s", server.Port, port)
		}
	})

	t.Run("it uses port 3000 by default", func(t *testing.T) {
		port := "3000"
		os.Setenv("PORT", "")
		server := NewServer()

		if server.Port != port {
			t.Errorf("got %s want %s", server.Port, port)
		}
	})

	t.Run("it connects to the database", func(t *testing.T) {
		NewServer()
		db, err := config.Conn.DB()
		if err != nil {
			t.Errorf("error connecting to the database: %v", err)
		}
		if err = db.Ping(); err != nil {
			t.Errorf("error connecting to the database: %v", err)
		}
	})

	t.Run("it returns and error if it can't connect to the database", func(t *testing.T) {
		gin.SetMode(gin.ReleaseMode)
		os.Setenv("DATABASE_URL", "foo/bar")
		var err interface{}
		defer func() {
			err = recover()
		}()

		NewServer()

		if err == nil {
			t.Errorf("expected database connection error but did not get one")
		}
	})
}

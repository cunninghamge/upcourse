package config

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestConnect(t *testing.T) {
	testCases := map[string]struct {
		mode      string
		wantError bool
	}{
		"release mode": {
			mode:      gin.ReleaseMode,
			wantError: true,
		},
		"test mode": {
			mode:      gin.TestMode,
			wantError: false,
		},
		"debug mode": {
			mode:      gin.DebugMode,
			wantError: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			gin.SetMode(tc.mode)

			err := Connect()
			if tc.wantError && err == nil {
				t.Error("expected an error but didn't get one")
			} else if !tc.wantError && err != nil {
				t.Errorf("error connecting to database, got err: %v", err)
			}
		})
	}

	gin.SetMode(gin.TestMode)
	Connect() //nolint
}

func TestDBConnect(t *testing.T) {
	testCases := map[string]struct {
		dsn       string
		wantError bool
	}{
		"release mode": {
			dsn:       os.Getenv("DATABASE_URL") + "?sslmode=require",
			wantError: true,
		},
		"test mode": {
			dsn:       "host=localhost port=5432 user=postgres sslmode=disable dbname=upcourse_test",
			wantError: false,
		},
		"debug mode": {
			dsn:       "host=localhost port=5432 user=postgres sslmode=disable dbname=upcourse",
			wantError: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			_, err := DBConnect(tc.dsn)
			if tc.wantError && err == nil {
				t.Error("expected an error but didn't get one")
			} else if !tc.wantError && err != nil {
				t.Errorf("error connecting to database, got err: %v", err)
			}
		})
	}
}

package config

import (
	"fmt"
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
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			gin.SetMode(tc.mode)

			err := Connect()
			if tc.wantError && err == nil {
				t.Error("expected an error connecting to database but didn't get one")
			} else if !tc.wantError && err != nil {
				t.Errorf("error connecting to database, got err: %v", err)
			}
		})
	}
}

func TestBaseDSN(t *testing.T) {
	testCases := map[string]string{
		gin.DebugMode: "upcourse",
		gin.TestMode:  "upcourse_test",
	}

	for mode, dbName := range testCases {
		t.Run(mode, func(t *testing.T) {
			dsn := baseDSN(mode)

			expected := fmt.Sprintf("host=localhost port=5432 user=postgres dbname=%s sslmode=disable", dbName)

			if dsn != expected {
				t.Errorf("got %s want %s for dsn", dsn, expected)
			}
		})
	}
}

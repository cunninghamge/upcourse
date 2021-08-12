package config

import (
	"errors"
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

			fns := map[string]func() error{
				"Connect": Connect,
				"pingDB":  pingDB,
			}

			for fnName, fn := range fns {
				err := fn()
				if tc.wantError && err == nil {
					t.Errorf("expected an error executing %s but didn't get one", fnName)
				} else if !tc.wantError && err != nil {
					t.Errorf("error connecting to database, got err: %v", err)
				}
			}
		})
	}
}

func pingDB() error {
	if Conn == nil {
		return errors.New("database pointer reference is nil")
	}
	sqlDB, err := Conn.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Ping()
	return err
}

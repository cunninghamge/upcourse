package config

import (
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TestConnect(t *testing.T) {
	testCases := []struct {
		name      string
		mode      string
		wantError bool
	}{
		{
			name:      "release mode",
			mode:      gin.ReleaseMode,
			wantError: true,
		},
		{
			name:      "test mode",
			mode:      gin.TestMode,
			wantError: false,
		},
		{
			name:      "debug mode",
			mode:      gin.DebugMode,
			wantError: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(tt.mode)
			err := Connect()
			if tt.wantError && err == nil {
				t.Error("expected an error but didn't get one")
			} else if !tt.wantError && err != nil {
				t.Errorf("error connecting to database, got err: %v", err)
			}

			err = pingDB(Conn)
			if tt.wantError && err == nil {
				t.Error("expected an error but didn't get one")
			} else if !tt.wantError && err != nil {
				t.Errorf("database is not connected, got err: %v", err)
			}
		})
	}

	gin.SetMode(gin.TestMode)
	Connect() //nolint
}

func TestDBConnectRelease(t *testing.T) {
	db, err := DBConnectRelease()
	if err == nil {
		t.Error("expected an error but didn't get one")
	}

	err = pingDB(db)
	if err == nil {
		t.Errorf("expected an error but didn't get one")
	}
}

func TestDBConnect(t *testing.T) {
	testCases := []struct {
		name   string
		dbName string
	}{
		{
			name:   "connect to test database",
			dbName: "upcourse_test",
		}, {
			name:   "connect to dev database",
			dbName: "upcourse",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			db, err := DBConnect(tt.dbName)
			if err != nil {
				t.Errorf("error connecting to test database, got err: %v", err)
			}

			err = pingDB(db)
			if err != nil {
				t.Errorf("database is not connected, got err: %v", err)
			}
		})
	}
}

func pingDB(db *gorm.DB) error {
	sqlDB, _ := db.DB()
	return sqlDB.Ping()
}

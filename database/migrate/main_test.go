package main

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
	db "upcourse/database"
	"upcourse/models"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	if err := db.Connect(); err != nil {
		log.Fatal("could not connect to test database")
	}

	dropTables()

	code := m.Run()

	db.Conn.Where("1=1").Delete(&models.Course{})

	os.Exit(code)
}

func dropTables() {
	db.Conn.Migrator().DropTable(&models.ModuleActivity{})
	db.Conn.Migrator().DropTable(&models.Module{})
	db.Conn.Migrator().DropTable(&models.Activity{})
	db.Conn.Migrator().DropTable(&models.Course{})
}
func TestPackageMain(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		defer log.SetOutput(os.Stderr)

		main()

		messages := []string{
			"Completed automigration of database models",
			"Migration complete",
		}
		for _, message := range messages {
			if !strings.Contains(buf.String(), message) {
				t.Errorf("expected but did not receive output message: %s", message)
			}
		}
	})

	t.Run("configuration is specific to mode", func(t *testing.T) {
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

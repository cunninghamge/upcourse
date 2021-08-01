package handlers

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"

	"upcourse/config"
	"upcourse/models"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	config.Connect()
	code := m.Run()
	os.Exit(code)
}

func teardown() {
	config.Conn.Unscoped().Where("1=1").Delete(&models.ModuleActivity{})
	config.Conn.Unscoped().Where("1=1").Delete(&models.Module{})
	config.Conn.Unscoped().Where("1=1").Delete(&models.Course{})
	config.Conn.Unscoped().Where("custom=true").Delete(&models.Activity{})
}

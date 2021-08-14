package helpers

import (
	"errors"
	"upcourse/config"
	"upcourse/models"
)

const DatabaseErr = "some database error"

func Teardown() {
	config.Conn.Where("1=1").Delete(&models.ModuleActivity{})
	config.Conn.Where("1=1").Delete(&models.Module{})
	config.Conn.Where("1=1").Delete(&models.Course{})
	config.Conn.Where("custom=true").Delete(&models.Activity{})
}

func ForceError() {
	config.Conn.Error = errors.New(DatabaseErr)
}

func ClearError() {
	config.Conn.Error = nil
}

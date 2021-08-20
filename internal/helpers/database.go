package helpers

import (
	"errors"
	db "upcourse/database"
	"upcourse/models"
)

const DatabaseErr = "some database error"

func Teardown() {
	db.Conn.Where("1=1").Delete(&models.ModuleActivity{})
	db.Conn.Where("1=1").Delete(&models.Module{})
	db.Conn.Where("1=1").Delete(&models.Course{})
	db.Conn.Where("custom=true").Delete(&models.Activity{})
}

func ForceError() {
	db.Conn.Error = errors.New(DatabaseErr)
}

func ClearError() {
	db.Conn.Error = nil
}

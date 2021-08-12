package mocks

import (
	"upcourse/config"
	"upcourse/models"
)

func DefaultActivities() []models.Activity {
	var activities []models.Activity
	config.Conn.Find(&activities).Where("custom = false")
	return activities
}

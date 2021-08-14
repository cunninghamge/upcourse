package mocks

import (
	"upcourse/config"
	"upcourse/models"
)

func DefaultActivities() []*models.Activity {
	var activities []*models.Activity
	config.Conn.Select("id, name, description, metric, multiplier").Where("custom = false").Find(&activities)
	return activities
}

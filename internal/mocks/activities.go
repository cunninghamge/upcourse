package mocks

import (
	db "upcourse/database"
	"upcourse/models"
)

func DefaultActivities() []*models.Activity {
	var activities []*models.Activity
	db.Conn.Select("id, name, description, metric, multiplier").Where("custom = false").Find(&activities)
	return activities
}

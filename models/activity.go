package models

import (
	"time"

	db "upcourse/database"
)

type Activity struct {
	ID                   int     `jsonapi:"primary,activity"`
	Name                 string  `jsonapi:"attr,name"`
	Description          string  `jsonapi:"attr,description"`
	Metric               string  `jsonapi:"attr,metric"`
	Multiplier           float32 `jsonapi:"attr,multiplier"`
	Custom               bool
	CreatedAt, UpdatedAt time.Time `gorm:"type:timestamp with time zone;default:CURRENT_TIMESTAMP"`
}

func GetActivities() ([]*Activity, error) {
	var activities []*Activity
	tx := db.Conn.Select("id, name, description, metric, multiplier").Where("custom = FALSE").Find(&activities)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return activities, nil
}

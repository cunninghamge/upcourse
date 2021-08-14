package models

import (
	"time"
)

type Activity struct {
	ID                   int     `jsonapi:"primary,activity"`
	Name                 string  `jsonapi:"attr,name"`
	Description          string  `jsonapi:"attr,description"`
	Metric               string  `jsonapi:"attr,metric"`
	Multiplier           float32 `jsonapi:"attr,multiplier"`
	Custom               bool
	CreatedAt, UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

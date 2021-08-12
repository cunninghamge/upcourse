package models

import (
	"time"
)

type Activity struct {
	ID                   int       `json:"-"`
	Name                 string    `json:"name"`
	Description          string    `json:"description"`
	Metric               string    `json:"metric"`
	Multiplier           float32   `json:"multiplier"`
	Custom               bool      `json:"-"`
	CreatedAt, UpdatedAt time.Time `json:"-" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

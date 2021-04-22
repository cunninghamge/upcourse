package models

import "time"

type ModuleActivity struct {
	ID                   int       `json:"id"`
	Input                int       `json:"input"`
	Notes                string    `json:"notes"`
	ModuleId             int       `json:"-"`
	ActivityId           int       `json:"activityId"`
	CreatedAt, UpdatedAt time.Time `json:"-" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	Activity             Activity  `json:"activity"`
}

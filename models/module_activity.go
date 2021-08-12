package models

import (
	"time"
)

type ModuleActivity struct {
	ID                   int       `json:"-"`
	Input                int       `json:"input"`
	Notes                string    `json:"notes"`
	ModuleId             int       `json:"-"`
	Module               Module    `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ActivityId           int       `json:"activityId"`
	Activity             Activity  `json:"-"`
	CreatedAt, UpdatedAt time.Time `json:"-" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

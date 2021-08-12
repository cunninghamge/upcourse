package models

import (
	"time"
)

// TODO: figure out how to set moduleActivities to -
type Module struct {
	ID                   int              `json:"-"`
	Name                 string           `json:"name" validate:"onCreate"`
	Number               int              `json:"number,omitempty" validate:"onCreate"`
	CourseId             int              `json:"-"`
	Course               Course           `json:"-" gorm:"constraint:OnDelete:CASCADE"`
	ModuleActivities     []ModuleActivity `json:"moduleActivities"`
	CreatedAt, UpdatedAt time.Time        `json:"-" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

package models

import (
	"time"
)

type Module struct {
	ID                   int              `json:"id"`
	Name                 string           `json:"name" validate:"onCreate"`
	Number               int              `json:"number,omitempty" validate:"onCreate"`
	CourseId             int              `json:"courseId"`
	Course               Course           `json:"-" gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt, UpdatedAt time.Time        `json:"-" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	ModuleActivities     []ModuleActivity `json:"moduleActivities,omitempty"`
}

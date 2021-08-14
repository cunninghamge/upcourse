package models

import (
	"time"
)

type Module struct {
	ID                   int    `jsonapi:"primary,module"`
	Name                 string `jsonapi:"attr,name,omitempty" validate:"onCreate"`
	Number               int    `jsonapi:"attr,number,omitempty" validate:"onCreate"`
	CourseId             int
	CreatedAt, UpdatedAt time.Time         `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	ModuleActivities     []*ModuleActivity `jsonapi:"relation,module_activities,omitempty"`
}

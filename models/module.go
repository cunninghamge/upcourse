package models

import (
	"time"
)

type Module struct {
	ID                   int    `jsonapi:"primary,module"`
	Name                 string `jsonapi:"attr,name,omitempty"`
	Number               int    `jsonapi:"attr,number" gorm:"not null"`
	CourseId             int
	CreatedAt, UpdatedAt time.Time         `gorm:"type:timestamp with time zone;default:CURRENT_TIMESTAMP"`
	ModuleActivities     []*ModuleActivity `jsonapi:"relation,module_activities,omitempty"`
}

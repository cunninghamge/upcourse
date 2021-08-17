package models

import (
	"time"
)

type ModuleActivity struct {
	ID                   int    `jsonapi:"primary,module_activity"`
	Input                int    `jsonapi:"attr,input" gorm:"not null"`
	Notes                string `jsonapi:"attr,notes"`
	ModuleId             int
	ActivityId           int       `gorm:"not null"`
	CreatedAt, UpdatedAt time.Time `gorm:"type:timestamp with time zone;default:CURRENT_TIMESTAMP"`
	Activity             *Activity `jsonapi:"relation,activity"`
}

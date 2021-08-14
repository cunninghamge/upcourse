package models

import (
	"time"
)

type ModuleActivity struct {
	ID                   int    `jsonapi:"primary,module_activity"`
	Input                int    `jsonapi:"attr,input"`
	Notes                string `jsonapi:"attr,notes"`
	ModuleId             int
	ActivityId           int
	CreatedAt, UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	Activity             *Activity `jsonapi:"relation,activity"`
}

package models

import (
	"time"
)

type Course struct {
	ID                   int       `jsonapi:"primary,course"`
	Name                 string    `jsonapi:"attr,name" validate:"onCreate"`
	Institution          string    `jsonapi:"attr,institution,omitempty" validate:"onCreate"`
	CreditHours          int       `jsonapi:"attr,credit_hours,omitempty" validate:"onCreate"`
	Length               int       `jsonapi:"attr,length,omitempty" validate:"onCreate"`
	Goal                 string    `jsonapi:"attr,goal,omitempty"`
	CreatedAt, UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	Modules              []*Module `jsonapi:"relation,modules"`
}

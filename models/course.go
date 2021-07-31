package models

import (
	"time"

	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	ID                   int       `json:"id"`
	Name                 string    `json:"name" validate:"onCreate"`
	Institution          string    `json:"institution,omitempty" validate:"onCreate"`
	CreditHours          int       `json:"creditHours,omitempty" validate:"onCreate"`
	Length               int       `json:"length,omitempty" validate:"onCreate"`
	Goal                 string    `json:"goal,omitempty"`
	CreatedAt, UpdatedAt time.Time `json:"-" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	Modules              []Module  `json:"modules"`
}

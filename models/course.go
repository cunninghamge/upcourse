package models

import (
	"time"
)

type Course struct {
	ID                   int       `jsonapi:"primary,course"`
	Name                 string    `jsonapi:"attr,name" gorm:"not null"`
	Institution          string    `jsonapi:"attr,institution" gorm:"not null"`
	CreditHours          int       `jsonapi:"attr,credit_hours" gorm:"not null"`
	Length               int       `jsonapi:"attr,length" gorm:"not null"`
	Goal                 string    `jsonapi:"attr,goal" gorm:"not null"`
	CreatedAt, UpdatedAt time.Time `gorm:"type:timestamp with time zone;default:CURRENT_TIMESTAMP"`
	Modules              []*Module `jsonapi:"relation,modules"`
}

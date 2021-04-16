package models

import "time"

type Course struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Institution string    `json:"institution,omitempty" binding:"required"`
	CreditHours int       `json:"creditHours,omitempty" binding:"required"`
	Length      int       `json:"length,omitempty" binding:"required"`
	Goal        string    `json:"goal,omitempty"`
	CreatedAt   time.Time `json:"-" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"-" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	Modules     []Module  `json:"modules"`
}

type UpdatableCourse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Institution string `json:"institution"`
	CreditHours int    `json:"creditHours"`
	Length      int    `json:"length"`
	Goal        string `json:"goal"`
}

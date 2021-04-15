package models

import "time"

type Course struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Institution string    `json:"institution" binding:"required"`
	CreditHours int       `json:"creditHours" binding:"required"`
	Length      int       `json:"length" binding:"required"`
	Goal        string    `json:"goal"`
	CreatedAt   time.Time `json:"-" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"-" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	Modules     []Module  `json:"modules"`
}

type CourseIdentifier struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

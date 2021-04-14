package models

import "time"

type Course struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Institution string    `json:"institution"`
	CreditHours int       `json:"creditHours"`
	Length      int       `json:"length"`
	Goal        string    `json:"goal"`
	CreatedAt   time.Time `json:"-" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"-" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	Modules     []Module  `json:"modules"`
}

type CourseIdentifier struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

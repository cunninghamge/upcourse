package models

import "time"

type ModuleActivity struct {
	Id         int       `json:"id"`
	Input      int       `json:"input"`
	Notes      string    `json:"notes"`
	ModuleId   int       `json:"-"`
	ActivityId int       `json:"-"`
	CreatedAt  time.Time `json:"-" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `json:"-" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	Activity   Activity  `json:"activity"`
}

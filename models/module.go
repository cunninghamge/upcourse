package models

import "time"

type Module struct {
	ID               int              `json:"id"`
	Name             string           `json:"name"`
	Number           int              `json:"number"`
	CourseId         int              `json:"courseId"`
	CreatedAt        time.Time        `json:"-" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt        time.Time        `json:"-" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	ModuleActivities []ModuleActivity `json:"moduleActivities"`
}

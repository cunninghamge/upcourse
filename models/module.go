package models

import "time"

type Module struct {
	ID                   int              `json:"id"`
	Name                 string           `json:"name" binding:"required"`
	Number               int              `json:"number,omitempty" binding:"required"`
	CourseId             int              `json:"courseId"`
	CreatedAt, UpdatedAt time.Time        `json:"-" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	ModuleActivities     []ModuleActivity `json:"moduleActivities,omitempty"`
}

type UpdatableModule struct {
	ID               int              `json:"id"`
	Name             string           `json:"name"`
	Number           int              `json:"number"`
	ModuleActivities []ModuleActivity `json:"moduleActivities"`
}

package models

import "time"

type Activity struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Metric      string    `json:"metric"`
	Multiplier  float32   `json:"multiplier"`
	Custom      bool      `json:"-"`
	CreatedAt   time.Time `json:"-" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"-" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

type ActivityTotals struct {
	Id         int     `json:"activityId"`
	Name       string  `json:"activityName"`
	ModuleId   int     `json:"moduleId"`
	ModuleName string  `json:"moduleName"`
	Minutes    float32 `json:"minutes"`
}

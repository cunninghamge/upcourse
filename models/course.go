package models

import (
	"time"
	db "upcourse/database"

	"gorm.io/gorm"
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

func GetCourse(id string) (*Course, error) {
	var course Course
	tx := db.Conn.Preload("Modules.ModuleActivities.Activity").
		First(&course, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &course, nil
}

func GetCourseList() ([]*Course, error) {
	var courses []*Course
	tx := db.Conn.Preload("Modules", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name, number, course_id")
	}).Select("courses.id, courses.name").Find(&courses)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return courses, nil
}

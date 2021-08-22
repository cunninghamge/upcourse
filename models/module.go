package models

import (
	"strconv"
	"time"
	db "upcourse/database"
)

type Module struct {
	ID                   int    `jsonapi:"primary,module"`
	Name                 string `jsonapi:"attr,name,omitempty"`
	Number               int    `jsonapi:"attr,number" gorm:"not null"`
	CourseId             int
	CreatedAt, UpdatedAt time.Time         `gorm:"type:timestamp with time zone;default:CURRENT_TIMESTAMP"`
	ModuleActivities     []*ModuleActivity `jsonapi:"relation,module_activities,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func GetModule(id string) (*Module, error) {
	var module Module
	tx := db.Conn.Preload("ModuleActivities.Activity").
		First(&module, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &module, nil
}

func CreateModule(module *Module) error {
	return db.Conn.Create(module).Error
}

func UpdateModule(module *Module, id string) error {
	tx := db.Conn.First(&Module{}, id).Updates(&module)
	if tx.Error != nil {
		return tx.Error
	}

	moduleId, _ := strconv.Atoi(id)
	for _, ma := range module.ModuleActivities {
		tx = db.Conn.Where(ModuleActivity{ModuleId: moduleId, ActivityId: ma.ActivityId}).
			FirstOrCreate(&ModuleActivity{}).
			Updates(&ma)
		if tx.Error != nil {
			return tx.Error
		}
	}

	return nil
}

func DeleteModule(id string) error {
	return db.Conn.Delete(&Module{}, id).Error
}

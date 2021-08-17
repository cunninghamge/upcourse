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
	ModuleActivities     []*ModuleActivity `jsonapi:"relation,module_activities,omitempty"`
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
<<<<<<< HEAD
	return db.Conn.Create(module).Error
}

func UpdateModule(module *Module, id string) error {
	tx := db.Conn.First(&Module{}, id).Updates(&module)
=======
	tx := db.Conn.Create(module)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func UpdateModule(module *Module, moduleId string) error {
	tx := db.Conn.First(&Module{}, moduleId).Updates(&module)
>>>>>>> 898d7bb (move module database logic to models package)
	if tx.Error != nil {
		return tx.Error
	}

<<<<<<< HEAD
	moduleId, _ := strconv.Atoi(id)
	for _, ma := range module.ModuleActivities {
		tx = db.Conn.Where(ModuleActivity{ModuleId: moduleId, ActivityId: ma.ActivityId}).
=======
	id, _ := strconv.Atoi(moduleId)
	for _, ma := range module.ModuleActivities {
		tx = db.Conn.Where(ModuleActivity{ModuleId: id, ActivityId: ma.ActivityId}).
>>>>>>> 898d7bb (move module database logic to models package)
			FirstOrCreate(&ModuleActivity{}).
			Updates(&ma)
		if tx.Error != nil {
			return tx.Error
		}
	}

	return nil
}

<<<<<<< HEAD
func DeleteModule(id string) error {
	return db.Conn.Delete(&Module{}, id).Error
=======
func DeleteModule(moduleId string) error {
	tx := db.Conn.Delete(&Module{}, moduleId)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
>>>>>>> 898d7bb (move module database logic to models package)
}

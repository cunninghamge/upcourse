package models

import (
	"fmt"
	"time"

	db "upcourse/config"
)

// TODO: move CourseID to path params instead of body in update module request
// TODO: figure out how to set moduleActivities to -
type Module struct {
	ID                   int              `json:"-"`
	Name                 string           `json:"name" validate:"onCreate"`
	Number               int              `json:"number,omitempty" validate:"onCreate"`
	CourseId             int              `json:"courseId"`
	Course               Course           `json:"-" gorm:"constraint:OnDelete:CASCADE"`
	ModuleActivities     []ModuleActivity `json:"moduleActivities"`
	CreatedAt, UpdatedAt time.Time        `json:"-" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

func (m Module) UpdateModuleActivities() error {
	var existingActivityIds []int
	if err := db.Conn.Model(&ModuleActivity{}).Where("module_id = ?", m.ID).Select("activity_id").Scan(&existingActivityIds).Error; err != nil {
		return err
	}
	fmt.Printf("len(m.ModuleActivities): %v\n", len(m.ModuleActivities))
	fmt.Printf("existingActivityIds: %v\n", existingActivityIds)
	for _, modActivity := range m.ModuleActivities {
		modActivity.ModuleId = m.ID
		activityExists := func() bool {
			for _, activityId := range existingActivityIds {
				if modActivity.ActivityId == activityId {
					return true
				}
			}
			return false
		}()

		if activityExists {
			if err := db.Conn.Model(&ModuleActivity{}).
				Where("module_id = ? AND activity_id = ?", modActivity.ModuleId, modActivity.ActivityId).
				Updates(&modActivity).Error; err != nil {
				return err
			}
		} else {
			if err := db.Conn.Model(&ModuleActivity{}).Create(&modActivity).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

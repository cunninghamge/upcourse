package mocks

import (
	"github.com/Pallinder/go-randomdata"

	"upcourse/config"
	"upcourse/models"
)

func Module() *models.Module {
	course := SimpleCourse()

	module := models.Module{
		Name:     "Test Module",
		Number:   1,
		CourseId: course.ID,
	}
	config.Conn.Create(&module)

	var moduleActivities []models.ModuleActivity
	for i := 0; i < 4; i++ {
		moduleActivities = append(moduleActivities, models.ModuleActivity{
			Input:      randomdata.Number(200),
			Notes:      "notes",
			ActivityId: i + 1,
			ModuleId:   module.ID,
		})
	}
	config.Conn.Create(&moduleActivities)

	config.Conn.Preload("ModuleActivities.Activity").First(&module)
	return &module
}

func SimpleModule() *models.Module {
	course := SimpleCourse()

	module := models.Module{
		Name:     "Test Module",
		Number:   1,
		CourseId: course.ID,
	}
	config.Conn.Create(&module)

	return &module
}

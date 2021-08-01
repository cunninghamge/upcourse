package mocks

import (
	"strconv"

	"github.com/Pallinder/go-randomdata"

	"upcourse/config"
	"upcourse/models"
)

func NewFullCourse() *models.Course {
	course := models.Course{
		Name:        "Test Course",
		Institution: randomdata.LastName() + " University",
		CreditHours: randomdata.Number(5),
		Length:      randomdata.Number(16),
	}
	config.Conn.Create(&course)

	var modules []models.Module
	for i := 1; i < 5; i++ {
		modules = append(modules, models.Module{
			Name:     "Module " + strconv.Itoa(i),
			Number:   i,
			CourseId: course.ID,
		})
	}
	config.Conn.Create(&modules)

	var moduleActivities []models.ModuleActivity
	for _, module := range modules {
		for i := 1; i < 5; i++ {
			moduleActivities = append(moduleActivities, models.ModuleActivity{
				Input:      randomdata.Number(200),
				Notes:      "notes",
				ActivityId: randomdata.Number(13) + 1,
				ModuleId:   module.ID,
			})
		}
	}
	config.Conn.Create(&moduleActivities)

	config.Conn.Preload("Modules.ModuleActivities.Activity").First(&course)
	return &course
}
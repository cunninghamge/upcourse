package mocks

import (
	"strconv"
	"time"

	"github.com/Pallinder/go-randomdata"

	db "upcourse/database"
	"upcourse/models"
)

func FullCourse() *models.Course {
	course := models.Course{
		Name:        "Test Course",
		Institution: randomdata.LastName() + " University",
		CreditHours: randomdata.Number(5),
		Length:      randomdata.Number(16),
	}
	db.Conn.Create(&course)

	var modules []models.Module
	for i := 1; i < 5; i++ {
		modules = append(modules, models.Module{
			Name:     "Module " + strconv.Itoa(i),
			Number:   i,
			CourseId: course.ID,
		})
	}
	db.Conn.Create(&modules)

	var moduleActivities []models.ModuleActivity
	for _, module := range modules {
		for i := 1; i < 5; i++ {
			moduleActivities = append(moduleActivities, models.ModuleActivity{
				Input:      randomdata.Number(200),
				Notes:      "notes",
				ActivityId: i,
				ModuleId:   module.ID,
			})
		}
	}
	db.Conn.Create(&moduleActivities)

	db.Conn.Preload("Modules.ModuleActivities.Activity").First(&course)

	course.CreatedAt = time.Time{}
	course.UpdatedAt = time.Time{}
	for _, m := range course.Modules {
		m.CreatedAt = time.Time{}
		m.UpdatedAt = time.Time{}
		for _, ma := range m.ModuleActivities {
			ma.CreatedAt = time.Time{}
			ma.UpdatedAt = time.Time{}
			ma.Activity.CreatedAt = time.Time{}
			ma.Activity.UpdatedAt = time.Time{}
		}
	}

	return &course
}

func SimpleCourse() *models.Course {
	course := models.Course{ID: 1, Name: "Foundations of Nursing"}
	db.Conn.Create(&course)
	return &course
}

func CourseList() []*models.Course {
	var courseList []*models.Course
	for i := 0; i < 3; i++ {
		courseList = append(courseList, &models.Course{
			Name: "Test Course " + strconv.Itoa(i),
			Modules: []*models.Module{
				{
					Name: "Module 1",
				},
				{
					Name: "Module 2",
				},
				{
					Name: "Module 3",
				},
			},
		})
	}
	db.Conn.Create(&courseList)

	return courseList
}

package main

import (
	"reflect"
	"regexp"
	"strings"
	"testing"
	"time"
	db "upcourse/database"
	"upcourse/models"
)

const preloaderStr = "Modules.ModuleActivities.Activity"

var tables = map[string]interface{}{
	"courses":           models.Course{},
	"modules":           models.Module{},
	"module_activities": models.ModuleActivity{},
	"activities":        models.Activity{},
}

func TestAutoMigration(t *testing.T) {
	for tableName, model := range tables {
		t.Run(tableName+" table exists", func(t *testing.T) {
			if !db.Conn.Migrator().HasTable(tableName) {
				t.Errorf("database is missing table %s", tableName)
			}
		})

		t.Run(tableName+" has expected columns", func(t *testing.T) {
			v := reflect.ValueOf(model)
			for i := 0; i < v.NumField(); i++ {
				colName := v.Type().Field(i).Name
				if strings.Contains(preloaderStr, colName) || colName == "Model" {
					continue
				}

				if !db.Conn.Migrator().HasColumn(model, colName) {
					t.Errorf("table %s missing column %s", tableName, colName)
				}
			}
		})

		t.Run(tableName+" has correct nullable columns", func(t *testing.T) {
			v := reflect.ValueOf(model)
			for i := 0; i < v.NumField(); i++ {
				colName := toSnakeCase(v.Type().Field(i).Name)
				tag := v.Type().Field(i).Tag.Get("gorm")
				wantNotNull := strings.Contains(tag, "not null") || colName == "id"
				var nullable string
				db.Conn.Raw(`SELECT is_nullable FROM information_schema.columns WHERE table_name=? AND column_name=?`, tableName, colName).Scan(&nullable)
				if wantNotNull && nullable == "YES" {
					t.Errorf("missing NOT NULL constraint on column %s of table %s", colName, tableName)
				} else if !wantNotNull && nullable == "NO" {
					t.Errorf("unexpected NOT NULL constraint on column %s of table %s", colName, tableName)
				}
			}
		})
	}
}

func TestIndexes(t *testing.T) {
	indexes := map[string]interface{}{
		"index_module_activities_on_activities_modules": models.ModuleActivity{},
		"index_modules_on_courses_number":               models.Module{},
	}

	for name, model := range indexes {
		t.Run(name+" exists", func(t *testing.T) {
			if !db.Conn.Migrator().HasIndex(model, name) {
				t.Errorf("missing index %s", name)
			}
		})
	}

	t.Run("index_module_activities_on_activities_modules works as expected", func(t *testing.T) {
		module := mockCourse().Modules[0]
		defer db.Conn.Where("1=1").Delete(&models.Course{})

		invalidInsertion := []models.ModuleActivity{
			{Input: 15, ModuleId: module.ID, ActivityId: 1},
			{Input: 10, ModuleId: module.ID, ActivityId: 1},
		}

		err := db.Conn.Create(&invalidInsertion).Error
		if err == nil {
			t.Errorf("expected an error inserting invalid moduleActivities but didn't get one")
		}
		expectedErr := "duplicate key value violates unique constraint \"index_module_activities_on_activities_modules\""
		if !strings.Contains(err.Error(), expectedErr) {
			t.Errorf("got %v want %s for error message", err, expectedErr)
		}
	})

	t.Run("index_modules_on_courses_number works as expected", func(t *testing.T) {
		courseId := mockCourse().ID
		defer db.Conn.Where("1=1").Delete(&models.Course{})

		invalidInsertion := models.Module{CourseId: courseId, Number: 1}

		err := db.Conn.Create(&invalidInsertion).Error
		if err == nil {
			t.Errorf("expected an error inserting invalid module but didn't get one")
		}
		expectedErr := "duplicate key value violates unique constraint \"index_modules_on_courses_number\""
		if !strings.Contains(err.Error(), expectedErr) {
			t.Errorf("got %v want %s for error message", err, expectedErr)
		}
	})
}

func TestForeignKeyConstraints(t *testing.T) {
	constraints := []string{
		"fk_modules_module_activities",
		"fk_activities_module_activities",
		"fk_courses_modules",
	}

	for _, name := range constraints {
		t.Run(name+" ON DELETE", func(t *testing.T) {
			var actionType string
			db.Conn.Raw(`SELECT confdeltype FROM pg_constraint WHERE conname = ?`, name).Scan(&actionType)
			if actionType != "c" {
				t.Errorf("ON DELETE not set to CASCADE for constraint %s", name)
			}
		})
	}

	t.Run("fk_modules_module_activities ON UPDATE", func(t *testing.T) {
		var actionType string
		db.Conn.Raw(`SELECT confupdtype FROM pg_constraint WHERE conname = 'fk_modules_module_activities'`).Scan(&actionType)
		if actionType != "c" {
			t.Errorf("ON UPDATE not set to CASCADE for constraint fk_modules_module_activities")
		}
	})
}

func TestTimestamps(t *testing.T) {
	course := mockCourse()
	defer db.Conn.Where("1=1").Delete(&models.Course{})

	nilTime := time.Time{}
	if course.CreatedAt == nilTime {
		t.Errorf("CreatedAt is empty but should not have been")
	}
	if course.UpdatedAt == nilTime {
		t.Errorf("UpdatedAt is empty but should not have been")
	}

	time.Sleep(5)
	db.Conn.Model(&course).Update("credit_hours", 4)
	if course.CreatedAt == course.UpdatedAt {
		t.Errorf("UpdatedAt did not change with record update")
	}
}

func mockCourse() *models.Course {
	var course models.Course
	db.Conn.Create(&models.Course{
		Name:        "Mock Course",
		Institution: "Mock University",
		CreditHours: 3,
		Length:      8,
		Goal:        "8-10 hours",
	}).First(&course)
	db.Conn.Create(&models.Module{
		Number:   1,
		CourseId: course.ID,
	})
	db.Conn.Preload(preloaderStr).First(&course)

	return &course
}

func toSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

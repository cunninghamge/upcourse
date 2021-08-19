package main

import (
	"reflect"
	"regexp"
	"strings"
	"testing"
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
				if !strings.Contains(preloaderStr, colName) && !db.Conn.Migrator().HasColumn(model, colName) {
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
	t.Run("index exists", func(t *testing.T) {
		idxName := "index_module_activities_on_activities_modules"
		if !db.Conn.Migrator().HasIndex(&models.ModuleActivity{}, idxName) {
			t.Errorf("missing index %s", idxName)
		}
	})

	t.Run("index works as expected", func(t *testing.T) {
		module := mockCourse().Modules[0]
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
}

func TestTriggers(t *testing.T) {
	tableNames := []string{
		"courses",
		"modules",
		"module_activities",
		"activities",
	}

	for _, name := range tableNames {
		var trigger string
		db.Conn.Raw(`SELECT action_statement FROM information_schema.triggers WHERE event_object_table=?`, name).Scan(&trigger)
		expTrigger := "EXECUTE FUNCTION update_updated_at_column()"
		if trigger != expTrigger {
			t.Errorf("expected %s got %s for update trigger on %s table", expTrigger, trigger, name)
		}
	}
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

func mockCourse() models.Course {
	var course models.Course
	db.Conn.Create(&models.Course{
		Name:        "Mock Course",
		Institution: "Mock University",
		CreditHours: 3,
		Length:      8,
		Goal:        "8-10 hours",
		Modules:     []*models.Module{{Number: 1}},
	}).Preload(preloaderStr).First(&course)

	return course
}

func toSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

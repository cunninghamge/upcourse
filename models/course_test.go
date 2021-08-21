package models

import (
	"reflect"
	"strconv"
	"testing"
	db "upcourse/database"

	"github.com/Pallinder/go-randomdata"
)

func TestGetCourse(t *testing.T) {
	mockCourse := mockFullCourse()
	defer teardown()

	t.Run("finds a course by id", func(t *testing.T) {
		id := strconv.Itoa(mockCourse.ID)
		course, err := GetCourse(id)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		foundCourseValues := reflect.ValueOf(*course)
		mockCourseValues := reflect.ValueOf(*mockCourse)
		fieldNames := []string{
			"ID",
			"Name",
			"Institution",
			"CreditHours",
			"Length",
			"Goal",
		}
		for _, field := range fieldNames {
			got := foundCourseValues.FieldByName(field)
			want := mockCourseValues.FieldByName(field)
			if got.String() != want.String() {
				t.Errorf("got %v want %v for course field %s", got, want, field)
			}
		}

		foundActivity := course.Modules[0].ModuleActivities[0].Activity
		if foundActivity == nil {
			t.Errorf("activity not included in response but should have been")
		}
	})

	t.Run("returns an error if course not found", func(t *testing.T) {
		id := strconv.Itoa(mockCourse.ID + 1)
		course, err := GetCourse(id)

		if course != nil {
			t.Errorf("got %v want nil for course", course)
		}
		wantErr := "record not found"
		if err.Error() != wantErr {
			t.Errorf("got %v want %s", err, wantErr)
		}
	})
}

func TestGetCourseList(t *testing.T) {
	t.Run("finds a list of courses", func(t *testing.T) {
		mockCourseList := mockCourseList()
		defer teardown()

		courses, err := GetCourseList()

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if len(courses) != len(mockCourseList) {
			t.Errorf("got %d want %d for number of results", len(courses), len(mockCourseList))
		}

		for i := 0; i < len(courses); i++ {
			foundCourseValues := reflect.ValueOf(*courses[i])
			mockCourseValues := reflect.ValueOf(*mockCourseList[i])
			expectedFields := []string{
				"ID",
				"Name",
			}
			for _, field := range expectedFields {
				got := foundCourseValues.FieldByName(field)
				want := mockCourseValues.FieldByName(field)
				if got.String() != want.String() {
					t.Errorf("got %v want %v for course field %s", got, want, field)
				}
			}

			emptyFields := []string{
				"Institution",
				"CreditHours",
				"Length",
				"Goal",
			}
			for _, field := range emptyFields {
				got := foundCourseValues.FieldByName(field)
				if got.String() != "" && got.Int() != 0 {
					t.Errorf("got %v for course field %s expected to be empty", got, field)
				}
			}

			for _, m := range courses[i].Modules[0].ModuleActivities {
				if m != nil {
					t.Error("module activities should not have been included in response")
				}
			}
		}
	})

	t.Run("does not return an error if no courses are found", func(t *testing.T) {
		courses, err := GetCourseList()

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if len(courses) != 0 {
			t.Errorf("got %d want 0 for number of results", len(courses))
		}
	})

	t.Run("returns database errors if they occur", func(t *testing.T) {
		mockBasicCourse()
		defer teardown()
		forceError()
		defer clearError()

		courses, err := GetCourseList()

		if err.Error() != testDBErr {
			t.Errorf("got %v want %s for error message", err, testDBErr)
		}

		if courses != nil {
			t.Errorf("got %v expected nil for courses", courses)
		}
	})
}

func mockBasicCourse() *Course {
	course := Course{Name: "Models Test Basic Course"}
	db.Conn.Create(&course)
	return &course
}

func mockFullCourse() *Course {
	course := Course{
		Name:        "Models Test Full Course",
		Institution: "Test University",
		CreditHours: randomdata.Number(5),
		Length:      randomdata.Number(16),
		Goal:        "8-10 hours",
	}
	db.Conn.Create(&course)

	for i := 1; i < 3; i++ {
		course.Modules = append(course.Modules, mockFullModule(course.ID, i))
	}

	db.Conn.Preload("Modules.ModuleActivities.Activity").First(&course)

	return &course
}

func mockCourseList() []*Course {
	var courses []*Course
	for i := 0; i < 2; i++ {
		courses = append(courses, mockFullCourse())
	}

	return courses
}

func mockFullModule(params ...int) *Module {
	if len(params) == 0 {
		params = []int{mockBasicCourse().ID, 1}
	}

	module := Module{
		Name:     "Models Test Full Module",
		CourseId: params[0],
		Number:   params[1],
	}
	for i := 0; i < 3; i++ {
		module.ModuleActivities = append(module.ModuleActivities, &ModuleActivity{
			Input:      randomdata.Number(200),
			ActivityId: i + 1,
		})
	}

	db.Conn.Create(&module).Preload("ModuleActivities.Activity").First(&module)

	return &module
}

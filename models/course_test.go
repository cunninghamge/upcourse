package models

import (
	"reflect"
	"strconv"
	"testing"
	db "upcourse/database"

	"github.com/Pallinder/go-randomdata"
)

func mockBasicCourse() *Course {
	course := Course{Name: "Models Test Basic Course"}
	db.Conn.Create(&course)
	courseIds = append(courseIds, course.ID)
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
	courseIds = append(courseIds, course.ID)

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

func TestCreateCourse(t *testing.T) {
	defer teardown()

	t.Run("creates a course", func(t *testing.T) {
		name := "Successful creation"
		newCourse := Course{
			Name:        name,
			Institution: "Test Institution",
			CreditHours: 3,
			Length:      8,
			Goal:        "8-10 hours",
		}

		err := CreateCourse(&newCourse)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		var course Course
		if err := db.Conn.Where("name = ?", name).First(&course).Error; err != nil {
			t.Error("new course not found in database")
		}
		courseIds = append(courseIds, course.ID)
	})

	t.Run("returns an error if course is not created", func(t *testing.T) {
		forceError()
		defer clearError()

		name := "Failed creation"
		newCourse := Course{
			Name:        name,
			Institution: "Test Institution",
			CreditHours: 3,
			Length:      8,
			Goal:        "8-10 hours",
		}

		err := CreateCourse(&newCourse)

		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}

		db.Conn.Error = nil
		if err := db.Conn.Where("name = ?", name).First(&Course{}).Error; err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})
}

func TestUpdateCourse(t *testing.T) {
	mockCourse := mockBasicCourse()
	defer teardown()

	t.Run("updates a course", func(t *testing.T) {
		name := "New Course name"
		mockCourse.Name = name

		err := UpdateCourse(mockCourse, strconv.Itoa(mockCourse.ID))

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		var foundCourse Course
		err = db.Conn.Where("name = ?", name).First(&foundCourse).Error
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if foundCourse.Name != name {
			t.Errorf("got %s want %s for updated course name", foundCourse.Name, name)
		}
	})

	t.Run("returns an error if course is not updated", func(t *testing.T) {
		number := mockCourse.Length + 1
		mockCourse.Length = number

		err := UpdateCourse(mockCourse, strconv.Itoa(mockCourse.ID+1))

		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}

		var foundCourse Course
		err = db.Conn.First(&foundCourse, mockCourse.ID).Error
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if foundCourse.Length == number {
			t.Errorf("got %d want %d for course number", foundCourse.Length, mockCourse.Length)
		}
	})
}

func TestDeleteCourse(t *testing.T) {
	t.Run("deletes a course", func(t *testing.T) {
		mockCourse := mockBasicCourse()

		err := DeleteCourse(strconv.Itoa(mockCourse.ID))

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if err = db.Conn.First(&mockCourse, mockCourse.ID).Error; err == nil {
			t.Errorf("course not deleted from database")
		}
	})

	t.Run("returns an error if course can't be deleted", func(t *testing.T) {
		mockCourse := mockBasicCourse()
		defer teardown()
		forceError()

		err := DeleteCourse(strconv.Itoa(mockCourse.ID))

		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}

		db.Conn.Error = nil
		if err = db.Conn.First(&mockCourse, mockCourse.ID).Error; err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}

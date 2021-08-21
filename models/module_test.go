package models

import (
	"reflect"
	"strconv"
	"testing"
	db "upcourse/database"

	"github.com/Pallinder/go-randomdata"
)

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

func TestGetModule(t *testing.T) {
	mockModule := mockFullModule()
	defer teardown()

	t.Run("finds a module by id", func(t *testing.T) {
		id := strconv.Itoa(mockModule.ID)
		module, err := GetModule(id)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if !reflect.DeepEqual(*module, *mockModule) {
			t.Errorf("got %v want %v for module", module, mockModule)
		}

		foundModuleValues := reflect.ValueOf(*module)
		mockModuleValues := reflect.ValueOf(*mockModule)
		fieldNames := []string{
			"ID",
			"Name",
			"Number",
		}
		for _, field := range fieldNames {
			got := foundModuleValues.FieldByName(field)
			want := mockModuleValues.FieldByName(field)
			if got.String() != want.String() {
				t.Errorf("got %v want %v for module field %s", got, want, field)
			}
		}

		foundActivity := module.ModuleActivities[0].Activity
		if foundActivity == nil {
			t.Errorf("activity not included in response but should have been")
		}
	})

	t.Run("returns an error if module not found", func(t *testing.T) {
		id := strconv.Itoa(mockModule.ID + 1)
		module, err := GetModule(id)

		if module != nil {
			t.Errorf("got %v want nil for module", module)
		}

		wantErr := "record not found"
		if err.Error() != wantErr {
			t.Errorf("got %v want %v", err, wantErr)
		}
	})
}

func TestCreateModule(t *testing.T) {
	course := mockBasicCourse()
	defer teardown()

	t.Run("creates a module", func(t *testing.T) {
		name := "Successful creation"
		newModule := Module{
			Name:     name,
			Number:   1,
			CourseId: course.ID,
			ModuleActivities: func() []*ModuleActivity {
				var moduleActivities []*ModuleActivity
				for i := 0; i < 2; i++ {
					moduleActivities = append(moduleActivities, &ModuleActivity{
						Input:      randomdata.Number(200),
						ActivityId: i + 1,
					})
				}
				return moduleActivities
			}(),
		}

		err := CreateModule(&newModule)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if err := db.Conn.Where("name = ?", name).First(&Module{}).Error; err != nil {
			t.Error("new module not found in database")
		}
	})

	t.Run("returns an error if module is not created", func(t *testing.T) {
		name := "Failed creation"
		newModule := Module{
			Name:     name,
			Number:   1,
			CourseId: course.ID + 1,
			ModuleActivities: func() []*ModuleActivity {
				var moduleActivities []*ModuleActivity
				for i := 0; i < 2; i++ {
					moduleActivities = append(moduleActivities, &ModuleActivity{
						Input:      randomdata.Number(200),
						ActivityId: i + 1,
					})
				}
				return moduleActivities
			}(),
		}

		err := CreateModule(&newModule)

		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
		if err := db.Conn.Where("name = ?", name).First(&Module{}).Error; err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})
}

func TestUpdateModule(t *testing.T) {
	mockModule := mockFullModule()
	defer teardown()

	t.Run("updates a module", func(t *testing.T) {
		name := "New Module name"
		mockModule.Name = name

		err := UpdateModule(mockModule, strconv.Itoa(mockModule.ID))

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		var foundModule Module
		err = db.Conn.Where("name = ?", name).First(&foundModule).Error
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if foundModule.Name != name {
			t.Errorf("got %s want %s for updated module name", foundModule.Name, name)
		}
	})

	t.Run("returns an error if module is not updated", func(t *testing.T) {
		number := mockModule.Number + 1
		mockModule.Number = number

		err := UpdateModule(mockModule, strconv.Itoa(mockModule.ID+1))

		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}

		var foundModule Module
		err = db.Conn.First(&foundModule, mockModule.ID).Error
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if foundModule.Number == number {
			t.Errorf("got %d want %d for module number", foundModule.Number, mockModule.Number)
		}
	})

	t.Run("returns an error if a moduleActivity can't be updated", func(t *testing.T) {
		activityId := mockModule.ModuleActivities[0].ID
		for _, ma := range mockModule.ModuleActivities {
			ma.ActivityId = activityId
		}

		err := UpdateModule(mockModule, strconv.Itoa(mockModule.ID))

		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})
}

func TestDeleteModule(t *testing.T) {
	t.Run("deletes a module", func(t *testing.T) {
		mockModule := mockFullModule()

		err := DeleteModule(strconv.Itoa(mockModule.ID))

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if err = db.Conn.First(&mockModule, mockModule.ID).Error; err == nil {
			t.Errorf("module not deleted from database")
		}
	})

	t.Run("returns an error if module can't be deleted", func(t *testing.T) {
		mockModule := mockFullModule()
		defer teardown()
		forceError()

		err := DeleteModule(strconv.Itoa(mockModule.ID))

		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}

		db.Conn.Error = nil
		if err = db.Conn.First(&mockModule, mockModule.ID).Error; err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}

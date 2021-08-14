package handlers

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	"gorm.io/gorm"

	"upcourse/config"
	testHelpers "upcourse/internal/helpers"
	"upcourse/internal/mocks"
	"upcourse/models"
)

func TestGetModule(t *testing.T) {
	mockModule := mocks.Module()
	defer testHelpers.Teardown()

	t.Run("returns a module and its moduleActivities", func(t *testing.T) {
		params := map[string]string{"id": fmt.Sprint(mockModule.ID)}
		w := testHelpers.NewRequest(params, "", GetModule)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusOK)

		response := testHelpers.UnmarshalPayload(t, w, new(models.Module))
		responseModule, ok := response.(*models.Module)
		if !ok {
			t.Errorf("error casting response element as module")
		}

		modActivitiesLen := len(responseModule.ModuleActivities)
		testHelpers.AssertEqualLength(t, modActivitiesLen, len(mockModule.ModuleActivities), "module activities")

		for i := 0; i < modActivitiesLen; i++ {
			gotModuleActivity := responseModule.ModuleActivities[i]
			wantModuleActivity := mockModule.ModuleActivities[i]

			gotActivity := gotModuleActivity.Activity
			wantActivity := wantModuleActivity.Activity
			wantActivity.CreatedAt, wantActivity.UpdatedAt = time.Time{}, time.Time{}
			if !reflect.DeepEqual(gotActivity, wantActivity) {
				t.Errorf("got %v want %v for activity for module_activity %d", gotActivity, wantActivity, wantModuleActivity.ID)
			}

			gotModuleActivity.Activity, wantModuleActivity.Activity = nil, nil
			wantModuleActivity.ModuleId, wantModuleActivity.ActivityId = 0, 0
			wantModuleActivity.CreatedAt, wantModuleActivity.UpdatedAt = time.Time{}, time.Time{}
			if !reflect.DeepEqual(gotModuleActivity, wantModuleActivity) {
				t.Errorf("got %v want %v for module_activity %d", gotModuleActivity, wantModuleActivity, wantModuleActivity.ID)
			}
		}

		responseModule.ModuleActivities, mockModule.ModuleActivities = nil, nil
		mockModule.CourseId = 0
		mockModule.CreatedAt, mockModule.UpdatedAt = time.Time{}, time.Time{}
		if !reflect.DeepEqual(responseModule, mockModule) {
			t.Errorf("got %v want %v for module", responseModule, mockModule)
		}
	})

	t.Run("returns a message if the module is not found", func(t *testing.T) {
		params := map[string]string{"id": fmt.Sprint(mockModule.ID + 1)}
		w := testHelpers.NewRequest(params, "", GetModule)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusNotFound)

		response := testHelpers.UnmarshalErrors(t, w)
		testHelpers.AssertError(t, response[0], ErrNotFound)
	})

	t.Run("returns database errors if they occur", func(t *testing.T) {
		testHelpers.ForceError()
		defer testHelpers.ClearError()

		params := map[string]string{"id": fmt.Sprint(mockModule.ID)}
		w := testHelpers.NewRequest(params, "", GetModule)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusInternalServerError)

		response := testHelpers.UnmarshalErrors(t, w)
		if response[0] != testHelpers.DatabaseErr {
			t.Errorf("got %s want %s for error message", response[0], testHelpers.DatabaseErr)
		}
	})
}

func TestCreateModule(t *testing.T) {
	course := mocks.SimpleCourse()
	defer testHelpers.Teardown()

	t.Run("it posts a new module with associated module activities", func(t *testing.T) {
		var moduleCount int64
		config.Conn.Model(models.Module{}).Count(&moduleCount)
		var moduleActivityCount int64
		config.Conn.Model(models.ModuleActivity{}).Count(&moduleActivityCount)

		newModuleInfo := `{
			"name": "Module 9",
			"number": 9,
			"moduleActivities":[
				{
					"input": 30,
					"notes": "A note",
					"activityId": 1
				},
				{
					"input": 8,
					"notes": null,
					"activityId": 2
				},
				{
					"input": 180,
					"notes": "",
					"activityId": 11
				}
			]
		}`
		params := map[string]string{"id": fmt.Sprint(course.ID)}
		w := testHelpers.NewRequest(params, newModuleInfo, CreateModule)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusCreated)

		errs := testHelpers.UnmarshalErrors(t, w)
		if len(errs) != 0 {
			t.Errorf("got unexpected errors: %s", errs)
		}

		var newModuleCount int64
		config.Conn.Model(models.Module{}).Count(&newModuleCount)
		if newModuleCount != (moduleCount + 1) {
			t.Errorf("module count did not change")
		}

		var newModuleActivityCount int64
		config.Conn.Model(models.ModuleActivity{}).Count(&newModuleActivityCount)
		if newModuleActivityCount != (moduleActivityCount + 3) {
			t.Errorf("activity count did not change")
		}
	})

	t.Run("it returns an error if a required field is missing", func(t *testing.T) {
		var moduleCount int64
		config.Conn.Model(models.Module{}).Count(&moduleCount)
		var moduleActivityCount int64
		config.Conn.Model(models.ModuleActivity{}).Count(&moduleActivityCount)

		newModuleInfo := `{
			"number": 9,
			"moduleActivities":[
				{
					"input": 30,
					"notes": "A note",
					"activityId": 1
				},
				{
					"input": 8,
					"notes": null,
					"activityId": 2
				},
				{
					"input": 180,
					"notes": "",
					"activityId": 11
				}
			]
		}`
		params := map[string]string{"id": fmt.Sprint(course.ID)}
		w := testHelpers.NewRequest(params, newModuleInfo, CreateModule)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusBadRequest)

		errs := testHelpers.UnmarshalErrors(t, w)

		testHelpers.AssertError(t, errs[0], "Name is required")

		var newModuleCount int64
		config.Conn.Model(models.Module{}).Count(&newModuleCount)
		if newModuleCount != moduleCount {
			t.Errorf("module count changed but should not have")
		}

		var newModuleActivityCount int64
		config.Conn.Model(models.ModuleActivity{}).Count(&newModuleActivityCount)
		if newModuleActivityCount != moduleActivityCount {
			t.Errorf("module activity count changed but should not have")
		}
	})

	t.Run("it returns an error if no body is sent", func(t *testing.T) {
		params := map[string]string{"id": fmt.Sprint(course.ID)}
		w := testHelpers.NewRequest(params, "", CreateModule)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusBadRequest)

		errs := testHelpers.UnmarshalErrors(t, w)
		testHelpers.AssertError(t, errs[0], ErrBadRequest)
	})

	t.Run("returns database errors if they occur", func(t *testing.T) {
		testHelpers.ForceError()
		defer testHelpers.ClearError()

		newModuleInfo := `{
			"name": "Module 9",
			"number": 9,
			"moduleActivities":[
				{
					"input": 180,
					"notes": "",
					"activityId": 11
				}
			]
		}`
		params := map[string]string{"id": fmt.Sprint(course.ID)}
		w := testHelpers.NewRequest(params, newModuleInfo, CreateModule)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusInternalServerError)

		response := testHelpers.UnmarshalErrors(t, w)
		if response[0] != testHelpers.DatabaseErr {
			t.Errorf("got %s want %s for error message", response[0], testHelpers.DatabaseErr)
		}
	})

	t.Run("it returns an error if an invalid course id is sent", func(t *testing.T) {
		newModuleInfo := `{
			"name": "Module 9",
			"number": 9,
			"moduleActivities":[
				{
					"input": 180,
					"notes": "",
					"activityId": 11
				}
			]
		}`
		params := map[string]string{"id": fmt.Sprint("-")}
		w := testHelpers.NewRequest(params, newModuleInfo, CreateModule)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusBadRequest)
	})
}

func TestUpdateModule(t *testing.T) {
	mockModule := mocks.Module()
	defer testHelpers.Teardown()

	t.Run("it updates an existing module", func(t *testing.T) {
		var moduleCount int64
		config.Conn.Model(models.Module{}).Count(&moduleCount)
		var moduleActivityCount int64
		config.Conn.Model(models.ModuleActivity{}).Count(&moduleActivityCount)

		newModuleInfo := `{
			"name": "new module name",
			"moduleActivities":[
				{
					"input": 30,
					"notes": "A new note",
					"activityId": 1
				}
			]
		}`
		params := map[string]string{"id": fmt.Sprint(mockModule.ID)}
		w := testHelpers.NewRequest(params, newModuleInfo, UpdateModule)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusOK)

		errs := testHelpers.UnmarshalErrors(t, w)
		if len(errs) != 0 {
			t.Errorf("got unexpected errors: %s", errs)
		}

		var newCount int64
		config.Conn.Model(models.Module{}).Count(&newCount)
		if newCount != moduleCount {
			t.Errorf("patch request should not have changed module count but did")
		}

		var newModActivityCount int64
		config.Conn.Model(models.ModuleActivity{}).Count(&newModActivityCount)
		if newModActivityCount != moduleActivityCount {
			t.Errorf("added a new module activity but should not have")
		}

		var updatedModule models.Module
		config.Conn.Preload("ModuleActivities", func(db *gorm.DB) *gorm.DB {
			return db.Order("id")
		}).First(&updatedModule, mockModule.ID)

		if updatedModule.Name != "new module name" {
			t.Errorf("did not update the module record")
		}

		if updatedModule.ModuleActivities[0].Notes != "A new note" {
			t.Errorf("did not update the module activity for the module")
		}

		if updatedModule.Number != mockModule.Number {
			t.Errorf("updated a field that should not have been updated")
		}
	})

	t.Run("it can add a new module activity to an existing module", func(t *testing.T) {
		var moduleCount int64
		config.Conn.Model(models.Module{}).Count(&moduleCount)
		var moduleActivityCount int64
		config.Conn.Model(models.ModuleActivity{}).Where("module_id = ?", mockModule.ID).Count(&moduleActivityCount)

		newModuleInfo := `{
			"name": "new module name",
			"moduleActivities":[
				{
					"input": 30,
					"notes": "A new note",
					"activityId": 14
				}
			]
		}`
		params := map[string]string{"id": fmt.Sprint(mockModule.ID)}
		w := testHelpers.NewRequest(params, newModuleInfo, UpdateModule)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusOK)

		errs := testHelpers.UnmarshalErrors(t, w)
		if len(errs) != 0 {
			t.Errorf("got unexpected errors: %s", errs)
		}

		var newCount int64
		config.Conn.Model(models.Course{}).Count(&newCount)
		if newCount != moduleCount {
			t.Errorf("patch request should not have changed module count but did")
		}

		var newModActivityCount int64
		config.Conn.Model(models.ModuleActivity{}).Where("module_id = ?", mockModule.ID).Count(&newModActivityCount)
		if newModActivityCount != (moduleActivityCount + 1) {
			t.Errorf("did not create a new module activity")
		}
	})

	t.Run("it returns a 404 if module not found", func(t *testing.T) {
		newModuleInfo := `{
			"name": "new module name",
			"moduleActivities":[
				{
					"id": 1,
					"input": 30,
					"notes": "A new note",
					"activityId": 1
				}
			]
		}`
		params := map[string]string{"id": fmt.Sprint(mockModule.ID + 1)}
		w := testHelpers.NewRequest(params, newModuleInfo, UpdateModule)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusNotFound)
	})

	t.Run("it returns an error if no body is sent", func(t *testing.T) {
		params := map[string]string{"id": fmt.Sprint(mockModule.ID + 1)}
		w := testHelpers.NewRequest(params, "", UpdateModule)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusBadRequest)
	})

	t.Run("returns database errors if they occur", func(t *testing.T) {
		testHelpers.ForceError()
		defer testHelpers.ClearError()

		newCourseInfo := `{
			"institution": "Tampa Bay Nurses United University",
			"creditHours": 3,
			"length": 16,
			"goal": "8-10 hours"
		}`
		params := map[string]string{"id": fmt.Sprint(mockModule.ID)}
		w := testHelpers.NewRequest(params, newCourseInfo, UpdateModule)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusInternalServerError)

		response := testHelpers.UnmarshalErrors(t, w)
		if response[0] != testHelpers.DatabaseErr {
			t.Errorf("got %s want %s for error message", response[0], testHelpers.DatabaseErr)
		}
	})
}

func TestDeleteModule(t *testing.T) {
	t.Run("Deletes a module and its children", func(t *testing.T) {
		mockModule := mocks.Module()
		defer testHelpers.Teardown()

		params := map[string]string{"id": fmt.Sprint(mockModule.ID)}
		w := testHelpers.NewRequest(params, "", DeleteModule)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusOK)

		errs := testHelpers.UnmarshalErrors(t, w)
		if len(errs) != 0 {
			t.Errorf("got unexpected errors: %s", errs)
		}

		var moduleCount int64
		config.Conn.Model(models.Module{}).Count(&moduleCount)
		if moduleCount > 0 {
			t.Errorf("Did not delete module")
		}

		var moduleActivityCount int64
		config.Conn.Model(models.ModuleActivity{}).Count(&moduleActivityCount)
		if moduleActivityCount > 0 {
			t.Errorf("Did not delete associated module activities")
		}
	})

	t.Run("returns database errors if they occur", func(t *testing.T) {
		testHelpers.ForceError()
		defer testHelpers.ClearError()

		mockModule := mocks.Module()
		defer testHelpers.Teardown()

		params := map[string]string{"id": fmt.Sprint(mockModule.ID)}
		w := testHelpers.NewRequest(params, "", DeleteCourse)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusInternalServerError)

		response := testHelpers.UnmarshalErrors(t, w)
		if response[0] != testHelpers.DatabaseErr {
			t.Errorf("got %s want %s for error message", response[0], testHelpers.DatabaseErr)
		}
	})
}

package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http/httptest"
	"reflect"
	"testing"

	"gorm.io/gorm"

	"upcourse/config"
	"upcourse/internal/mocks"
	"upcourse/models"
)

func TestGetModule(t *testing.T) {
	mockModule := mocks.Module()
	defer teardown()

	t.Run("returns a module and its moduleActivities", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{
			"id": fmt.Sprint(mockModule.ID),
		})

		GetModule(ctx)

		if w.Code != 200 {
			t.Errorf("expected response code to be 200, got %d", w.Code)
		}

		var response struct {
			Data struct {
				Type          string
				ID            int
				Attributes    models.Module
				Relationships map[string][]struct {
					Type          string
					ID            int
					Attributes    models.ModuleActivity
					Relationships map[string]struct {
						Type       string
						ID         int
						Attributes models.Activity
					}
				}
			}
		}
		if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
			t.Errorf("error unmarshaling json response: %v", err)
		}
	})

	t.Run("returns a message if the module is not found", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{
			"id": fmt.Sprint(mockModule.ID + 20),
		})

		GetModule(ctx)

		if w.Code != 404 {
			t.Errorf("expected response code to be 404, got %d", w.Code)
		}

		response := unmarshalResponse(t, w.Body)
		assertResponseValue(t, response.Errors[0], "record not found", "error message")
	})

	t.Run("returns database errors if they occur", func(t *testing.T) {
		err := "some database error"
		config.Conn.Error = errors.New(err)
		defer func() {
			config.Conn.Error = nil
		}()

		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{
			"id": fmt.Sprint(mockModule.ID),
		})

		GetModule(ctx)

		response := unmarshalResponse(t, w.Body)
		if response.Errors[0] != err {
			t.Errorf("got %s want %s for error message", response.Errors[0], err)
		}
	})
}

func TestCreateModule(t *testing.T) {
	course := mocks.SimpleCourse()
	defer teardown()

	t.Run("it posts a new module with associated module activities", func(t *testing.T) {
		var moduleCount int64
		config.Conn.Model(models.Module{}).Count(&moduleCount)
		var moduleActivityCount int64
		config.Conn.Model(models.ModuleActivity{}).Count(&moduleActivityCount)

		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{
			"id": fmt.Sprint(course.ID),
		})
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
		mocks.SetRequestBody(ctx, newModuleInfo)

		CreateModule(ctx)

		if w.Code != 201 {
			t.Errorf("expected response code to be 201, got %d", w.Code)
		}

		response := unmarshalResponse(t, w.Body)
		if len(response.Errors) != 0 {
			t.Errorf("got unexpected errors: %s", response.Errors)
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

		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{
			"id": fmt.Sprint(course.ID),
		})
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
		mocks.SetRequestBody(ctx, newModuleInfo)

		CreateModule(ctx)

		if w.Code != 400 {
			t.Errorf("expected response code to be %d, got %d", 400, w.Code)
		}

		response := unmarshalResponse(t, w.Body)

		expected := []string{"Name is required"}
		if !reflect.DeepEqual(response.Errors, expected) {
			t.Errorf("got %v, wanted %v for field Error messages", response.Errors, expected)
		}

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
		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{})

		CreateModule(ctx)

		if w.Code != 400 {
			t.Errorf("expected response code to be %d, got %d", 400, w.Code)
		}
	})

	t.Run("returns database errors if they occur", func(t *testing.T) {
		err := "some database error"
		config.Conn.Error = errors.New(err)
		defer func() {
			config.Conn.Error = nil
		}()

		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{
			"id": fmt.Sprint(course.ID),
		})
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
		mocks.SetRequestBody(ctx, newModuleInfo)

		CreateModule(ctx)

		response := unmarshalResponse(t, w.Body)
		if response.Errors[0] != err {
			t.Errorf("got %s want %s for error message", response.Errors[0], err)
		}
	})

	t.Run("it returns an error if an invalid course id is sent", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{
			"id": "-",
		})
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
		mocks.SetRequestBody(ctx, newModuleInfo)

		CreateModule(ctx)

		if w.Code != 400 {
			t.Errorf("expected response code to be %d, got %d", 400, w.Code)
		}
	})

	t.Run("it returns an error if an no course id is sent", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, nil)
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
		mocks.SetRequestBody(ctx, newModuleInfo)

		CreateModule(ctx)

		if w.Code != 400 {
			t.Errorf("expected response code to be %d, got %d", 400, w.Code)
		}
	})
}

func TestUpdateModule(t *testing.T) {
	mockModule := mocks.Module()
	defer teardown()

	t.Run("it updates an existing module", func(t *testing.T) {
		var moduleCount int64
		config.Conn.Model(models.Module{}).Count(&moduleCount)
		var moduleActivityCount int64
		config.Conn.Model(models.ModuleActivity{}).Count(&moduleActivityCount)

		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{
			"id": fmt.Sprint(mockModule.ID),
		})
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
		mocks.SetRequestBody(ctx, newModuleInfo)

		UpdateModule(ctx)

		if w.Code != 200 {
			t.Errorf("expected response code to be %d, got %d", 200, w.Code)
		}

		response := unmarshalResponse(t, w.Body)
		if len(response.Errors) != 0 {
			t.Errorf("got unexpected errors: %s", response.Errors)
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

		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{
			"id": fmt.Sprint(mockModule.ID),
		})
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
		mocks.SetRequestBody(ctx, newModuleInfo)

		UpdateModule(ctx)

		if w.Code != 200 {
			t.Errorf("expected response code to be %d, got %d", 200, w.Code)
		}

		response := unmarshalResponse(t, w.Body)
		if len(response.Errors) != 0 {
			t.Errorf("got unexpected errors: %s", response.Errors)
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
		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{
			"id": fmt.Sprint(mockModule.ID + 20),
		})
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
		mocks.SetRequestBody(ctx, newModuleInfo)

		UpdateModule(ctx)

		if w.Code != 404 {
			t.Errorf("expected response code to be %d, got %d", 404, w.Code)
		}
	})

	t.Run("it returns an error if no body is sent", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{
			"id": fmt.Sprint(mockModule.ID),
		})
		UpdateModule(ctx)

		if w.Code != 400 {
			t.Errorf("expected response code to be %d, got %d", 400, w.Code)
		}
	})
}

func TestDELETEmodule(t *testing.T) {
	mockModule := mocks.Module()

	t.Run("Deletes a module and its children", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{
			"id": fmt.Sprint(mockModule.ID),
		})

		DeleteModule(ctx)

		if w.Code != 200 {
			t.Errorf("expected response code to be %d, got %d", 200, w.Code)
		}

		response := unmarshalResponse(t, w.Body)
		if len(response.Errors) != 0 {
			t.Errorf("got unexpected errors: %s", response.Errors)
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
		err := "some database error"
		config.Conn.Error = errors.New(err)
		defer func() {
			config.Conn.Error = nil
		}()

		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{
			"id": fmt.Sprint(mockModule.ID),
		})

		DeleteModule(ctx)

		response := unmarshalResponse(t, w.Body)
		if response.Errors[0] != err {
			t.Errorf("got %s want %s for error message", response.Errors[0], err)
		}
	})
}

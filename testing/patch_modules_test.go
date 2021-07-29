package testing

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"upcourse/config"
	"upcourse/models"
	"upcourse/routes"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestPATCHModule(t *testing.T) {
	mockModule := newModule()
	defer teardown()

	t.Run("it updates a module", func(t *testing.T) {
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

		response := newPatchModuleRequest(newModuleInfo, mockModule.ID)

		assert.Equal(t, 200, response.Code)

		parsedResponse := UnmarshalSuccess(t, response.Body)

		assertResponseValue(t, parsedResponse.Message, "Module updated successfully", "Message")

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

		response := newPatchModuleRequest(newModuleInfo, mockModule.ID)

		assert.Equal(t, 200, response.Code)

		parsedResponse := UnmarshalSuccess(t, response.Body)

		assertResponseValue(t, parsedResponse.Message, "Module updated successfully", "Message")

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

	t.Run("it returns an error if database is unavailable", func(t *testing.T) {
		db, _ := config.Conn.DB()
		db.Close()

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
		response := newPatchModuleRequest(newModuleInfo, mockModule.ID)

		assert.Equal(t, 503, response.Code)

		config.Connect()
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

		response := newPatchModuleRequest(newModuleInfo, mockModule.ID+20)

		assert.Equal(t, 404, response.Code)
	})

	t.Run("it returns an error if no body is sent", func(t *testing.T) {
		response := newPatchModuleRequest("", mockModule.ID)

		assert.Equal(t, 400, response.Code)
	})
}

func newPatchModuleRequest(json string, moduleId int) *httptest.ResponseRecorder {
	router := routes.GetRoutes()
	request, _ := http.NewRequest("PATCH", fmt.Sprintf("/modules/%d", moduleId), bytes.NewBufferString(json))
	response := httptest.NewRecorder()

	request.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(response, request)

	return response
}

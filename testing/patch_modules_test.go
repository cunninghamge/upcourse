package testing

import (
	"bytes"
	"course-chart/config"
	"course-chart/models"
	"course-chart/routes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestPATCHModule(t *testing.T) {
	mockModule := newModule()
	defer teardown()

	t.Run("it updates a module", func(t *testing.T) {
		var moduleCount int64
		config.Conn.Model(models.Module{}).Count(&moduleCount)

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

		assert.Equal(t, 200, response.Code)

		parsedResponse := UnmarshalSuccess(t, response.Body)

		assertResponseValue(t, parsedResponse.Message, "Module updated successfully", "Message")

		var newCount int64
		config.Conn.Model(models.Course{}).Count(&newCount)

		if newCount != moduleCount {
			t.Errorf("patch request should not have changed module count but did")
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
}

func newPatchModuleRequest(json string, moduleId int) *httptest.ResponseRecorder {
	router := routes.GetRoutes()
	request, _ := http.NewRequest("PATCH", fmt.Sprintf("/modules/%d", moduleId), bytes.NewBufferString(json))
	response := httptest.NewRecorder()

	request.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(response, request)

	return response
}

package testing

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"upcourse/config"
	"upcourse/models"
	"upcourse/routes"

	"github.com/stretchr/testify/assert"
)

func TestDELETEmodule(t *testing.T) {
	mockModule := newModule()
	defer teardown()

	t.Run("Deletes a module and its children", func(t *testing.T) {
		var moduleCount int64
		config.Conn.Model(models.Module{}).Count(&moduleCount)

		response := newDeleteModuleRequest(mockModule.ID)

		assert.Equal(t, 200, response.Code)

		parsedResponse := UnmarshalSuccess(t, response.Body)

		assertResponseValue(t, parsedResponse.Message, "Module deleted successfully", "Message")

		var newModuleCount int64
		config.Conn.Model(models.Module{}).Count(&newModuleCount)

		if moduleCount == newModuleCount {
			t.Errorf("Did not delete module")
		}

		err := config.Conn.First(&models.Module{}, mockModule.ID).Error
		if err == nil {
			t.Errorf("Also did not delete module")
		}

		err = config.Conn.First(&models.ModuleActivity{}, mockModule.ModuleActivities[0].ID).Error
		if err == nil {
			t.Errorf("Did not delete associated module activities")
		}
	})

	t.Run("returns an error if database is unavailable", func(t *testing.T) {
		db, _ := config.Conn.DB()
		db.Close()
		response := newDeleteModuleRequest(mockModule.ID)

		assert.Equal(t, 503, response.Code)

		config.Connect()
	})
}

func newDeleteModuleRequest(moduleId int) *httptest.ResponseRecorder {
	router := routes.GetRoutes()
	request, _ := http.NewRequest("DELETE", fmt.Sprintf("/modules/%d", moduleId), nil)
	response := httptest.NewRecorder()

	request.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(response, request)

	return response
}

package handlers

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"upcourse/config"
	"upcourse/internal/mocks"
	"upcourse/models"

	"github.com/gin-gonic/gin"
)

func TestDeleteCourse(t *testing.T) {
	mockCourse := mocks.NewFullCourse()
	defer teardown()

	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: fmt.Sprint(mockCourse.ID)})

	t.Run("Deletes a course and its children", func(t *testing.T) {
		var courseCount int64
		config.Conn.Model(models.Course{}).Count(&courseCount)

		DeleteCourse(ctx)

		var newCourseCount int64
		config.Conn.Model(models.Course{}).Count(&newCourseCount)
		if courseCount == newCourseCount {
			t.Errorf("Did not delete course: course count did not change")
		}

		err := config.Conn.First(&models.Course{}, mockCourse.ID).Error
		if err == nil {
			t.Errorf("Did not delete course: course still found in database")
		}

		var moduleCount int64
		config.Conn.Model(models.Module{}).Count(&moduleCount)
		if moduleCount > 0 {
			t.Errorf("Did not delete associated modules")
		}

		var moduleActivityCount int64
		config.Conn.Model(models.ModuleActivity{}).Count(&moduleActivityCount)
		if moduleActivityCount > 0 {
			t.Errorf("Did not delete associated module activities")
		}
	})
}

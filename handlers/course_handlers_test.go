package handlers

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	"upcourse/config"
	testHelpers "upcourse/internal/helpers"
	"upcourse/internal/mocks"
	"upcourse/models"
)

func TestGetCourse(t *testing.T) {
	mockCourse := mocks.FullCourse()
	defer testHelpers.Teardown()

	t.Run("returns a course and its modules and moduleActivities", func(t *testing.T) {
		params := map[string]string{"id": fmt.Sprint(mockCourse.ID)}
		w := testHelpers.NewRequest(params, "", GetCourse)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusOK)

		response := testHelpers.UnmarshalPayload(t, w, new(models.Course))
		responseCourse, ok := response.(*models.Course)
		if !ok {
			t.Errorf("error casting response element as course")
		}

		modulesLen := len(responseCourse.Modules)
		testHelpers.AssertEqualLength(t, modulesLen, len(mockCourse.Modules), "modules")

		for i := 0; i < modulesLen; i++ {
			gotModule := responseCourse.Modules[i]
			wantModule := mockCourse.Modules[i]

			modActivitiesLen := len(gotModule.ModuleActivities)
			testHelpers.AssertEqualLength(t, modActivitiesLen, len(wantModule.ModuleActivities), "module activities")

			for i := 0; i < modActivitiesLen; i++ {
				gotModuleActivity := gotModule.ModuleActivities[i]
				wantModuleActivity := wantModule.ModuleActivities[i]

				gotActivity := gotModuleActivity.Activity
				wantActivity := wantModuleActivity.Activity
				if !reflect.DeepEqual(gotActivity, wantActivity) {
					t.Errorf("got %v want %v for activity for module_activity %d", gotActivity, wantActivity, wantModuleActivity.ID)
				}

				gotModuleActivity.Activity, wantModuleActivity.Activity = nil, nil
				wantModuleActivity.ModuleId, wantModuleActivity.ActivityId = 0, 0
				if !reflect.DeepEqual(gotModuleActivity, wantModuleActivity) {
					t.Errorf("got %v want %v for module_activity %d", gotModuleActivity, wantModuleActivity, wantModuleActivity.ID)
				}
			}

			gotModule.ModuleActivities, wantModule.ModuleActivities = nil, nil
			wantModule.CourseId = 0
			if !reflect.DeepEqual(gotModule, wantModule) {
				t.Errorf("got %v want %v for module %d", gotModule, wantModule, wantModule.ID)
			}
		}

		responseCourse.Modules, mockCourse.Modules = nil, nil
		if !reflect.DeepEqual(responseCourse, mockCourse) {
			t.Errorf("got %v want %v for course", *responseCourse, mockCourse)
		}
	})

	t.Run("returns an error if the course is not found", func(t *testing.T) {
		params := map[string]string{"id": fmt.Sprint(mockCourse.ID + 1)}
		w := testHelpers.NewRequest(params, "", GetCourse)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusNotFound)

		response := testHelpers.UnmarshalErrors(t, w)
		testHelpers.AssertError(t, response[0], ErrNotFound)
	})

	t.Run("returns database errors if they occur", func(t *testing.T) {
		testHelpers.ForceError()
		defer testHelpers.ClearError()

		params := map[string]string{"id": fmt.Sprint(mockCourse.ID)}
		w := testHelpers.NewRequest(params, "", GetCourse)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusInternalServerError)

		response := testHelpers.UnmarshalErrors(t, w)
		testHelpers.AssertError(t, response[0], testHelpers.DatabaseErr)
	})
}

func TestGetCourses(t *testing.T) {
	t.Run("returns a list of courses with its modules", func(t *testing.T) {
		mockCourses := mocks.CourseList()
		defer testHelpers.Teardown()

		w := testHelpers.NewRequest(nil, "", GetCourses)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusOK)

		response := testHelpers.UnmarshalManyPayload(t, w, new(models.Course))
		if len(response) != len(mockCourses) {
			t.Errorf("got %d expected %d for length of results", len(response), len(mockCourses))
		}

		for i := 0; i < len(response); i++ {
			gotCourse, ok := response[i].(*models.Course)
			if !ok {
				t.Errorf("error casting response element as course")
			}
			wantCourse := mockCourses[i]

			gotModulesLen := len(gotCourse.Modules)
			wantModulesLen := len(wantCourse.Modules)
			if gotModulesLen != wantModulesLen {
				t.Errorf("got %d want %d for number of modules", gotModulesLen, wantModulesLen)
			}

			for i := 0; i < gotModulesLen; i++ {
				gotModule := gotCourse.Modules[i]
				wantModule := wantCourse.Modules[i]

				if gotModule.ModuleActivities != nil {
					t.Errorf("moduleActivities should not have been included in response but were")
				}

				wantModule.ModuleActivities = nil
				wantModule.CourseId = 0
				wantModule.CreatedAt, wantModule.UpdatedAt = time.Time{}, time.Time{}
				if !reflect.DeepEqual(gotModule, wantModule) {
					t.Errorf("got %v want %v for module %d", gotModule, wantModule, wantModule.ID)
				}
			}

			gotCourse.Modules, wantCourse.Modules = nil, nil
			wantCourse.CreatedAt, wantCourse.UpdatedAt = time.Time{}, time.Time{}
			if !reflect.DeepEqual(gotCourse, wantCourse) {
				t.Errorf("got %v want %v for course", gotCourse, wantCourse)
			}
		}
	})

	t.Run("retuns an array even if only one course is found", func(t *testing.T) {
		mocks.SimpleCourse()
		defer testHelpers.Teardown()

		w := testHelpers.NewRequest(nil, "", GetCourses)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusOK)

		response := testHelpers.UnmarshalManyPayload(t, w, new(models.Course))

		testHelpers.AssertEqualLength(t, len(response), 1, "courses")
	})

	t.Run("returns an array if no courses are found", func(t *testing.T) {
		w := testHelpers.NewRequest(nil, "", GetCourses)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusOK)

		response := testHelpers.UnmarshalManyPayload(t, w, new(models.Course))

		testHelpers.AssertEqualLength(t, len(response), 0, "courses")
	})

	t.Run("returns database errors if they occur", func(t *testing.T) {
		testHelpers.ForceError()
		defer testHelpers.ClearError()

		w := testHelpers.NewRequest(nil, "", GetCourses)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusInternalServerError)

		response := testHelpers.UnmarshalErrors(t, w)
		if response[0] != testHelpers.DatabaseErr {
			t.Errorf("got %s want %s for error message", response[0], testHelpers.DatabaseErr)
		}
	})
}

func TestCreateCourse(t *testing.T) {
	t.Run("it creates a new course", func(t *testing.T) {
		defer testHelpers.Teardown()

		var courseCount int64
		config.Conn.Model(models.Course{}).Count(&courseCount)

		newCourseInfo := `{
			"name": "Nursing 101",
			"institution": "Tampa Bay Nurses United University",
			"creditHours": 3,
			"length": 16,
			"goal": "8-10 hours"
		}`
		w := testHelpers.NewRequest(nil, newCourseInfo, CreateCourse)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusCreated)

		errs := testHelpers.UnmarshalErrors(t, w)
		if len(errs) != 0 {
			t.Errorf("got unexpected errors: %s", errs)
		}

		var newCount int64
		config.Conn.Model(models.Course{}).Count(&newCount)

		if newCount != (courseCount + 1) {
			t.Errorf("course count did not change")
		}
	})

	t.Run("it returns an error if a required field is missing", func(t *testing.T) {
		var courseCount int64
		config.Conn.Model(models.Course{}).Count(&courseCount)

		newCourseInfo := `{
			"creditHours": 3,
			"length": 16,
			"goal": "8-10 hours"
		}`
		w := testHelpers.NewRequest(nil, newCourseInfo, CreateCourse)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusBadRequest)

		errs := testHelpers.UnmarshalErrors(t, w)

		testHelpers.AssertError(t, errs[0], "Name is required")
		testHelpers.AssertError(t, errs[1], "Institution is required")

		var newCount int64
		config.Conn.Model(models.Course{}).Count(&newCount)

		if newCount != courseCount {
			t.Errorf("course count changed but should not have")
		}
	})

	t.Run("it returns an error if no body is sent", func(t *testing.T) {
		w := testHelpers.NewRequest(nil, "", CreateCourse)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusBadRequest)

		errs := testHelpers.UnmarshalErrors(t, w)
		testHelpers.AssertError(t, errs[0], ErrBadRequest)
	})

	t.Run("returns database errors if they occur", func(t *testing.T) {
		testHelpers.ForceError()
		defer testHelpers.ClearError()

		newCourseInfo := `{
			"name": "Nursing 101",
			"institution": "Tampa Bay Nurses United University",
			"creditHours": 3,
			"length": 16,
			"goal": "8-10 hours"
		}`
		w := testHelpers.NewRequest(nil, newCourseInfo, CreateCourse)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusInternalServerError)

		response := testHelpers.UnmarshalErrors(t, w)
		if response[0] != testHelpers.DatabaseErr {
			t.Errorf("got %s want %s for error message", response[0], testHelpers.DatabaseErr)
		}
	})
}

func TestUpdateCourse(t *testing.T) {
	mockCourse := mocks.SimpleCourse()
	defer testHelpers.Teardown()

	t.Run("it updates an existing course", func(t *testing.T) {
		var courseCount int64
		config.Conn.Model(models.Course{}).Count(&courseCount)

		params := map[string]string{"id": fmt.Sprint(mockCourse.ID)}
		newCourseInfo := `{
			"institution": "Tampa Bay Nurses United University",
			"creditHours": 3,
			"length": 16,
			"goal": "8-10 hours"
		}`
		w := testHelpers.NewRequest(params, newCourseInfo, UpdateCourse)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusOK)

		errs := testHelpers.UnmarshalErrors(t, w)
		if len(errs) != 0 {
			t.Errorf("got unexpected errors: %s", errs)
		}

		var newCount int64
		config.Conn.Model(models.Course{}).Count(&newCount)

		if newCount != courseCount {
			t.Errorf("patch request should not have changed course count")
		}

		var updatedCourse models.Course
		config.Conn.First(&updatedCourse, mockCourse.ID)

		if updatedCourse.Institution != "Tampa Bay Nurses United University" {
			t.Errorf("did not update the course record")
		}

		if updatedCourse.Name != "Foundations of Nursing" {
			t.Errorf("changed a field that should not have been changed")
		}
	})

	t.Run("it returns a 404 if course not found", func(t *testing.T) {
		params := map[string]string{"id": fmt.Sprint(mockCourse.ID + 1)}
		newCourseInfo := `{
			"institution": "Tampa Bay Nurses United University",
			"creditHours": 3,
			"length": 16,
			"goal": "8-10 hours"
		}`
		w := testHelpers.NewRequest(params, newCourseInfo, UpdateCourse)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusNotFound)
	})

	t.Run("it returns an error if no body is sent", func(t *testing.T) {
		params := map[string]string{"id": fmt.Sprint(mockCourse.ID)}
		w := testHelpers.NewRequest(params, "", UpdateCourse)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusBadRequest)
	})

	t.Run("returns database errors if they occur", func(t *testing.T) {
		testHelpers.ForceError()
		defer testHelpers.ClearError()

		params := map[string]string{"id": fmt.Sprint(mockCourse.ID)}
		newCourseInfo := `{
			"institution": "Tampa Bay Nurses United University",
			"creditHours": 3,
			"length": 16,
			"goal": "8-10 hours"
		}`
		w := testHelpers.NewRequest(params, newCourseInfo, UpdateCourse)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusInternalServerError)

		response := testHelpers.UnmarshalErrors(t, w)
		if response[0] != testHelpers.DatabaseErr {
			t.Errorf("got %s want %s for error message", response[0], testHelpers.DatabaseErr)
		}
	})
}

func TestDeleteCourse(t *testing.T) {
	t.Run("deletes the course and dependent records", func(t *testing.T) {
		mockCourse := mocks.FullCourse()
		defer testHelpers.Teardown()

		var courseCount int64
		config.Conn.Model(models.Course{}).Count(&courseCount)

		params := map[string]string{"id": fmt.Sprint(mockCourse.ID)}
		w := testHelpers.NewRequest(params, "", DeleteCourse)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusOK)

		errs := testHelpers.UnmarshalErrors(t, w)
		if len(errs) != 0 {
			t.Errorf("got unexpected errors: %s", errs)
		}

		var newCourseCount int64
		config.Conn.Model(models.Course{}).Count(&newCourseCount)
		if courseCount == newCourseCount {
			t.Errorf("Did not delete course")
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

	t.Run("returns database errors if they occur", func(t *testing.T) {
		testHelpers.ForceError()
		defer testHelpers.ClearError()

		mockCourse := mocks.FullCourse()
		defer testHelpers.Teardown()

		params := map[string]string{"id": fmt.Sprint(mockCourse.ID)}
		w := testHelpers.NewRequest(params, "", DeleteCourse)

		testHelpers.AssertStatusCode(t, w.Code, http.StatusInternalServerError)

		response := testHelpers.UnmarshalErrors(t, w)
		if response[0] != testHelpers.DatabaseErr {
			t.Errorf("got %s want %s for error message", response[0], testHelpers.DatabaseErr)
		}
	})
}

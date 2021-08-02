package handlers

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"reflect"
	"testing"

	"upcourse/config"
	"upcourse/internal/mocks"
	"upcourse/models"
)

type getCourseResponse struct {
	Data struct {
		Course         models.Course
		ActivityTotals []models.ActivityTotals
	}
	Message string
	Status  int
}

type getCoursesResponse struct {
	Data    []models.Course
	Message string
	Status  int
}

func TestGetCourse(t *testing.T) {
	mockCourse := mocks.NewFullCourse()
	defer teardown()

	t.Run("returns a course and its modules and moduleActivities", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{
			"id": fmt.Sprint(mockCourse.ID),
		})

		GetCourse(ctx)

		if w.Code != 200 {
			t.Errorf("expected response code to be 200, got %d", w.Code)
		}

		response := getCourseResponse{}
		json.Unmarshal(w.Body.Bytes(), &response)

		course := response.Data.Course
		assertResponseValue(t, course.ID, mockCourse.ID, "Id")
		assertResponseValue(t, course.Name, mockCourse.Name, "Name")
		assertResponseValue(t, course.CreditHours, mockCourse.CreditHours, "CreditHours")
		assertResponseValue(t, course.Length, mockCourse.Length, "Length")

		firstResponseModule := course.Modules[0]
		firstMockModule := mockCourse.Modules[0]
		assertResponseValue(t, firstResponseModule.ID, firstMockModule.ID, "Module Id")
		assertResponseValue(t, firstResponseModule.Name, firstMockModule.Name, "Module Name")
		assertResponseValue(t, firstResponseModule.Number, firstMockModule.Number, "Module Number")

		firstResponseModActivity := firstResponseModule.ModuleActivities[0]
		firstMockModActivity := firstMockModule.ModuleActivities[0]
		assertResponseValue(t, firstResponseModActivity.Input, firstMockModActivity.Input, "Module Activity Input")
		assertResponseValue(t, firstResponseModActivity.Notes, firstMockModActivity.Notes, "Module Activity Notes")

		firstResponseActivity := firstResponseModActivity.Activity
		firstMockActivity := firstMockModActivity.Activity
		assertResponseValue(t, firstResponseActivity.ID, firstMockActivity.ID, "Activity Id")
		assertResponseValue(t, firstResponseActivity.Description, firstMockActivity.Description, "Activity Description")
		assertResponseValue(t, firstResponseActivity.Metric, firstMockActivity.Metric, "Activity Metric")
		assertResponseValue(t, firstResponseActivity.Multiplier, firstMockActivity.Multiplier, "Activity Multiplier")
	})

	t.Run("returns an error if the course is not found", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{
			"id": fmt.Sprint(mockCourse.ID + 1),
		})

		GetCourse(ctx)

		if w.Code != 404 {
			t.Errorf("expected response code to be 404, got %d", w.Code)
		}

		response := unmarshalErrorResponse(t, w.Body)
		assertResponseValue(t, response.Errors, "Record not found", "Response message")
	})
}

func TestGetCourses(t *testing.T) {
	t.Run("returns a list of courses", func(t *testing.T) {
		mockCourses := mocks.NewCourseList()
		defer teardown()

		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{})

		GetCourses(ctx)

		if w.Code != 200 {
			t.Errorf("expected response code to be 200, got %d", w.Code)
		}

		response := getCoursesResponse{}
		json.Unmarshal(w.Body.Bytes(), &response)

		course := response.Data[0]
		assertResponseValue(t, course.ID, mockCourses[0].ID, "Id")
		assertResponseValue(t, course.Name, mockCourses[0].Name, "Name")

		firstModule := course.Modules[0]
		assertResponseValue(t, firstModule.Name, mockCourses[0].Modules[0].Name, "Module Name")
	})

	t.Run("retuns an array even if only one course is found", func(t *testing.T) {
		mocks.NewSimpleCourse()
		defer teardown()

		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{})

		GetCourses(ctx)

		if w.Code != 200 {
			t.Errorf("expected response code to be 200, got %d", w.Code)
		}

		response := getCoursesResponse{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assertResponseValue(t, len(response.Data), 1, "Number of results")
	})

	t.Run("returns an array if no courses are found", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{})

		GetCourses(ctx)

		if w.Code != 200 {
			t.Errorf("expected response code to be 200, got %d", w.Code)
		}

		response := getCoursesResponse{}
		json.Unmarshal(w.Body.Bytes(), &response)
		assertResponseValue(t, len(response.Data), 0, "Number of results")
	})
}

func TestCreateCourse(t *testing.T) {
	defer teardown()

	t.Run("it creates a new course", func(t *testing.T) {
		var courseCount int64
		config.Conn.Model(models.Course{}).Count(&courseCount)

		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{})
		newCourseInfo := `{
			"name": "Nursing 101",
			"institution": "Tampa Bay Nurses United University",
			"creditHours": 3,
			"length": 16,
			"goal": "8-10 hours"
		}`
		mocks.SetRequestBody(ctx, newCourseInfo)

		CreateCourse(ctx)

		if w.Code != 201 {
			t.Errorf("expected response code to be 201, got %d", w.Code)
		}

		response := unmarshalSuccessResponse(t, w.Body)
		assertResponseValue(t, response.Message, "Course created successfully", "Message")

		var newCount int64
		config.Conn.Model(models.Course{}).Count(&newCount)

		if newCount != (courseCount + 1) {
			t.Errorf("course count did not change")
		}
	})

	t.Run("it returns an error if a required field is missing", func(t *testing.T) {
		var courseCount int64
		config.Conn.Model(models.Course{}).Count(&courseCount)

		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{})
		newCourseInfo := `{
			"creditHours": 3,
			"length": 16,
			"goal": "8-10 hours"
		}`
		mocks.SetRequestBody(ctx, newCourseInfo)

		CreateCourse(ctx)

		if w.Code != 400 {
			t.Errorf("expected response code to be %d, got %d", 400, w.Code)
		}

		response := unmarshalResponse(t, w.Body)

		expected := []string{"Name is required", "Institution is required"}
		if !reflect.DeepEqual(response.Errors, expected) {
			t.Errorf("got %v, wanted %v for field Error messages", response.Errors, expected)
		}

		var newCount int64
		config.Conn.Model(models.Course{}).Count(&newCount)

		if newCount != courseCount {
			t.Errorf("course count changed but should not have")
		}
	})

	t.Run("it returns an error if no body is sent", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{})

		CreateCourse(ctx)

		if w.Code != 400 {
			t.Errorf("expected response code to be %d, got %d", 400, w.Code)
		}
	})
}

func TestUpdateCourse(t *testing.T) {
	mockCourse := mocks.NewSimpleCourse()
	defer teardown()

	t.Run("it updates an existing course", func(t *testing.T) {
		var courseCount int64
		config.Conn.Model(models.Course{}).Count(&courseCount)

		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{
			"id": fmt.Sprint(mockCourse.ID),
		})
		newCourseInfo := `{
			"institution": "Tampa Bay Nurses United University",
			"creditHours": 3,
			"length": 16,
			"goal": "8-10 hours"
		}`
		mocks.SetRequestBody(ctx, newCourseInfo)

		UpdateCourse(ctx)

		if w.Code != 200 {
			t.Errorf("expected response code to be %d, got %d", 200, w.Code)
		}

		response := unmarshalSuccessResponse(t, w.Body)
		assertResponseValue(t, response.Message, "Course updated successfully", "Message")

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
		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{
			"id": fmt.Sprint(mockCourse.ID + 1),
		})
		newCourseInfo := `{
			"institution": "Tampa Bay Nurses United University",
			"creditHours": 3,
			"length": 16,
			"goal": "8-10 hours"
		}`
		mocks.SetRequestBody(ctx, newCourseInfo)

		UpdateCourse(ctx)

		if w.Code != 404 {
			t.Errorf("expected response code to be %d, got %d", 404, w.Code)
		}
	})

	t.Run("it returns an error if no body is sent", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{
			"id": fmt.Sprint(mockCourse.ID),
		})
		UpdateCourse(ctx)

		if w.Code != 400 {
			t.Errorf("expected response code to be %d, got %d", 400, w.Code)
		}
	})
}

func TestDeleteCourse(t *testing.T) {
	t.Run("deletes the course and dependent records", func(t *testing.T) {
		mockCourse := mocks.NewFullCourse()
		defer teardown()

		var courseCount int64
		config.Conn.Model(models.Course{}).Count(&courseCount)

		ctx := mocks.NewMockContext(httptest.NewRecorder(), map[string]string{
			"id": fmt.Sprint(mockCourse.ID),
		})

		DeleteCourse(ctx)

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

	t.Run("returns a successful http response", func(t *testing.T) {
		mockCourse := mocks.NewFullCourse()
		defer teardown()

		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{
			"id": fmt.Sprint(mockCourse.ID),
		})

		DeleteCourse(ctx)

		if w.Code != 200 {
			t.Errorf("expected response code to be 200, got %d", w.Code)
		}

		response := unmarshalSuccessResponse(t, w.Body)
		assertResponseValue(t, response.Message, "Course deleted successfully", "Message")
	})
}

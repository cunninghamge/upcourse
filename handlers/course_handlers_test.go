package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http/httptest"
	"reflect"
	"testing"

	"upcourse/config"
	"upcourse/internal/mocks"
	"upcourse/models"
)

func TestGetCourse(t *testing.T) {
	mockCourse := mocks.FullCourse()
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

		var response struct {
			Data struct {
				Type          string
				ID            int
				Attributes    models.Course
				Relationships map[string][]struct {
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
		}
		if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
			t.Errorf("error unmarshaling json response: %v", err)
		}
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
			"id": fmt.Sprint(mockCourse.ID + 1),
		})

		GetCourse(ctx)

		response := unmarshalResponse(t, w.Body)
		if response.Errors[0] != err {
			t.Errorf("got %s want %s for error message", response.Errors[0], err)
		}
	})
}

func TestGetCourses(t *testing.T) {
	t.Run("returns a list of courses", func(t *testing.T) {
		mockCourses := mocks.CourseList()
		defer teardown()

		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{})

		GetCourses(ctx)

		if w.Code != 200 {
			t.Errorf("expected response code to be 200, got %d", w.Code)
		}

		var response struct {
			Data []struct {
				Type          string
				ID            int
				Attributes    models.Course
				Relationships map[string][]SerializedResource
			}
		}
		if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
			t.Errorf("error unmarshaling json response: %v", err)
		}

		if len(response.Data) != len(mockCourses) {
			t.Errorf("got %d expected %d for length of results", len(response.Data), len(mockCourses))
		}
	})

	t.Run("retuns an array even if only one course is found", func(t *testing.T) {
		mocks.SimpleCourse()
		defer teardown()

		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{})

		GetCourses(ctx)

		if w.Code != 200 {
			t.Errorf("expected response code to be 200, got %d", w.Code)
		}

		var response struct {
			Data []SerializedResource
		}
		if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
			t.Errorf("error unmarshaling json response: %v", err)
		}

		assertResponseValue(t, len(response.Data), 1, "Number of results")
	})

	t.Run("returns an array if no courses are found", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{})

		GetCourses(ctx)

		if w.Code != 200 {
			t.Errorf("expected response code to be 200, got %d", w.Code)
		}

		var response struct {
			Data []SerializedResource
		}
		if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
			t.Errorf("error unmarshaling json response: %v", err)
		}

		assertResponseValue(t, len(response.Data), 0, "Number of results")
	})

	t.Run("returns database errors if they occur", func(t *testing.T) {
		err := "some database error"
		config.Conn.Error = errors.New(err)
		defer func() {
			config.Conn.Error = nil
		}()

		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{})

		GetCourses(ctx)

		response := unmarshalResponse(t, w.Body)
		if response.Errors[0] != err {
			t.Errorf("got %s want %s for error message", response.Errors[0], err)
		}
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

		response := unmarshalResponse(t, w.Body)
		if len(response.Errors) != 0 {
			t.Errorf("got unexpected errors: %s", response.Errors)
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

	t.Run("returns database errors if they occur", func(t *testing.T) {
		err := "some database error"
		config.Conn.Error = errors.New(err)
		defer func() {
			config.Conn.Error = nil
		}()

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

		response := unmarshalResponse(t, w.Body)
		if response.Errors[0] != err {
			t.Errorf("got %s want %s for error message", response.Errors[0], err)
		}
	})
}

func TestUpdateCourse(t *testing.T) {
	mockCourse := mocks.SimpleCourse()
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

		response := unmarshalResponse(t, w.Body)
		if len(response.Errors) != 0 {
			t.Errorf("got unexpected errors: %s", response.Errors)
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
		mockCourse := mocks.FullCourse()
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
		mockCourse := mocks.FullCourse()
		defer teardown()

		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{
			"id": fmt.Sprint(mockCourse.ID),
		})

		DeleteCourse(ctx)

		if w.Code != 200 {
			t.Errorf("expected response code to be 200, got %d", w.Code)
		}

		response := unmarshalResponse(t, w.Body)
		if len(response.Errors) != 0 {
			t.Errorf("got unexpected errors: %s", response.Errors)
		}
	})

	t.Run("returns database errors if they occur", func(t *testing.T) {
		err := "some database error"
		config.Conn.Error = errors.New(err)
		defer func() {
			config.Conn.Error = nil
		}()

		mockCourse := mocks.FullCourse()
		defer teardown()

		w := httptest.NewRecorder()
		ctx := mocks.NewMockContext(w, map[string]string{
			"id": fmt.Sprint(mockCourse.ID),
		})

		DeleteCourse(ctx)

		response := unmarshalResponse(t, w.Body)
		if response.Errors[0] != err {
			t.Errorf("got %s want %s for error message", response.Errors[0], err)
		}
	})
}

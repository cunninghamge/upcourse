package testing

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"testing"

	"github.com/Pallinder/go-randomdata"
	"github.com/gin-gonic/gin"

	"upcourse/config"
	"upcourse/models"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	config.Connect()
	code := m.Run()
	os.Exit(code)
}

func assertResponseValue(t *testing.T, got, want interface{}, field string) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v for field %q", got, want, field)
	}
}

func teardown() {
	config.Conn.Where("1=1").Delete(&models.ModuleActivity{})
	config.Conn.Where("1=1").Delete(&models.Module{})
	config.Conn.Where("1=1").Delete(&models.Course{})
	config.Conn.Where("custom=true").Delete(&models.Activity{})
}

func newSimpleCourse() *models.Course {
	course := models.Course{ID: 1, Name: "Foundations of Nursing"}
	config.Conn.Create(&course)
	return &course
}

func newFullCourse() *models.Course {
	course := models.Course{
		Name:        "Test Course",
		Institution: randomdata.LastName() + " University",
		CreditHours: randomdata.Number(5),
		Length:      randomdata.Number(16),
	}
	config.Conn.Create(&course)

	var modules []models.Module
	for i := 1; i < 5; i++ {
		modules = append(modules, models.Module{
			Name:     "Module " + strconv.Itoa(i),
			Number:   i,
			CourseId: course.ID,
		})
	}
	config.Conn.Create(&modules)

	var moduleActivities []models.ModuleActivity
	for _, module := range modules {
		for i := 1; i < 5; i++ {
			moduleActivities = append(moduleActivities, models.ModuleActivity{
				Input:      randomdata.Number(200),
				Notes:      "notes",
				ActivityId: randomdata.Number(13) + 1,
				ModuleId:   module.ID,
			})
		}
	}
	config.Conn.Create(&moduleActivities)

	config.Conn.Preload("Modules.ModuleActivities.Activity").First(&course)
	return &course
}

func newCourseList() []models.Course {
	var courseList []models.Course
	for i := 0; i < 3; i++ {
		courseList = append(courseList, models.Course{
			Name: "Test Course " + strconv.Itoa(i),
			Modules: []models.Module{
				{
					Name: "Module 1",
				},
				{
					Name: "Module 2",
				},
				{
					Name: "Module 3",
				},
			},
		})
	}
	config.Conn.Create(&courseList)

	return courseList
}

func newModule() models.Module {
	course := newSimpleCourse()

	module := models.Module{
		Name:     "Test Module",
		Number:   1,
		CourseId: course.ID,
	}
	config.Conn.Create(&module)

	var moduleActivities []models.ModuleActivity
	for i := 0; i < 4; i++ {
		moduleActivities = append(moduleActivities, models.ModuleActivity{
			ID:         i + 1,
			Input:      randomdata.Number(200),
			Notes:      "notes",
			ActivityId: i + 1,
			ModuleId:   module.ID,
		})
	}
	config.Conn.Create(&moduleActivities)

	config.Conn.Preload("ModuleActivities.Activity").First(&module)
	return module
}

func coreActivities() []models.Activity {
	var activities []models.Activity
	config.Conn.Find(&activities).Where("custom = false")
	return activities
}

func UnmarshalError(t *testing.T, response io.Reader) ErrorResponse {
	t.Helper()
	body, _ := ioutil.ReadAll(response)

	responseError := ErrorResponse{}
	err := json.Unmarshal([]byte(body), &responseError)
	if err != nil {
		t.Errorf("Error marshaling JSON response\nError: %v", err)
	}

	return responseError
}

type ErrorResponse struct {
	Status int
	Errors string
}

func UnmarshalErrors(t *testing.T, response io.Reader) MultipleErrorResponse {
	t.Helper()
	body, _ := ioutil.ReadAll(response)

	responseErrors := MultipleErrorResponse{}
	err := json.Unmarshal([]byte(body), &responseErrors)
	if err != nil {
		t.Errorf("Error marshaling JSON response\nError: %v", err)
	}

	return responseErrors
}

type MultipleErrorResponse struct {
	Status int
	Errors []string
}

func UnmarshalSuccess(t *testing.T, response io.Reader) SuccessResponse {
	t.Helper()

	body, _ := ioutil.ReadAll(response)
	successResponse := SuccessResponse{}
	err := json.Unmarshal([]byte(body), &successResponse)

	if err != nil {
		t.Errorf("Error marshaling JSON response\nError: %v", err)
	}

	return successResponse
}

type SuccessResponse struct {
	Status  int
	Message string
}

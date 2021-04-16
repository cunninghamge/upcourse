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

	"course-chart/config"
	"course-chart/models"
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

func newMockCourse() *models.Course {
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

func newMockCourseList() []models.CourseIdentifier {
	var courseList []models.Course
	for i := 0; i < 3; i++ {
		courseList = append(courseList, models.Course{
			Name: "Test Course " + strconv.Itoa(i),
		})
	}
	config.Conn.Create(&courseList)

	var courses []models.CourseIdentifier
	for _, course := range courseList {
		courses = append(courses, models.CourseIdentifier{
			ID:   course.ID,
			Name: course.Name,
		})
	}
	return courses
}

func newMockModule() models.Module {
	course := models.Course{}
	config.Conn.Create(&course)

	module := models.Module{
		Name:     "Test Module",
		Number:   1,
		CourseId: course.ID,
	}
	config.Conn.Create(&module)

	var moduleActivities []models.ModuleActivity
	for i := 1; i < 5; i++ {
		moduleActivities = append(moduleActivities, models.ModuleActivity{
			Input:      randomdata.Number(200),
			Notes:      "notes",
			ActivityId: randomdata.Number(13) + 1,
			ModuleId:   module.ID,
		})
	}
	config.Conn.Create(&moduleActivities)

	config.Conn.Preload("ModuleActivities.Activity").First(&module)
	return module
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

package main

import (
	"context"
	"course-chart/config"
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

var ModuleActivityFactory = factory.NewFactory(
	&ModuleActivity{},
).SeqInt("Id", func(n int) (interface{}, error) {
	return n, nil
}).Attr("Input", func(args, factory.Args) (interface{}, error) {
	return randomdata.Number(100), nil
}).Attr("Notes", func(args, factory.Args) (interface{}, error) {
	return "notes", nil
}).Attr("ActivityId", func(args, factory.Args) (interface{}, error) {
	return randomdata.Number(14), nil
}).Attr("ModuleId", func(args, factory.Args) (interface{}, error) {
	if parent := args.Parent(); parent != nil {
		return parent.Instance(), nil
	}
	return nil, nil
})

var ModuleFactory = factory.NewFactory(
	&Module{},
).SeqInt("Id", func(n int) (interface{}, error) {
	return n, nil
}).SeqInt("Number", func(n int) (interface{}, error) {
	return n, nil
}).Attr("Name", func(args, factory.Args) (interface{}, error) {
	module := args.Instance().(*Module)
	return "Module" + module.Number, nil
}).Attr("CourseId", func(args factory.Args) (interface{}, error) {
	if parent := args.Parent(); parent != nil {
		return parent.Instance(), nil
	}
	return nil, nil
}).OnCreate(func(args factory.Args) error {
	db := args.Context().Value("db").(*gorm.DB)
	return db.Create(args.Instance()).Error
}).SubSliceFactory("ModuleActivities", ModuleActivityFactory, func() int { return 6 })

type Course struct {
	Id          int
	Name        string
	Institution string
	CreditHours int
	Lengh       int
	Goal        string
	Modules     *[]Module
}

type Module struct {
	Id               int
	Name             string
	Number           int
	Course           *Course
	ModuleActivities *[]ModuleActivity
}

type ModuleActivity struct {
	Id         int
	Input      int
	Notes      string
	ActivityId int
	Module     *Module
}

var CourseFactory = factory.NewFactory(
	&Course{},
).SeqInt("Id", func(n int) (interface{}, error) {
	return n, nil
}).Attr("Name", func(args, factory.Args) (interface{}, error) {
	course := args.Instance().(*Course)
	return "Test Course" + course.Id, nil
}).Attr("Institution", func(args, factory.Args) (interface{}, error) {
	return randomdata.LastName() + "University", nil
}).Attr("CreditHours", func(args, factory.Args) (interface{}, error) {
	return randomdata.Number(5), nil
}).Attr("Length", func(args, factory.Args) (interface{}, error) {
	return randomdata.Number(16), nil
}).Attr("Goal", func(args, factory.Args) (interface{}, error) {
	return "16-18 hrs", nil
}).OnCreate(func(args factory.Args) error {
	db := args.Context().Value("db").(*gorm.DB)
	return db.Create(args.Instance()).Error
}).SubSliceFactory("Modules", ModuleFactory, func() int { return 4 })

func CreateCourse() *Course {
	db := config.Conn
	tx := db.Begin()
	ctx := context.WithValue(context.Background(), "db", tx)
	v, err := CourseFactory.CreateWithContext(ctx)
	if err != nil {
		panic(err)
	}
	course := v.(*Course)
	tx.Commit()
	fmt.Println(course)
	return course
}

type User struct {
	ID       int
	Name     string
	Location string
}

// 'Location: "Tokyo"' is default value.
var UserFactory = factory.NewFactory(
	&User{},
).SeqInt("ID", func(n int) (interface{}, error) {
	return n, nil
}).Attr("Name", func(args factory.Args) (interface{}, error) {
	return randomdata.FullName(randomdata.RandomGender), nil
}).Attr("Location", func(args factory.Args) (interface{}, error) {
	return randomdata.City(), nil
})

func n() {
	for i := 0; i < 3; i++ {
		user := UserFactory.MustCreate().(*User)
		fmt.Println("ID:", user.ID, " Name:", user.Name, " Location:", user.Location)
	}
}

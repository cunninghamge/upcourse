[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/cunninghamge/upcourse.svg)](https://github.com/gomods/athens)
[![CircleCI](https://circleci.com/gh/cunninghamge/upcourse.svg?style=svg)](https://circleci.com/gh/circleci/circleci-docs)


<h1 align="center">Upcourse</h1>

<p style="margin: 5px">
Upcourse is an instructional design tool to assist education professionals in designing and mapping both new and existing courses. Users can build courses and modules around the amount of time students should spend on a variety of different learning activities (such as readings, lectures, and exams).
</p>
<br>
<p style="margin: 5px">
Like many software tools, Upcourse was born from a desire to replace a difficult-to-use spreadsheet with a more interactive and user-friendly solution. The original app, CourseChart, was built by a team of students at the Turing School of Software & Design as our capstone project.Our team worked closely with an instructional design expert to ensure that our product met the needs of professionals in the field. 
</p>

## Table of Contents
 - [About](#about)
 - [Database Schema](#database-schema)
 - [API Contract](#api-contract)
 - [Technologies Used](#technologies-used)
 - [Local Setup](#local-setup)

## Database Schema

*coming soon*

## API Contract

*coming soon*

## Technologies Used

* [Go](https://golang.org/)
* [Gin](https://github.com/gin-gonic/gin)
* [GORM](https://gorm.io/)
* [PostgreSQL](https://www.postgresql.org/)
* [CircleCI](https://circleci.com/)
* [Heroku](https://heroku.com)

## Local Setup
  To run the project in your local environment, please follow the instructions below:

  1. Clone the repository:<br>
    `git clone git@github.com:cunninghamge/upcourse.git`
    `cd upcourse`
  2. Install Go with<br>
    `brew install go`
  3. Set up the database:<br>
    a. create the databases:<br>
      `psql -c "CREATE DATABASE upcourse`<br>
      `psql -c "CREATE DATABASE upcourse_test`<br>
    b. run the database migrations:<br>
      `go run ./database/migrate`<br>
      `GIN_MODE=test go run ./database/migrate`<br>
      `go run ./database/seed`<br>
      `GIN_MODE=test go run ./database/seed`<br>
  4. To launch a local server:<br>
    `go run upcourse`<br>
    Once the server is running you can send requests to `localhost:8080`<br>
    ex: `http://localhost:8080/v1/courses/1`
  5. To run tests:<br>
    `go test ./...`<br>
    or, to run tests with a detailed coverage report, run:<br>
    `go test ./... -v -coverprofile cover.out`<br>
    `go tool cover -html=cover.out`<br>

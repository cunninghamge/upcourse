# Course Chart BE

## Table of Contents
 - [Description](#description)
 - [System Design](#database-schema)
 - [Database Schema](#database-schema)
 - [API Contract](#api-contract)
 - [Technologies Used](#technologies-used)
 - [Local Setup](#local-setup)
 - [Learning Goals](#learning-goals)
 - [Authors](#authors)
 - [Statistics](#statistics)

## Description

## System Design

## Database Schema

## API Contract

## Technologies Used

## Local Setup
  To run the project in your local environment, please follow the instructions below:

  1. Clone the repository:<br>
    `git clone git@github.com:CourseChart/course-chart-be.git`
    `cd course-chart-be`
  2. Install Go with<br>
    `brew install go`
  3. Set up the database:<br>
    a. create the databases:
      `psql -c "CREATE DATABASE course_chart`
      `psql -c "CREATE DATABASE course_chart_test`
    b. cd into the migrations directory:
      `cd migrations`
    c. initialize the go-pg/migrations package:
      `go run *.go init`
    d. run the migrations:
      `go run *.go`
    e. return to the root directory:
      `cd ..`
  4. To launch a local server:<br>
    `go run course-chart`<br>
    Once the server is running you can send requests to `localhost:8080`<br>
    ex: `http://localhost:8080`
  5. To run tests and view the test coverage report:<br>
    `go test -cover` 


## Learning Goals

## Authors

## Statistics
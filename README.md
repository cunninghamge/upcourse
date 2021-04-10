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
    a. create the databases:<br>
      `psql -c "CREATE DATABASE course_chart`
      `psql -c "CREATE DATABASE course_chart_test`<br>
    b. cd into the migrations directory:<br>
      `cd migrations`<br>
    c. initialize the go-pg/migrations package:<br>
      `go run *.go init`<br>
    d. run the migrations:<br>
      `go run *.go`<br>
    e. return to the root directory:<br>
      `cd ..`<br>
    f. Create a `.env` file and add the following:
      ```
      DB_USER: <your postgres username>
      DB_ADDRESS: "localhost:5432"
      DB_NAME: "course_chart"
      ```
      To get your postgres username, enter the following in the command line:<br>
      `psql postgres`<br>
      `\du`
  4. To launch a local server:<br>
    `go run course-chart`<br>
    Once the server is running you can send requests to `localhost:8080`<br>
    ex: `http://localhost:8080`
  5. To run tests and view the test coverage report:<br>
    `go test -cover` 


## Learning Goals

## Authors

## Statistics
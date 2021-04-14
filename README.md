[![CircleCI](https://circleci.com/gh/CourseChart/course-chart-be.svg?style=svg)](https://circleci.com/gh/circleci/circleci-docs)

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

[API Endpoint Documentation](https://documenter.getpostman.com/view/14310262/TzJpgevK)

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
      `psql -c "CREATE DATABASE course_chart`<br>
      `psql -c "CREATE DATABASE course_chart_test`<br>
    b. Create a `.env` file and add the following:
      ```
      POSTGRES_USER: <your postgres username>
      POSTGRES_ADDRESS: "localhost:5432"
      POSTGRES_NAME: "course_chart"
      PORT: "8080"
      ```
      To get your postgres username, enter the following in the command line:<br>
      `psql postgres`<br>
      `\du`
    c. run the database migrations:<br>
      `go run ./migrations`<br>
  4. To launch a local server:<br>
    `go run course-chart`<br>
    Once the server is running you can send requests to `localhost:8080`<br>
    ex: `http://localhost:8080`
  5. To run tests and view the test coverage report:<br>
    `go test -cover` 


## Learning Goals

## Authors

## Statistics

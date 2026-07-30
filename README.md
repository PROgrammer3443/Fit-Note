# FitNote

FitNote is a full-stack fitness tracking web application built with **Go**. It helps users monitor their health by tracking workouts, calories, and personal fitness information through a clean web interface.

The project was built to practice backend development with Go, user authentication, database management, server-side rendering, and web application architecture.

---

# Features

## User Authentication

* User registration
* Secure login and logout
* Session-based authentication
* Protected routes
* Persistent login sessions

---

## User Profiles

Users can store personal fitness information including:

* Name
* Age
* Height
* Weight
* Activity Level
* Daily calorie goals

---

## Workout Tracking

Create and complete daily workouts.

Features include:

* Workout plans
* Exercise library
* Mark exercises as completed
* Swap exercises
* Track calories burned
* Daily workout progress

---

## Nutrition Tracking

Track daily calorie intake.

Features include:

* Food logging
* Calories consumed
* Daily calorie totals
* Nutrition history

---

## Dashboard

The home page displays a summary of the user's progress.

Examples include:

* Calories Consumed
* Calories Burned
* Daily Progress
* Current Workout
* Personal Statistics

---

# Technologies Used

## Backend

* Go
* net/http
* html/template
* database/sql

## Database

* MySQL

## Authentication

* scs Session Manager
* MySQL Session Store

## Frontend

* HTML
* CSS
* JavaScript

---

# Project Structure

```text
FitNote
├── cmd
│   └── web
│       ├── handlers.go
│       ├── helpers.go
│       ├── middleware.go
│       ├── routes.go
│       └── main.go
│
├── internal
│   ├── models
│   ├── validator
│   └── ...
│
├── migrations
│
├── ui
│   ├── html
│   │   ├── pages
│   │   ├── partials
│   │   └── base.tmpl.html
│   │
│   └── static
│       ├── css
│       ├── js
│       └── img
│
└── go.mod
```

---

# Main Features

## Authentication

* Sign Up
* Login
* Logout
* Session Management

---

## Fitness

* Daily Workouts
* Exercise Tracking
* Workout Completion
* Exercise Swapping
* Calories Burned

---

## Nutrition

* Food Logging
* Calorie Tracking
* Daily Intake

---

## Dashboard

* User Statistics
* Progress Overview
* Workout Summary
* Calorie Summary

---

# Database

The application uses **MySQL** for persistent storage.

Data stored includes:

* Users
* Sessions
* Food Logs
* Workout Data
* Exercises
* Daily Statistics

---

# Learning Objectives

This project was created to gain practical experience with:

* Backend web development in Go
* REST-style request handling
* Authentication and authorization
* Session management
* SQL database design
* Database migrations
* HTML templating
* Form validation
* Middleware
* MVC-style project organization

---

# Future Improvements

Potential future additions include:

* Weekly and monthly progress charts
* Exercise recommendations
* Nutrition recommendations
* AI-powered workout suggestions
* Barcode food scanner
* Mobile-friendly responsive design
* Email verification
* Password reset
* Export workout history
* REST API
* Docker deployment

---

# Screenshots

Add screenshots of:

* Login Page
* Dashboard
* Workout Page
* Food Tracker
* User Profile

---

# Goal

FitNote was built as a learning project to explore modern backend development with Go while creating a practical fitness tracking application. It combines authentication, database management, server-side rendering, and health tracking into a complete full-stack web application.

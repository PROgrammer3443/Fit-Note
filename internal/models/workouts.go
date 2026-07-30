package models

import (
	"database/sql"
	"time"
)

type Workout struct {
	ID int
	User_id int
	Date time.Time
	Calories int
}

type WorkoutModel struct {
	DB *sql.DB
}
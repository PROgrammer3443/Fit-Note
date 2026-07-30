package models

import (
	"database/sql"
)

type Logs struct {
	Id int
	Meal_id int
	Food_name string
	Calories int
}

type LogModel struct {
	DB *sql.DB
}
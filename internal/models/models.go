package models

import (
	"database/sql"
	"errors"
	"time"
)

func (m *UserModel) Insert(username, email, hashedPass string, age, height, weight, ca_level, ta_level int) error {
	stmt := "INSERT INTO users (name, email, hashed_password, age, ta_level, ca_level, weight, height) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := m.DB.Exec(stmt, username, email, hashedPass, age, ta_level, ca_level, weight, height)
	return err
}

var ErrNoRecord = errors.New("models: no matching record found")

func (m *UserModel) Get(id int) (*User, error) {
    stmt := `
        SELECT id, name, email, hashed_password,
               age, weight, height,
               ca_level, ta_level
        FROM users
        WHERE id = ?
    `

    u := &User{}

    err := m.DB.QueryRow(stmt, id).Scan(
        &u.ID,
        &u.Name,
        &u.Email,
        &u.Hashed_password,
        &u.Age,
        &u.Weight,
        &u.Height,
        &u.Ca_level,
        &u.Ta_level,
    )

    if err != nil {
        if err == sql.ErrNoRows {
            return nil, ErrNoRecord
        }
        return nil, err
    }

    return u, nil
}

func (f *FoodModel) Insert(meal_Type string, user_id int, date time.Time) error {
	stmt := "INSERT INTO food_logs (user_id, date, meal_type, total_calories) VALUES (?, ?, ?, ?)"

	_, err := f.DB.Exec(stmt, user_id, date, meal_Type, 0)
	return err
}

func (m *FoodModel) GetByUser(userID int) ([]*Food, error) {
    stmt := `
        SELECT meal_id, meal_type, COALESCE(total_calories, 0), user_id, date
        FROM food_logs
        WHERE user_id = ?
        ORDER BY date DESC
    `

    rows, err := m.DB.Query(stmt, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    foods := []*Food{}

    for rows.Next() {
        f := &Food{}

        err := rows.Scan(
            &f.Meal_id,
            &f.Meal_Type,
            &f.Total_calories,
            &f.User_id,
            &f.Date,
        )
        if err != nil {
            return nil, err
        }

        foods = append(foods, f)
    }

    return foods, rows.Err()
}

func (m *LogModel) Insert(food_name string, meal_id, calories int) error {
	stmt := "INSERT INTO log_entries (meal_ID, food_name, calories) VALUES (?, ?, ?)"

	_, err := m.DB.Exec(stmt, meal_id, food_name, calories)
	return err
}

func (m *LogModel) GetByMeal(mealID int) ([]*Logs, error) {
    stmt := `
        SELECT id, meal_id, food_name, calories
        FROM log_entries
        WHERE meal_id = ?
        ORDER BY id ASC
    `

    rows, err := m.DB.Query(stmt, mealID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    logs := []*Logs{}

    for rows.Next() {
        l := &Logs{}

        err := rows.Scan(
            &l.Id,
            &l.Meal_id,
            &l.Food_name,
            &l.Calories,
        )
        if err != nil {
            return nil, err
        }

        logs = append(logs, l)
    }

    return logs, rows.Err()
}

func (w *WorkoutModel) Insert(user_id, calories int, workout_date time.Time) error {
	stmt := "INSERT INTO workout_logs (user_id, workout_date, calories_burned) VALUES (?, ?, ?)"

	_, err := w.DB.Exec(stmt, user_id, workout_date, calories)
	return err
}

func (w *WorkoutModel) GetBurnedByUser(userID int) (int, error) {
    stmt := `
        SELECT COALESCE(SUM(calories_burned), 0)
        FROM workout_logs
        WHERE user_id = ?
          AND DATE(workout_date) = CURDATE()
    `

    var calories int
    err := w.DB.QueryRow(stmt, userID).Scan(&calories)
    if err != nil {
        if err == sql.ErrNoRows {
            return 0, nil 
        }
        return 0, err
    }

    return calories, nil
}   
package models

import (
	"database/sql"
	"time"
)

type Food struct {
	Meal_id int
	User_id int
	Date time.Time
	Meal_Type string
	Total_calories int
}

type FoodModel struct {
	DB *sql.DB
}


func (m *FoodModel) GetTotalCaloriesByUser(userID int) (int, error) {
    stmt := `
        SELECT COALESCE(SUM(le.calories), 0)
        FROM food_logs fl
        LEFT JOIN log_entries le
            ON fl.meal_id = le.meal_id
        WHERE fl.user_id = ?
          AND DATE(fl.date) = CURDATE()
    `

    var total int

    err := m.DB.QueryRow(stmt, userID).Scan(&total)
    if err != nil {
        return 0, err
    }

    return total, nil
}

func (m *FoodModel) GetLogsByUser(userID int) ([]*Food, error) {
    query := `
        SELECT meal_id,
               meal_type,
               COALESCE(total_calories, 0),
               user_id,
               date
        FROM food_logs
        WHERE user_id = ?
          AND DATE(date) = CURDATE()
        ORDER BY date DESC
    `

    rows, err := m.DB.Query(query, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    meals := []*Food{}

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

        meals = append(meals, f)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return meals, nil
}
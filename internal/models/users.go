package models

import "database/sql"

type User struct {
	ID int
	Name  string
	Email string
	Hashed_password string
	Age int
	Weight int
	Height int
	Ca_level int
	Ta_level int
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) GetByEmail(email string) (*User, error) {
    stmt := `
        SELECT id, name, email, hashed_password, age, ta_level, ca_level, weight, height
        FROM users
        WHERE email = ?
    `

    u := &User{}

    err := m.DB.QueryRow(stmt, email).Scan(
        &u.ID,
        &u.Name,
        &u.Email,
        &u.Hashed_password,
        &u.Age,
        &u.Ta_level,
        &u.Ca_level,
        &u.Weight,
        &u.Height,
    )

    if err != nil {
        return nil, err
    }

    return u, nil
}
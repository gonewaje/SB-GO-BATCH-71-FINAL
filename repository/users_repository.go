package repository

import (
	"database/sql"
	"errors"
)

type userRow struct {
	ID       int
	Name     string
	Email    string
	Password string
	Role     string
}

func CreateUser(db *sql.DB, name, email, password, role string) (int, error) {
	var id int
	err := db.QueryRow(`
		INSERT INTO users (name, email, password, role)
		VALUES ($1,$2,$3,$4) RETURNING id
	`, name, email, password, role).Scan(&id)
	return id, err
}

func GetUserByEmail(db *sql.DB, email string) (userRow, error) {
	var u userRow
	err := db.QueryRow(`
		SELECT id, name, email, password, role
		FROM users WHERE email=$1
	`, email).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Role)
	if err == sql.ErrNoRows {
		return u, errors.New("not found")
	}
	return u, err
}

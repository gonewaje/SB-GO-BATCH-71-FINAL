package repository

import (
	"database/sql"
	"errors"

	"gonewaje/final/structs"
)

func ListRestaurants(db *sql.DB) ([]structs.Restaurant, error) {
	rows, err := db.Query(`SELECT id, name, address, phone FROM restaurants ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []structs.Restaurant
	for rows.Next() {
		var r structs.Restaurant
		if err := rows.Scan(&r.ID, &r.Name, &r.Address, &r.Phone); err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
}

func CreateRestaurant(db *sql.DB, name, address, phone string) (int, error) {
	var id int
	err := db.QueryRow(`
		INSERT INTO restaurants (name, address, phone) VALUES ($1,$2,$3) RETURNING id
	`, name, address, phone).Scan(&id)
	return id, err
}

func GetRestaurant(db *sql.DB, id int) (structs.Restaurant, error) {
	var r structs.Restaurant
	err := db.QueryRow(`
		SELECT id, name, address, phone FROM restaurants WHERE id=$1
	`, id).Scan(&r.ID, &r.Name, &r.Address, &r.Phone)
	if err == sql.ErrNoRows {
		return r, errors.New("not found")
	}
	return r, err
}

func UpdateRestaurant(db *sql.DB, id int, name, address, phone string) (bool, error) {
	res, err := db.Exec(`
		UPDATE restaurants SET name=$1, address=$2, phone=$3 WHERE id=$4
	`, name, address, phone, id)
	if err != nil {
		return false, err
	}
	n, _ := res.RowsAffected()
	return n > 0, nil
}

func DeleteRestaurant(db *sql.DB, id int) (bool, error) {
	res, err := db.Exec(`DELETE FROM restaurants WHERE id=$1`, id)
	if err != nil {
		return false, err
	}
	n, _ := res.RowsAffected()
	return n > 0, nil
}

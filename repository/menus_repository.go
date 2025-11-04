package repository

import (
	"database/sql"
	"errors"

	"gonewaje/final/structs"
)

func ListMenuItemsByRestaurant(db *sql.DB, restaurantID int) ([]structs.MenuItem, error) {
	rows, err := db.Query(`
		SELECT id, restaurant_id, name, price, available
		FROM menu_items WHERE restaurant_id=$1
	`, restaurantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []structs.MenuItem
	for rows.Next() {
		var m structs.MenuItem
		if err := rows.Scan(&m.ID, &m.RestaurantID, &m.Name, &m.Price, &m.Available); err != nil {
			return nil, err
		}
		res = append(res, m)
	}
	return res, nil
}

func CreateMenuItem(db *sql.DB, restaurantID int, name string, price int, available bool) (int, error) {
	var id int
	err := db.QueryRow(`
		INSERT INTO menu_items (restaurant_id, name, price, available)
		VALUES ($1,$2,$3,$4) RETURNING id
	`, restaurantID, name, price, available).Scan(&id)
	return id, err
}

func UpdateMenuItem(db *sql.DB, id int, name string, price int, available bool) (bool, error) {
	res, err := db.Exec(`
		UPDATE menu_items SET name=$1, price=$2, available=$3 WHERE id=$4
	`, name, price, available, id)
	if err != nil {
		return false, err
	}
	n, _ := res.RowsAffected()
	return n > 0, nil
}

func DeleteMenuItem(db *sql.DB, id int) (bool, error) {
	res, err := db.Exec(`DELETE FROM menu_items WHERE id=$1`, id)
	if err != nil {
		return false, err
	}
	n, _ := res.RowsAffected()
	return n > 0, nil
}

// helper used by orders
func GetMenuItemForOrder(db *sql.DB, id int) (name string, price int, err error) {
	err = db.QueryRow(`SELECT name, price FROM menu_items WHERE id=$1`, id).Scan(&name, &price)
	if err == sql.ErrNoRows {
		return "", 0, errors.New("menu item not found")
	}
	return
}

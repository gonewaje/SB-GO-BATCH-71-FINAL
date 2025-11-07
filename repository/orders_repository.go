package repository

import (
	"database/sql"
	"errors"

	"gonewaje/final/structs"
)

func CreateOrderWithItems(db *sql.DB, userID int, restaurantID int, items []structs.CreateOrderItem) (orderID int, total int, err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	err = tx.QueryRow(`
		INSERT INTO orders (user_id, restaurant_id, total_price, status)
		VALUES ($1,$2,0,'pending') RETURNING id
	`, userID, restaurantID).Scan(&orderID)
	if err != nil {
		return
	}

	total = 0
	for _, it := range items {
		name, priceEach, e := GetMenuItemForOrder(db, it.MenuItemID)
		if e != nil {
			err = e
			return
		}
		sub := priceEach * it.Quantity
		total += sub
		_, err = tx.Exec(`
			INSERT INTO order_items (order_id, menu_item_id, quantity, price_each)
			VALUES ($1,$2,$3,$4)
		`, orderID, it.MenuItemID, it.Quantity, priceEach)
		if err != nil {
			return
		}
		_ = name // unused in DB, but kept for clarity
	}

	_, err = tx.Exec(`UPDATE orders SET total_price=$1 WHERE id=$2`, total, orderID)
	if err != nil {
		return
	}

	err = tx.Commit()
	return
}

func ListOrdersByUser(db *sql.DB, userID int) ([]structs.Order, error) {
	rows, err := db.Query(`
		SELECT id, user_id, restaurant_id, total_price, status
		FROM orders WHERE user_id=$1 ORDER BY id DESC
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []structs.Order
	for rows.Next() {
		var o structs.Order
		if err := rows.Scan(&o.ID, &o.UserID, &o.RestaurantID, &o.TotalPrice, &o.Status); err != nil {
			return nil, err
		}
		items, err := listOrderItems(db, o.ID)
		if err != nil {
			return nil, err
		}
		o.Items = items
		orders = append(orders, o)
	}
	return orders, nil
}

func listOrderItems(db *sql.DB, orderID int) ([]structs.OrderItemPrint, error) {
	rows, err := db.Query(`
		SELECT oi.menu_item_id, m.name, oi.quantity, oi.price_each
		FROM order_items oi
		LEFT JOIN menu_items m ON m.id=oi.menu_item_id
		WHERE oi.order_id=$1
	`, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []structs.OrderItemPrint
	for rows.Next() {
		var row structs.OrderItemPrint
		if err := rows.Scan(&row.MenuItemID, &row.Name, &row.Quantity, &row.PriceEach); err != nil {
			return nil, err
		}
		row.Subtotal = row.PriceEach * row.Quantity
		res = append(res, row)
	}
	return res, nil
}

func GetOrderOwner(db *sql.DB, orderID int) (userID int, err error) {
	err = db.QueryRow(`SELECT user_id FROM orders WHERE id=$1`, orderID).Scan(&userID)
	if err == sql.ErrNoRows {
		return 0, errors.New("order not found")
	}
	return
}

func UpdateOrderStatus(db *sql.DB, orderID int, status string) error {
	_, err := db.Exec(`UPDATE orders SET status=$1 WHERE id=$2`, status, orderID)
	return err
}

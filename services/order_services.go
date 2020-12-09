package services

import (
	"fmt"

	"github.com/yeric17/inventory-system/db"
	"github.com/yeric17/inventory-system/models"
)

const fullOrderQuery = `SELECT orders.id, orders.client_id, products.name product_name, order_descriptions.name description, orders.quantity, products.price * orders.quantity as total
FROM orders
LEFT JOIN products ON orders.product_id = products.id
LEFT JOIN order_descriptions ON orders.order_description_id = order_descriptions.id`

var (
	OrderServices = orderServices{}
)

type orderServices struct{}

func (*orderServices) GetAll(limit int, page int) (data models.OrderData, err1 error) {
	var query string
	query = fmt.Sprintf(`%s LIMIT %d OFFSET %d`, fullOrderQuery, limit, page)

	stmt, err := db.DBClient.Prepare(query)

	if err != nil {
		return models.OrderData{}, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		return models.OrderData{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Order

		if err := rows.Scan(
			&order.ID,
			&order.ClientID,
			&order.ProductName,
			&order.Description,
			&order.Quantity,
			&order.Total,
		); err != nil {
			return models.OrderData{}, err
		}
		data.Orders = append(data.Orders, order)
	}

	return data, nil
}

func (*orderServices) Get(orderID uint64) (models.Order, error) {
	var query string
	query = fmt.Sprintf(`%s WHERE orders.id = $1`, fullOrderQuery)

	stmt, err := db.DBClient.Prepare(query)

	if err != nil {
		return models.Order{}, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(orderID)
	var order models.Order

	if err := row.Scan(
		&order.ID,
		&order.ClientID,
		&order.ProductName,
		&order.Description,
		&order.Quantity,
		&order.Total,
	); err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func (*orderServices) Create(order models.OrderCreate) (models.Order, error) {

	var query string

	query = "INSERT INTO orders(client_id, quantity, product_id, order_description_id) VALUES($1,$2,$3,$4) RETURNING id"

	stmt, err := db.DBClient.Prepare(query)

	if err != nil {
		return models.Order{}, err
	}

	defer stmt.Close()

	var lastID uint64

	err = stmt.QueryRow(order.ClientID, order.Quantity, order.ProductID, order.Description).Scan(&lastID)

	if err != nil {
		return models.Order{}, err
	}

	newOrder, err := OrderServices.Get(lastID)

	if err != nil {
		return models.Order{}, err
	}

	return newOrder, nil
}

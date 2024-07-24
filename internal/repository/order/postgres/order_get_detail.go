package postgres

import (
	"log"

	"github.com/ramailh/technical-test-dbo/internal/model"
)

var queryOrderGetDetail = `
	SELECT order_id, customer_id, order_date, status, created_at, updated_at
	FROM order 
	WHERE order_id = $1
`

func (repo *orderRepository) GetDetail(orderID int64) (*model.Order, error) {
	var res model.Order
	err := repo.db.QueryRow(queryOrderGetDetail, orderID).
		Scan(&res.OrderID, &res.CustomerID, &res.OrderDate, &res.Status, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &res, nil
}

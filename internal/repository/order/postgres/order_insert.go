package postgres

import (
	"errors"
	"log"

	"github.com/ramailh/technical-test-dbo/internal/model"
)

var queryInsertOrder = `
	INSERT INTO order (customer_id, order_date, item, status) VALUES ($1, $2, $3, $4)
`

func (repo *orderRepository) Insert(req *model.Order) error {
	res, err := repo.db.Exec(queryInsertOrder, req.CustomerID, req.OrderDate, req.Item, req.Status)
	if err != nil {
		log.Println(err)
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return err
	}

	if id == 0 {
		log.Println("failed to insert order")
		return errors.New("insert failed")
	}

	return nil
}

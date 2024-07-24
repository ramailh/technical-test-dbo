package postgres

import (
	"errors"
	"log"

	"github.com/ramailh/technical-test-dbo/internal/model"
)

var queryUpdateOrder = `
	UPDATE order 
	SET customer_id = $1,
		item = $2,
		status = $3,
	WHERE order_id = $4

`

func (repo *orderRepository) Update(req *model.Order) error {
	res, err := repo.db.Exec(queryUpdateOrder, req.CustomerID, req.Item, req.Status, req.OrderID)
	if err != nil {
		log.Println(err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}

	if rowsAffected == 0 {
		log.Println("failed to update order")
		return errors.New("failed to update order")
	}

	return nil
}

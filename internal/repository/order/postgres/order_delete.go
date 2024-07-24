package postgres

import (
	"errors"
	"log"
)

var queryDelete = `
	UPDATE order
	SET deleted_at = NOW()
	WHERE order_id = $1
`

func (repo *orderRepository) Delete(orderID int64) error {
	res, err := repo.db.Exec(queryDelete, orderID)
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
		log.Println("failed to delete order")
		return errors.New("failed to delete order")
	}

	return nil
}

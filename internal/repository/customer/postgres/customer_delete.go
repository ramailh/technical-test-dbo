package postgres

import (
	"errors"
	"log"
)

var queryDelete = `
	UPDATE customer 
	SET deleted_at = NOW()
	WHERE customer_id = $1
`

func (repo *customerRepository) Delete(customerID int64) error {
	res, err := repo.db.Exec(queryDelete, customerID)
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
		log.Println("failed to update customer")
		return errors.New("failed to update customer")
	}

	return nil
}

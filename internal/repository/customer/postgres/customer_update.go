package postgres

import (
	"errors"
	"log"

	"github.com/ramailh/technical-test-dbo/internal/model"
)

var queryUpdateCustomer = `
	UPDATE customer 
	SET first_name = $1,
		last_name = $2,
		email = $3,
		phone = $4
	WHERE customer_id = $5

`

func (repo *customerRepository) Update(req *model.Customer) error {
	res, err := repo.db.Exec(queryUpdateCustomer, req.FirstName, req.LastName, req.Email, req.Phone, req.CustomerID)
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

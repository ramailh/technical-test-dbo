package postgres

import (
	"log"

	"github.com/ramailh/technical-test-dbo/internal/model"
)

var queryCustomerGetDetail = `
	SELECT customer_id, first_name, last_name, email, phone, created_at, updated_at
	FROM customer 
	WHERE customer_id = $1
`

func (repo *customerRepository) GetDetail(customerID int64) (*model.Customer, error) {
	var res model.Customer
	err := repo.db.QueryRow(queryCustomerGetDetail, customerID).
		Scan(&res.CustomerID, &res.FirstName, &res.LastName, &res.Email, &res.Phone, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &res, nil
}

package postgres

import (
	"log"

	"github.com/ramailh/technical-test-dbo/internal/model"
)

var queryInsertCustomer = `
	INSERT INTO customer (first_name, last_name, email, phone) VALUES ($1, $2, $3, $4)
`

func (repo *customerRepository) Insert(req *model.Customer) (int64, error) {
	res, err := repo.db.Exec(queryInsertCustomer, req.FirstName, req.LastName, req.Email, req.Phone)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return id, nil
}

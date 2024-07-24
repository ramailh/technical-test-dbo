package postgres

import (
	"database/sql"

	"github.com/ramailh/technical-test-dbo/internal/repository/customer"
)

type customerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) customer.CustomerRepository {
	return &customerRepository{
		db,
	}
}

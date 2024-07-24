package postgres

import (
	"database/sql"

	"github.com/ramailh/technical-test-dbo/internal/repository/order"
)

type orderRepository struct {
	db *sql.DB
}

func NeworderRepository(db *sql.DB) order.OrderRepository {
	return &orderRepository{
		db,
	}
}

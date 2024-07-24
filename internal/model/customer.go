package model

import (
	"time"
)

type Customer struct {
	CustomerID int64     `json:"customer_id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

type GetCustomerListResponse struct {
	Data []*Customer `json:"data"`
	Meta *Meta       `json:"meta"`
}

type RegisterCustomerRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

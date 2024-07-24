package model

import "time"

type Order struct {
	OrderID    int64     `json:"order_id"`
	CustomerID int64     `json:"customer_id"`
	OrderDate  time.Time `json:"order_date"`
	Item       string    `json:"item"`
	Status     int       `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

type GetOrderListResponse struct {
	Data []*Order `json:"data"`
	Meta *Meta    `json:"meta"`
}

type InsertOrderRequest struct {
	OrderDate  time.Time `json:"order_date"`
	Status     int       `json:"status"`
	CustomerID int64     `json:"customer_id"`
	Item       string    `json:"item"`
}

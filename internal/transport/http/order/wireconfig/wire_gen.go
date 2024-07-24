// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wireconfig

import (
	"database/sql"
	"github.com/ramailh/technical-test-dbo/internal/common/config/redis"
	"github.com/ramailh/technical-test-dbo/internal/repository/order/postgres"
	"github.com/ramailh/technical-test-dbo/internal/service/order"
	"github.com/ramailh/technical-test-dbo/internal/transport/http/order/controller"
)

// Injectors from order_wire.go:

func InitializeOrderController(pg *sql.DB, rds redis.Redis) (controller.OrderController, error) {
	orderRepository := postgres.NeworderRepository(pg)
	orderService := order.NewOrderService(orderRepository, rds)
	orderController := controller.NewOrderController(orderService)
	return orderController, nil
}
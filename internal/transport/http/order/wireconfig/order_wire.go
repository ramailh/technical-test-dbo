//go:build wireinject
// +build wireinject

package wireconfig

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/ramailh/technical-test-dbo/internal/common/config/redis"
	pgOrder "github.com/ramailh/technical-test-dbo/internal/repository/order/postgres"
	svcOrder "github.com/ramailh/technical-test-dbo/internal/service/order"
	orderCtrl "github.com/ramailh/technical-test-dbo/internal/transport/http/order/controller"
)

func InitializeOrderController(pg *sql.DB, rds redis.Redis) (orderCtrl.OrderController, error) {
	wire.Build(
		pgOrder.NeworderRepository,
		svcOrder.NewOrderService,
		orderCtrl.NewOrderController,
	)
	return nil, nil
}

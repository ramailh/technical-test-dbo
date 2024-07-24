//go:build wireinject
// +build wireinject

package wireconfig

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/ramailh/technical-test-dbo/internal/common/config/redis"
	pgAuth "github.com/ramailh/technical-test-dbo/internal/repository/auth/postgres"
	pgCust "github.com/ramailh/technical-test-dbo/internal/repository/customer/postgres"
	svcAuth "github.com/ramailh/technical-test-dbo/internal/service/auth"
	svcCust "github.com/ramailh/technical-test-dbo/internal/service/customer"
	customerCtrl "github.com/ramailh/technical-test-dbo/internal/transport/http/customer/controller"
)

func InitializeCustomerController(pg *sql.DB, rds redis.Redis) (customerCtrl.CustomerController, error) {
	wire.Build(
		pgAuth.NewAuthRepository, pgCust.NewCustomerRepository,
		svcAuth.NewAuthService, svcCust.NewCustomerService,
		customerCtrl.NewCustomerController,
	)
	return nil, nil
}

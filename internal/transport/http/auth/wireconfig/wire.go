//go:build wireinject
// +build wireinject

package wireconfig

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/ramailh/technical-test-dbo/internal/common/config/redis"
	pgAuth "github.com/ramailh/technical-test-dbo/internal/repository/auth/postgres"
	svcAuth "github.com/ramailh/technical-test-dbo/internal/service/auth"
	authCtrl "github.com/ramailh/technical-test-dbo/internal/transport/http/auth/controller"
)

func InitializeAuthController(pg *sql.DB, rds redis.Redis) (authCtrl.AuthController, error) {
	wire.Build(
		pgAuth.NewAuthRepository,
		svcAuth.NewAuthService,
		authCtrl.NewAuthController,
	)
	return nil, nil
}

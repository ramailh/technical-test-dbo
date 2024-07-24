package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/ramailh/technical-test-dbo/internal/common/config/postgres"
	"github.com/ramailh/technical-test-dbo/internal/common/config/redis"
	"github.com/ramailh/technical-test-dbo/internal/common/env"
	"github.com/ramailh/technical-test-dbo/internal/common/middleware"
	authRouter "github.com/ramailh/technical-test-dbo/internal/transport/http/auth/router"
	customerRouter "github.com/ramailh/technical-test-dbo/internal/transport/http/customer/router"
	orderRouter "github.com/ramailh/technical-test-dbo/internal/transport/http/order/router"

	wireAuth "github.com/ramailh/technical-test-dbo/internal/transport/http/auth/wireconfig"
	wireCust "github.com/ramailh/technical-test-dbo/internal/transport/http/customer/wireconfig"
	wireOrder "github.com/ramailh/technical-test-dbo/internal/transport/http/order/wireconfig"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	env.LoadEnv()
}

func main() {
	rds, err := redis.NewRedisConn(
		int64(env.RedisMaxIdle),
		int64(env.RedisMaxActive),
		int64(env.RedisDB),
		env.RedisHost,
		env.RedisPort,
		env.RedisPassword,
	)
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.NewSqlConnection(postgres.DBConfiguration{
		DBHost:            env.DBHost,
		DBPort:            env.DBPort,
		DBUser:            env.DBUser,
		DBPassword:        env.DBPassword,
		DBName:            env.DBName,
		MaxConnection:     env.DBMaxConn,
		MaxIdleConnection: env.DBMaxIdleConn,
	})
	if err != nil {
		log.Fatal(err)
	}

	custController, err := wireCust.InitializeCustomerController(db, rds)
	if err != nil {
		log.Fatal(err)
	}

	authController, err := wireAuth.InitializeAuthController(db, rds)
	if err != nil {
		log.Fatal(err)
	}

	orderController, err := wireOrder.InitializeOrderController(db, rds)
	if err != nil {
		log.Fatal(err)
	}

	md := middleware.NewMiddlewareAuth(rds)

	r := gin.Default()
	authRouter.AuthRouter(r, md, authController)
	orderRouter.OrderRouter(r, md, orderController)
	customerRouter.CustomerRouter(r, md, custController)

	go func() {
		if err := r.Run(":" + env.Port); err != nil {
			log.Fatal(err)
		}
	}()

	done := make(chan bool, 1)

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

		sig := <-sigCh
		log.Printf("Receive signal %v. Shutting down gracefully...\n", sig)

		done <- true
	}()

	<-done
}

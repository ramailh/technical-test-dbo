package order

import (
	"github.com/ramailh/technical-test-dbo/internal/common/config/redis"
	"github.com/ramailh/technical-test-dbo/internal/model"
	"github.com/ramailh/technical-test-dbo/internal/repository/order"
)

type OrderService interface {
	GetOrderList(req *model.PaginationRequest) (*model.GetOrderListResponse, error)
	GetOrderDetail(id int64) (*model.Order, error)
	InsertOrder(req *model.InsertOrderRequest) error
	UpdateOrder(req *model.Order) error
	DeleteOrder(id int64) error
}

type orderService struct {
	repo order.OrderRepository
	rds  redis.Redis
}

func NewOrderService(repo order.OrderRepository, rds redis.Redis) OrderService {
	return &orderService{
		repo: repo,
		rds:  rds,
	}
}

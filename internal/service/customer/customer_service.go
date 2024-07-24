package customer

import (
	"github.com/ramailh/technical-test-dbo/internal/common/config/redis"
	"github.com/ramailh/technical-test-dbo/internal/model"
	"github.com/ramailh/technical-test-dbo/internal/repository/customer"
	"github.com/ramailh/technical-test-dbo/internal/service/auth"
)

type CustomerService interface {
	GetCustomerList(req *model.PaginationRequest) (*model.GetCustomerListResponse, error)
	GetCustomerDetail(id int64) (*model.Customer, error)
	RegisterCustomer(req *model.RegisterCustomerRequest) error
	UpdateCustomer(req *model.Customer) error
	DeleteCustomer(id int64) error
}

type customerService struct {
	repo    customer.CustomerRepository
	svcAuth auth.CustAuthService
	rds     redis.Redis
}

func NewCustomerService(repo customer.CustomerRepository, svcAuth auth.AuthService, rds redis.Redis) CustomerService {
	return &customerService{
		repo, svcAuth, rds,
	}
}

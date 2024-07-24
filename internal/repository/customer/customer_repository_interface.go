package customer

import "github.com/ramailh/technical-test-dbo/internal/model"

type CustomerRepository interface {
	GetDetail(customerID int64) (*model.Customer, error)
	GetList(req *model.PaginationRequest) ([]*model.Customer, *model.Meta, error)
	Insert(req *model.Customer) (int64, error)
	Update(req *model.Customer) error
	Delete(customerID int64) error
}
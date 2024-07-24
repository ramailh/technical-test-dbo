package order

import "github.com/ramailh/technical-test-dbo/internal/model"

type OrderRepository interface {
	GetDetail(orderID int64) (*model.Order, error)
	GetList(req *model.PaginationRequest) ([]*model.Order, *model.Meta, error)
	Insert(req *model.Order) error
	Update(req *model.Order) error
	Delete(orderID int64) error
}

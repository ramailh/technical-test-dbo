package order

import (
	"log"

	"github.com/ramailh/technical-test-dbo/internal/model"
)

func (svc *orderService) GetOrderDetail(id int64) (*model.Order, error) {
	res, err := svc.repo.GetDetail(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return res, nil
}

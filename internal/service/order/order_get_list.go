package order

import (
	"log"

	"github.com/ramailh/technical-test-dbo/internal/model"
)

func (svc *orderService) GetOrderList(req *model.PaginationRequest) (*model.GetOrderListResponse, error) {
	res, meta, err := svc.repo.GetList(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &model.GetOrderListResponse{
		Data: res,
		Meta: meta,
	}, nil
}

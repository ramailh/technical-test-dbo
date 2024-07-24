package customer

import (
	"log"

	"github.com/ramailh/technical-test-dbo/internal/model"
)

func (svc *customerService) GetCustomerList(req *model.PaginationRequest) (*model.GetCustomerListResponse, error) {
	res, meta, err := svc.repo.GetList(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &model.GetCustomerListResponse{
		Data: res,
		Meta: meta,
	}, nil
}

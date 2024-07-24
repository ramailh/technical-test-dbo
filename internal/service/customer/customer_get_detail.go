package customer

import (
	"log"

	"github.com/ramailh/technical-test-dbo/internal/model"
)

func (svc *customerService) GetCustomerDetail(id int64) (*model.Customer, error) {
	res, err := svc.repo.GetDetail(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}

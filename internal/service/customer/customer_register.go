package customer

import (
	"log"

	"github.com/ramailh/technical-test-dbo/internal/model"
)

func (svc *customerService) RegisterCustomer(req *model.RegisterCustomerRequest) error {
	custID, err := svc.repo.Insert(&model.Customer{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
	})
	if err != nil {
		log.Println(err)
		return err
	}

	err = svc.svcAuth.Register(&model.RegisterRequest{
		CustomerID: custID,
		Username:   req.Username,
		Password:   req.Password,
	})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

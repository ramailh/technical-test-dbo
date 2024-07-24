package order

import (
	"log"

	"github.com/ramailh/technical-test-dbo/internal/model"
)

func (svc *orderService) UpdateOrder(req *model.Order) error {
	if err := svc.repo.Update(req); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

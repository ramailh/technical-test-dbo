package order

import (
	"log"
	"time"

	"github.com/ramailh/technical-test-dbo/internal/model"
)

func (svc *orderService) InsertOrder(req *model.InsertOrderRequest) error {
	err := svc.repo.Insert(&model.Order{
		CustomerID: req.CustomerID,
		OrderDate:  time.Now(),
		Item:       req.Item,
		Status:     req.Status,
	})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

package customer

import (
	"database/sql"
	"errors"
	"log"

	"github.com/ramailh/technical-test-dbo/internal/model"
)

func (svc *customerService) UpdateCustomer(req *model.Customer) error {
	user, err := svc.repo.GetDetail(req.CustomerID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println(err)
		return err
	}

	if user.CustomerID == 0 {
		err = errors.New("error customer not exist")
		log.Println(err)
		return err
	}

	if err = svc.repo.Update(req); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

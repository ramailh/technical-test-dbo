package customer

import (
	"database/sql"
	"errors"
	"log"
)

func (svc *customerService) DeleteCustomer(id int64) error {
	user, err := svc.repo.GetDetail(id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Println(err)
		return err
	}

	if user.CustomerID == 0 {
		err = errors.New("error customer not exist")
		log.Println(err)
		return err
	}

	if err = svc.repo.Delete(id); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

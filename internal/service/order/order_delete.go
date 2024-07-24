package order

import "log"

func (svc *orderService) DeleteOrder(id int64) error {
	if err := svc.repo.Delete(id); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

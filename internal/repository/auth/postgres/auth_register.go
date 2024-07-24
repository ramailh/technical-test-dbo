package postgres

import (
	"errors"
	"log"

	"github.com/ramailh/technical-test-dbo/internal/model"
)

var queryRegister = `
	INSERT INTO auth (customer_id, username, password_hash) VALUES ($1, $2, $3)
`

func (repo *authRepository) Register(req *model.Auth) error {
	res, err := repo.db.Exec(queryRegister, req.CustomerID, req.Username, req.PasswordHash)
	if err != nil {
		log.Println(err)
		return err
	}

	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		log.Println()
		return err
	}

	if lastInsertedID == 0 {
		log.Println("failed to insert")
		return errors.New("failed to insert")
	}

	return nil
}

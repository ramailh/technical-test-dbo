package postgres

import (
	"log"

	"github.com/ramailh/technical-test-dbo/internal/model"
)

var queryInsertLoginSession = `
		INSERT INTO session (auth_id, customer_id, expired_at) 
		VALUES ($1, $2, $3)
	`

func (repo *authRepository) InsertLoginSession(req *model.Session) (int64, error) {
	res, err := repo.db.Exec(queryInsertLoginSession, req.AuthID, req.CustomerID, req.ExpiredAt)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return id, nil
}

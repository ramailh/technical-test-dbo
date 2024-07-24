package postgres

import (
	"log"

	"github.com/ramailh/technical-test-dbo/internal/model"
)

var queryGetLoginSessionByID = `
	SELECT session_id, auth_id, customer_id, expired_at, created_at
	FROM session 
	WHERE session_id = $1
`

func (repo *authRepository) GetLoginSessionByID(sessionID int64) (*model.Session, error) {
	var res model.Session
	err := repo.db.QueryRow(queryGetLoginSessionByID, sessionID).Scan(&res.SessionID, &res.AuthID,
		&res.CustomerID, &res.ExpiredAt, &res.CreatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &res, nil
}

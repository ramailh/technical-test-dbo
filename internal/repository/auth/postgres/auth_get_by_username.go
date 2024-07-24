package postgres

import (
	"log"

	"github.com/ramailh/technical-test-dbo/internal/model"
)

var queryAuthGetByUsername = `
	SELECT auth_id, customer_id, username, password_hash, created_at, updated_at
	FROM auth 
	WHERE customer_id = $1
`

func (repo *authRepository) GetAuthByUsername(username string) (*model.Auth, error) {
	var res model.Auth
	err := repo.db.QueryRow(queryAuthGetByUsername, username).Scan(&res.AuthID, &res.CustomerID, &res.Username, &res.PasswordHash, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &res, nil
}

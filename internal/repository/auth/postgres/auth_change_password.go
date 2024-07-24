package postgres

import (
	"errors"
	"log"
)

var queryChangePassword = `
	UPDATE session
	SET password_hash = $2
	WHERE customer_id = $1
`

func (repo *authRepository) ChangePassword(username string, password string) error {
	res, err := repo.db.Exec(queryChangePassword, username, password)
	if err != nil {
		log.Println(err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}

	if rowsAffected == 0 {
		log.Println("failed to change password")
		return errors.New("failed to change password")
	}

	return nil
}

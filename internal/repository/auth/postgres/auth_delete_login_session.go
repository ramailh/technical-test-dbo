package postgres

import (
	"errors"
	"log"
)

var queryDeleteLoginSession = `
	DELETE FROM session 
	WHERE session_id = $1
`

func (repo *authRepository) DeleteLoginSession(sessionID int64) error {
	res, err := repo.db.Exec(queryDeleteLoginSession, sessionID)
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
		log.Println("failed to delete session")
		return errors.New("failed to delete session")
	}

	return nil
}

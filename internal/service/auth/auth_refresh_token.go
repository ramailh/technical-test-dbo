package auth

import (
	"errors"
	"log"
	"time"

	"github.com/ramailh/technical-test-dbo/internal/common/token"
	"github.com/ramailh/technical-test-dbo/internal/model"
)

func (svc *authService) RefreshToken(req *model.RefreshTokenRequest) (*model.Session, error) {
	session, err := svc.repo.GetLoginSessionByID(req.SessionID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if time.Now().Unix() > session.ExpiredAt.Unix() {
		err = errors.New("token expired, please login again")
		log.Println(err)
		return nil, err
	}

	session.ExpiredAt = time.Now().Add(24 * time.Hour)

	sessionID, err := svc.repo.InsertLoginSession(session)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	jwtToken, err := token.GenerateJWTCust(sessionID, session.CustomerID, session.AuthID, session.ExpiredAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	session.Token = jwtToken

	return session, nil
}

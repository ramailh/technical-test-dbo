package auth

import (
	"log"
	"time"

	"github.com/ramailh/technical-test-dbo/internal/common/token"
	"github.com/ramailh/technical-test-dbo/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func (svc *authService) Login(req *model.LoginRequest) (*model.Session, error) {
	auth, err := svc.repo.GetAuthByUsername(req.Username)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(auth.PasswordHash), []byte(req.Password)); err != nil {
		log.Println(err)
		return nil, err
	}

	session := &model.Session{
		AuthID:     auth.AuthID,
		CustomerID: auth.CustomerID,
		ExpiredAt:  time.Now().Add(24 * time.Hour),
	}

	sessionID, err := svc.repo.InsertLoginSession(session)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	jwtToken, err := token.GenerateJWTCust(sessionID, auth.CustomerID, auth.AuthID, session.ExpiredAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	session.Token = jwtToken

	return session, nil
}

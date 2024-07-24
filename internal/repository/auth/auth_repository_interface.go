package auth

import "github.com/ramailh/technical-test-dbo/internal/model"

type AuthRepository interface {
	Register(req *model.Auth) error
	GetAuthByUsername(username string) (*model.Auth, error)
	InsertLoginSession(req *model.Session) (int64, error)
	GetLoginSessionByID(sessionID int64) (*model.Session, error)
	DeleteLoginSession(sessionID int64) error
	ChangePassword(username string, password string) error
}

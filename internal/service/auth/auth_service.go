package auth

import (
	"github.com/ramailh/technical-test-dbo/internal/common/config/redis"
	"github.com/ramailh/technical-test-dbo/internal/model"
	repository "github.com/ramailh/technical-test-dbo/internal/repository/auth"
)

type (
	CustAuthService interface {
		Register(req *model.RegisterRequest) error
		ChangePassword(req *model.ChangePasswordRequest) error
	}

	AuthService interface {
		Login(req *model.LoginRequest) (*model.Session, error)
		Logout(sessionID int64) error
		RefreshToken(req *model.RefreshTokenRequest) (*model.Session, error)
		CustAuthService
	}
)

type authService struct {
	repo repository.AuthRepository
	rds  redis.Redis
}

func NewAuthService(repo repository.AuthRepository, rds redis.Redis) AuthService {
	return &authService{
		repo, rds,
	}
}

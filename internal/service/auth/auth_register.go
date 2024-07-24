package auth

import (
	"log"

	"github.com/ramailh/technical-test-dbo/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func (svc *authService) Register(req *model.RegisterRequest) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		log.Println(err)
		return err
	}

	err = svc.repo.Register(&model.Auth{
		CustomerID:   req.CustomerID,
		Username:     req.Username,
		PasswordHash: string(pass),
	})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

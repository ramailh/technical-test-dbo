package auth

import (
	"log"

	"github.com/ramailh/technical-test-dbo/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func (svc *authService) ChangePassword(req *model.ChangePasswordRequest) error {
	auth, err := svc.repo.GetAuthByUsername(req.Username)
	if err != nil {
		log.Println(err)
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(auth.PasswordHash), []byte(req.OldPassword)); err != nil {
		return err
	}

	newPass, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), 10)
	if err != nil {
		log.Println(err)
		return err
	}

	err = svc.repo.ChangePassword(req.Username, string(newPass))
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

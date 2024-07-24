package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ramailh/technical-test-dbo/internal/service/auth"
)

type AuthController interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
	RefreshToken(c *gin.Context)
}

type authController struct {
	svc auth.AuthService
}

func NewAuthController(svc auth.AuthService) AuthController {
	return &authController{
		svc: svc,
	}
}

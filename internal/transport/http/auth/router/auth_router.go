package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ramailh/technical-test-dbo/internal/common/middleware"
	"github.com/ramailh/technical-test-dbo/internal/transport/http/auth/controller"
)

func AuthRouter(r *gin.Engine, md middleware.AuthMiddleware, c controller.AuthController) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", c.Login)
		authGroup.Use(md.AuthJWT)
		authGroup.POST("/logout", c.Logout)
		authGroup.POST("/refresh-token", c.RefreshToken)
	}
}

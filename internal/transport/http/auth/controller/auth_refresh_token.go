package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ramailh/technical-test-dbo/internal/common/response"
	"github.com/ramailh/technical-test-dbo/internal/model"
)

func (ctr *authController) RefreshToken(c *gin.Context) {
	sessionID := c.GetInt64("session_id")
	custID := c.GetInt64("customer_id")
	authID := c.GetInt64("auth_id")

	res, err := ctr.svc.RefreshToken(&model.RefreshTokenRequest{
		AuthID:     authID,
		CustomerID: custID,
		SessionID:  sessionID,
	})
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, response.ErrInternal)
		return
	}

	c.JSON(http.StatusOK, res)
}

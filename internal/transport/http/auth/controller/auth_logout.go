package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ramailh/technical-test-dbo/internal/common/response"
)

func (ctr *authController) Logout(c *gin.Context) {
	sessionID := c.GetInt64("session_id")

	if sessionID == 0 {
		log.Println("failed to get session id from token")
		c.JSON(http.StatusBadRequest, response.ErrGetSessionID)
		return
	}

	err := ctr.svc.Logout(sessionID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, response.ResponseErr{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, response.CommonSuccess{
		Message: "success",
		Success: true,
	})
}

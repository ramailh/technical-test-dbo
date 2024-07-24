package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ramailh/technical-test-dbo/internal/common/response"
	"github.com/ramailh/technical-test-dbo/internal/model"
)

func (ctr *authController) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.ErrParsingJSON)
		return
	}

	res, err := ctr.svc.Login(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, response.ResponseErr{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

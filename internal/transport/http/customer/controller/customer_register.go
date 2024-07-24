package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ramailh/technical-test-dbo/internal/common/response"
	"github.com/ramailh/technical-test-dbo/internal/model"
)

func (ctr *customerController) RegisterCustomer(c *gin.Context) {
	var req model.RegisterCustomerRequest
	if err := c.BindJSON(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.ErrParsingJSON)
		return
	}

	if err := ctr.svc.RegisterCustomer(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, response.ErrInternal)
		return
	}

	c.JSON(http.StatusOK, response.CommonSuccess{
		Message: "success",
		Success: true,
	})
}

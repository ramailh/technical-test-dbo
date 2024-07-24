package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ramailh/technical-test-dbo/internal/common/response"
	"github.com/ramailh/technical-test-dbo/internal/model"
)

func (ctr *customerController) GetCustomerList(c *gin.Context) {
	var req model.PaginationRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.ErrParsingQueryString)
		return
	}

	res, err := ctr.svc.GetCustomerList(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, response.ErrInternal)
		return
	}

	c.JSON(http.StatusOK, res)
}

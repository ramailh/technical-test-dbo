package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ramailh/technical-test-dbo/internal/common/response"
)

func (ctr *customerController) GetCustomerDetail(c *gin.Context) {
	customerID := c.Param("customer_id")
	custID, err := strconv.Atoi(customerID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.ErrParsingURLParam)
		return
	}

	res, err := ctr.svc.GetCustomerDetail(int64(custID))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, response.ErrInternal)
		return
	}

	c.JSON(http.StatusOK, res)
}

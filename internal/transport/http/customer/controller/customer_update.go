package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ramailh/technical-test-dbo/internal/common/response"
	"github.com/ramailh/technical-test-dbo/internal/model"
)

func (ctr *customerController) UpdateCustomer(c *gin.Context) {
	var req model.Customer
	if err := c.BindJSON(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.ErrParsingJSON)
		return
	}

	customerID := c.Param("customer_id")
	custID, err := strconv.Atoi(customerID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.ErrParsingURLParam)
		return
	}

	req.CustomerID = int64(custID)

	if err := ctr.svc.UpdateCustomer(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, response.ErrInternal)
		return
	}

	c.JSON(http.StatusOK, response.CommonSuccess{
		Message: "success",
		Success: true,
	})
}

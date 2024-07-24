package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ramailh/technical-test-dbo/internal/common/response"
)

func (ctr *customerController) DeleteCustomer(c *gin.Context) {
	custIDParam := c.Param("customer_id")
	custID, err := strconv.Atoi(custIDParam)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.ErrParsingURLParam)
		return
	}

	if err := ctr.svc.DeleteCustomer(int64(custID)); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, response.ErrInternal)
		return
	}

	c.JSON(http.StatusOK, response.CommonSuccess{
		Message: "success",
		Success: true,
	})
}

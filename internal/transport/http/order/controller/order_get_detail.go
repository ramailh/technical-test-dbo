package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ramailh/technical-test-dbo/internal/common/response"
)

func (ctr *orderController) GetOrderDetail(c *gin.Context) {
	orderID := c.Param("order_id")
	orderIDInt, err := strconv.Atoi(orderID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.ErrParsingURLParam)
		return
	}

	res, err := ctr.svc.GetOrderDetail(int64(orderIDInt))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, response.ErrInternal)
		return
	}

	c.JSON(http.StatusOK, res)
}
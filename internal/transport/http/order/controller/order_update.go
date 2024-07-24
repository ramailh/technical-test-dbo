package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ramailh/technical-test-dbo/internal/common/response"
	"github.com/ramailh/technical-test-dbo/internal/model"
)

func (ctr *orderController) UpdateOrder(c *gin.Context) {
	var req model.Order
	if err := c.BindJSON(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.ErrParsingJSON)
		return
	}

	orderIDParam := c.Param("order_id")
	orderID, err := strconv.Atoi(orderIDParam)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.ErrParsingURLParam)
		return
	}

	req.OrderID = int64(orderID)

	if err := ctr.svc.UpdateOrder(&req); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, response.ErrInternal)
		return
	}

	c.JSON(http.StatusOK, response.CommonSuccess{
		Message: "success",
		Success: true,
	})
}

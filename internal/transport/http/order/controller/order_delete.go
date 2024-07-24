package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ramailh/technical-test-dbo/internal/common/response"
)

func (ctr *orderController) DeleteOrder(c *gin.Context) {
	orderIDParam := c.Param("order_id")
	orderID, err := strconv.Atoi(orderIDParam)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.ErrParsingURLParam)
		return
	}

	if err := ctr.svc.DeleteOrder(int64(orderID)); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, response.ErrInternal)
		return
	}

	c.JSON(http.StatusOK, response.CommonSuccess{
		Message: "success",
		Success: true,
	})
}

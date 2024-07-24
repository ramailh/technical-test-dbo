package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ramailh/technical-test-dbo/internal/service/order"
)

type OrderController interface {
	GetOrderList(c *gin.Context)
	GetOrderDetail(c *gin.Context)
	InsertOrder(c *gin.Context)
	UpdateOrder(c *gin.Context)
	DeleteOrder(c *gin.Context)
}

type orderController struct {
	svc order.OrderService
}

func NewOrderController(svc order.OrderService) OrderController {
	return &orderController{
		svc,
	}
}

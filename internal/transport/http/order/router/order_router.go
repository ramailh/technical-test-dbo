package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ramailh/technical-test-dbo/internal/common/middleware"
	"github.com/ramailh/technical-test-dbo/internal/transport/http/order/controller"
)

func OrderRouter(r *gin.Engine, md middleware.AuthMiddleware, c controller.OrderController) {
	orderGroup := r.Group("/order", md.AuthJWT)
	{
		orderGroup.GET("/:order_id", c.GetOrderDetail)
		orderGroup.GET("/", c.GetOrderList)
		orderGroup.POST("/", c.InsertOrder)
		orderGroup.PUT("/:order_id", c.UpdateOrder)
		orderGroup.DELETE("/:order_id", c.DeleteOrder)
	}
}

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ramailh/technical-test-dbo/internal/common/middleware"
	"github.com/ramailh/technical-test-dbo/internal/transport/http/customer/controller"
)

func CustomerRouter(r *gin.Engine, md middleware.AuthMiddleware, c controller.CustomerController) {
	custGroup := r.Group("/customer")
	{
		custGroup.POST("/", c.RegisterCustomer)
		custGroup.Use(md.AuthJWT)
		custGroup.GET("/:customer_id", c.GetCustomerDetail)
		custGroup.GET("/", c.GetCustomerList)
		custGroup.PUT("/:customer_id", c.UpdateCustomer)
		custGroup.DELETE("/:customer_id", c.DeleteCustomer)
	}
}

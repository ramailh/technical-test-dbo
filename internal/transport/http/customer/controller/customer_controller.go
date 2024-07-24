package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ramailh/technical-test-dbo/internal/service/customer"
)

type CustomerController interface {
	GetCustomerList(c *gin.Context)
	GetCustomerDetail(c *gin.Context)
	RegisterCustomer(c *gin.Context)
	UpdateCustomer(c *gin.Context)
	DeleteCustomer(c *gin.Context)
}

type customerController struct {
	svc customer.CustomerService
}

func NewCustomerController(svc customer.CustomerService) CustomerController {
	return &customerController{
		svc: svc,
	}
}

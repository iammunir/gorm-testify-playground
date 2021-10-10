package handler

import (
	"learn-gorm/controller"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetCustomerByName(*gin.Context)
	GetCustomerById(*gin.Context)
	GetCustomerByWithAccount(*gin.Context)
	GetCustomerByWithAccountGoroutine(*gin.Context)
	GetContractDetails(*gin.Context)
	GetCustomerByNameMongo(*gin.Context)
}

type handler struct {
	contoller controller.Controller
}

func NewHandler(ctrl controller.Controller) Handler {
	return &handler{
		contoller: ctrl,
	}
}

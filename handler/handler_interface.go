package handler

import (
	"learn-gorm/controller"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetUserByName(*gin.Context)
	GetUserById(*gin.Context)
}

type handler struct {
	contoller controller.Controller
}

func NewHandler(ctrl controller.Controller) Handler {
	return &handler{
		contoller: ctrl,
	}
}

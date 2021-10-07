package app

import (
	"learn-gorm/controller"
	"learn-gorm/handler"
	"learn-gorm/repository"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitRouter(mysqlConn *gorm.DB) *gin.Engine {

	repo := repository.GetCustomerRepository(mysqlConn)
	controller := controller.NewController(repo)
	handler := handler.NewHandler(controller)

	router := gin.Default()

	router.POST("/get-customer-by-name", handler.GetUserByName)

	router.POST("/get-customer-by-id", handler.GetUserById)

	return router
}

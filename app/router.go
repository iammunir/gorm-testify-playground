package app

import (
	"learn-gorm/controller"
	"learn-gorm/handler"
	"learn-gorm/repository"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitRouter(mysqlConn *gorm.DB) *gin.Engine {

	custRepo := repository.GetCustomerRepository(mysqlConn)
	accRepo := repository.GetAccountRepository(mysqlConn)
	contRepo := repository.GetContractRepository(mysqlConn)

	controller := controller.NewController(custRepo, accRepo, contRepo)
	handler := handler.NewHandler(controller)

	router := gin.Default()

	router.POST("/get-customer-by-name", handler.GetCustomerByName)

	router.POST("/get-customer-by-id", handler.GetCustomerById)

	router.POST("/get-customer-with-account", handler.GetCustomerByWithAccount)

	router.POST("/get-customer-with-account-goroutine", handler.GetCustomerByWithAccountGoroutine)

	router.POST("/get-contract-details", handler.GetContractDetails)

	return router
}

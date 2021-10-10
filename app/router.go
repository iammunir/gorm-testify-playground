package app

import (
	"learn-gorm/controller"
	"learn-gorm/handler"
	"learn-gorm/repository"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRouter(mysqlConn *gorm.DB, mongoConn *mongo.Database) *gin.Engine {

	custRepo := repository.GetCustomerRepository(mysqlConn)
	accRepo := repository.GetAccountRepository(mysqlConn)
	contRepo := repository.GetContractRepository(mysqlConn)
	custMongoRepo := repository.GetCustomerMongoRepository(mongoConn)
	custMysqlRepo := repository.GetCustomerMysqlRepository(mysqlConn)

	controller := controller.NewController(custRepo, accRepo, contRepo, custMongoRepo, custMysqlRepo)
	handler := handler.NewHandler(controller)

	router := gin.Default()

	router.POST("/get-customer-by-name", handler.GetCustomerByName)

	router.POST("/get-customer-by-id", handler.GetCustomerById)

	router.POST("/get-customer-with-account", handler.GetCustomerByWithAccount)

	router.POST("/get-customer-with-account-goroutine", handler.GetCustomerByWithAccountGoroutine)

	router.POST("/get-contract-details", handler.GetContractDetails)

	router.POST("/get-customer-by-name-mongo", handler.GetCustomerByNameMongo)

	return router
}

package controller

import (
	"learn-gorm/models"
	"learn-gorm/repository"
)

type Controller interface {
	GetCustomerByName(models.CustomerRequest) ([]models.Customer, *models.ErrorResponse)
	GetCustomerById(models.CustomerRequest) (*models.Customer, *models.ErrorResponse)

	// comparison between normal and concurent
	GetCustomerWithAccount(models.CustomerRequest) ([]models.Customer, *models.ErrorResponse)
	GetCustomerWithAccountGoroutine(models.CustomerRequest) ([]models.Customer, *models.ErrorResponse)

	// multiple primarykeys process
	GetContractDetails(models.ContractRequest) ([]models.ContractDetails, *models.ErrorResponse)

	// mapping with mongodb
	GetCustomerByNameMongo(models.ContractRequest) ([]models.CustomerMongo, *models.ErrorResponse)
	GetCustomerByNameMongoWithSwitchObj(models.ContractRequest) ([]models.CustomerMongo, *models.ErrorResponse)
}

type controller struct {
	CustomerRepo repository.CustomerRepository
	AccountRepo  repository.AccountRepository
	ContractRepo repository.ContractRepository
	CustMongo    repository.CustomerMongoRepository
	CustMysql    repository.CustomerMysqlRepository
}

func NewController(
	custRepo repository.CustomerRepository,
	accRepo repository.AccountRepository,
	contRepo repository.ContractRepository,
	custMongoRepo repository.CustomerMongoRepository,
	custMysqlRepo repository.CustomerMysqlRepository,
) Controller {
	return &controller{
		CustomerRepo: custRepo,
		AccountRepo:  accRepo,
		ContractRepo: contRepo,
		CustMongo:    custMongoRepo,
		CustMysql:    custMysqlRepo,
	}
}

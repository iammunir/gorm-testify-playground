package controller

import (
	"learn-gorm/models"
	"learn-gorm/repository"
)

type Controller interface {
	GetCustomerByName(models.CustomerRequest) ([]models.Customer, *models.ErrorResponse)
	GetCustomerById(models.CustomerRequest) (*models.Customer, *models.ErrorResponse)
	GetCustomerWithAccount(models.CustomerRequest) ([]models.Customer, *models.ErrorResponse)
	GetCustomerWithAccountGoroutine(models.CustomerRequest) ([]models.Customer, *models.ErrorResponse)
	GetContractDetails(models.ContractRequest) ([]models.ContractDetails, *models.ErrorResponse)
}

type controller struct {
	CustomerRepo repository.CustomerRepository
	AccountRepo  repository.AccountRepository
	ContractRepo repository.ContractRepository
}

func NewController(
	custRepo repository.CustomerRepository,
	accRepo repository.AccountRepository,
	contRepo repository.ContractRepository,
) Controller {
	return &controller{
		CustomerRepo: custRepo,
		AccountRepo:  accRepo,
		ContractRepo: contRepo,
	}
}

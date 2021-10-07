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
}

type controller struct {
	CustomerRepo repository.CustomerRepository
	AccountRepo  repository.AccountRepository
}

func NewController(custRepo repository.CustomerRepository, accRepo repository.AccountRepository) Controller {
	return &controller{
		CustomerRepo: custRepo,
		AccountRepo:  accRepo,
	}
}

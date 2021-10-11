package controller

import (
	"learn-gorm/models"

	"github.com/stretchr/testify/mock"
)

type ControllerMockInterface interface {
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

type ControllerMock struct {
	Mock mock.Mock
}

func (ctrl *ControllerMock) GetCustomerByName(req models.CustomerRequest) ([]models.Customer, *models.ErrorResponse) {
	arguments := ctrl.Mock.Called(req)

	// handling data not found
	if arguments.Get(0) == nil {
		return nil, models.BuildError("data not found", "record not found")
	}

	customers := arguments.Get(0).([]models.Customer)
	return customers, nil
}

func (ctrl *ControllerMock) GetCustomerById(req models.CustomerRequest) (*models.Customer, *models.ErrorResponse) {
	arguments := ctrl.Mock.Called(req)

	// handling data not found
	if arguments.Get(0) == nil {
		return nil, models.BuildError("data not found", "record not found")
	}

	customer := arguments.Get(0).(models.Customer)
	return &customer, nil
}

func (ctrl *ControllerMock) GetCustomerWithAccount(req models.CustomerRequest) ([]models.Customer, *models.ErrorResponse) {
	arguments := ctrl.Mock.Called(req)

	// handling data not found
	if arguments.Get(0) == nil {
		return nil, models.BuildError("data not found", "record not found")
	}

	customers := arguments.Get(0).([]models.Customer)
	return customers, nil
}

func (ctrl *ControllerMock) GetCustomerWithAccountGoroutine(req models.CustomerRequest) ([]models.Customer, *models.ErrorResponse) {
	arguments := ctrl.Mock.Called(req)

	// handling data not found
	if arguments.Get(0) == nil {
		return nil, models.BuildError("data not found", "record not found")
	}

	customers := arguments.Get(0).([]models.Customer)
	return customers, nil
}

func (ctrl *ControllerMock) GetContractDetails(req models.ContractRequest) ([]models.ContractDetails, *models.ErrorResponse) {
	arguments := ctrl.Mock.Called(req)

	// handling data not found
	if arguments.Get(0) == nil {
		return nil, models.BuildError("data not found", "record not found")
	}

	customers := arguments.Get(0).([]models.ContractDetails)
	return customers, nil
}

func (ctrl *ControllerMock) GetCustomerByNameMongo(req models.ContractRequest) ([]models.CustomerMongo, *models.ErrorResponse) {
	arguments := ctrl.Mock.Called(req)

	// handling data not found
	if arguments.Get(0) == nil {
		return nil, models.BuildError("data not found", "record not found")
	}

	customers := arguments.Get(0).([]models.CustomerMongo)
	return customers, nil
}

func (ctrl *ControllerMock) GetCustomerByNameMongoWithSwitchObj(req models.ContractRequest) ([]models.CustomerMongo, *models.ErrorResponse) {
	arguments := ctrl.Mock.Called(req)

	// handling data not found
	if arguments.Get(0) == nil {
		return nil, models.BuildError("data not found", "record not found")
	}

	customers := arguments.Get(0).([]models.CustomerMongo)
	return customers, nil
}

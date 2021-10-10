package controller

import (
	"learn-gorm/models"
	"learn-gorm/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var customerRepository = &repository.CustomerRepositoryMock{Mock: mock.Mock{}}
var accountRepository = &repository.AccountRepositoryMock{Mock: mock.Mock{}}
var contractRepository = &repository.ContractRepositoryMock{Mock: mock.Mock{}}
var customerMongoRepository = &repository.CustomerMongoRepositoryMock{Mock: mock.Mock{}}
var customerMysqlRepository = &repository.CustomerMysqlRepositoryMock{Mock: mock.Mock{}}
var customerController = NewController(
	customerRepository,
	accountRepository,
	contractRepository,
	customerMongoRepository,
	customerMysqlRepository)

func TestGetCustomerByIdNotFound(t *testing.T) {

	// set data mocking
	customerRepository.Mock.On("FindById", 99999).Return(nil)

	// prepare controller parameter
	dataRequest := models.CustomerRequest{
		CustomerId: "99999",
		RequestDetails: models.RequestDetails{
			Address: 1,
			Company: 1,
			Contact: 0,
		},
	}

	// execute controller to test
	customer, err := customerController.GetCustomerById(dataRequest)

	// assert
	assert.NotNil(t, err, "Error should not be nil")
	assert.Nil(t, customer, "Customer data should be nil")
}

func TestGetCustomerByIdFound(t *testing.T) {
	// set data mocking
	customerData := models.Customer{
		FirstName: "John",
		LastName:  "Doe",
	}
	customerRepository.Mock.On("FindById", 10).Return(customerData)

	// prepare controller parameter
	dataRequest := models.CustomerRequest{
		CustomerId: "10",
		RequestDetails: models.RequestDetails{
			Address: 0,
			Company: 0,
			Contact: 0,
		},
	}

	// execute controller to test
	customer, err := customerController.GetCustomerById(dataRequest)

	// assert
	assert.Nil(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, "John", customer.FirstName, "First name should match")
	assert.Equal(t, "Doe", customer.LastName, "Last name should match")
}

package repository

import (
	"errors"
	"learn-gorm/models"

	"github.com/stretchr/testify/mock"
)

type CustomerRepositoryMockInterface interface {
	FindById(int) (*models.Customer, error)
	FindByName(string) ([]models.Customer, error)
	FindByNik(string) ([]models.Customer, error)
}

type CustomerRepositoryMock struct {
	Mock mock.Mock
}

func (repo *CustomerRepositoryMock) FindById(id int) (*models.Customer, error) {
	arguments := repo.Mock.Called(id)

	// for handling data not found
	if arguments.Get(0) == nil {
		return nil, errors.New("record not found")
	}

	customer := arguments.Get(0).(models.Customer)
	return &customer, nil
}

func (repo *CustomerRepositoryMock) FindByName(name string) ([]models.Customer, error) {
	arguments := repo.Mock.Called(name)

	// for handling data not found
	if arguments.Get(0) == nil {
		return nil, errors.New("record not found")
	}

	customers := arguments.Get(0).([]models.Customer)
	return customers, nil
}

func (repo *CustomerRepositoryMock) FindByNik(nik string) ([]models.Customer, error) {
	arguments := repo.Mock.Called(nik)

	// for handling data not found
	if arguments.Get(0) == nil {
		return nil, errors.New("record not found")
	}

	customers := arguments.Get(0).([]models.Customer)
	return customers, nil
}

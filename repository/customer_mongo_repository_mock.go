package repository

import (
	"errors"
	"learn-gorm/models"

	"github.com/stretchr/testify/mock"
)

type CustomerMongoRepositoryMockInterface interface {
	FindByName(string) ([]models.Customer, error)
}

type CustomerMongoRepositoryMock struct {
	Mock mock.Mock
}

func (repo *CustomerMongoRepositoryMock) FindByName(name string) ([]models.CustomerMongo, error) {
	arguments := repo.Mock.Called(name)

	// for handling data not found
	if arguments.Get(0) == nil {
		return nil, errors.New("record not found")
	}

	customers := arguments.Get(0).([]models.CustomerMongo)
	return customers, nil
}

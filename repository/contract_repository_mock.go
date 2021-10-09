package repository

import (
	"errors"
	"learn-gorm/models"

	"github.com/stretchr/testify/mock"
)

type ContractRepositoryMockInterface interface {
	FindByName(string) ([]models.ContractDetails, error)
}

type ContractRepositoryMock struct {
	Mock mock.Mock
}

func (repo *ContractRepositoryMock) FindByName(name string) ([]models.ContractDetails, error) {
	arguments := repo.Mock.Called(name)

	// for handling data not found
	if arguments.Get(0) == nil {
		return nil, errors.New("record not found")
	}

	contracts := arguments.Get(0).([]models.ContractDetails)
	return contracts, nil
}

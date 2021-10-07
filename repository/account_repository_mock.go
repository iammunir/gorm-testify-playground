package repository

import (
	"errors"
	"learn-gorm/models"

	"github.com/stretchr/testify/mock"
)

type AccountRepositoryMockInterface interface {
	FindByNik(string) ([]models.Account, error)
}

type AccountRepositoryMock struct {
	Mock mock.Mock
}

func (repo *AccountRepositoryMock) FindByNik(nik string) ([]models.Account, error) {
	arguments := repo.Mock.Called(nik)

	// for handling data not found
	if arguments.Get(0) == nil {
		return nil, errors.New("record not found")
	}

	accounts := arguments.Get(0).([]models.Account)
	return accounts, nil
}

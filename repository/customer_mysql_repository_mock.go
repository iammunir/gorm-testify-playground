package repository

import (
	"errors"
	"learn-gorm/models"

	"github.com/stretchr/testify/mock"
)

type CustomerMysqlRepositoryInterface interface {
	FindCustomerDetail([]string) ([]models.CustomerMongo, error)
	FindContractDetail([]string) ([]models.ContractMongo, error)
	FindComponentDetail([]string) ([]models.ComponentMongo, error)
	FindPartyRoleDetail([]string) ([]models.PartyRoleMongo, error)
}

type CustomerMysqlRepositoryMock struct {
	Mock mock.Mock
}

func (repo *CustomerMysqlRepositoryMock) FindCustomerDetail(ids []string) ([]models.CustomerMongo, error) {
	arguments := repo.Mock.Called(ids)

	// for handling data not found
	if arguments.Get(0) == nil {
		return nil, errors.New("record not found")
	}

	customers := arguments.Get(0).([]models.CustomerMongo)
	return customers, nil
}

func (repo *CustomerMysqlRepositoryMock) FindContractDetail(ids []string) ([]models.ContractMongo, error) {
	arguments := repo.Mock.Called(ids)

	// for handling data not found
	if arguments.Get(0) == nil {
		return nil, errors.New("record not found")
	}

	contracts := arguments.Get(0).([]models.ContractMongo)
	return contracts, nil
}

func (repo *CustomerMysqlRepositoryMock) FindComponentDetail(ids []string) ([]models.ComponentMongo, error) {
	arguments := repo.Mock.Called(ids)

	// for handling data not found
	if arguments.Get(0) == nil {
		return nil, errors.New("record not found")
	}

	components := arguments.Get(0).([]models.ComponentMongo)
	return components, nil
}

func (repo *CustomerMysqlRepositoryMock) FindPartyRoleDetail(ids []string) ([]models.PartyRoleMongo, error) {
	arguments := repo.Mock.Called(ids)

	// for handling data not found
	if arguments.Get(0) == nil {
		return nil, errors.New("record not found")
	}

	partyRoles := arguments.Get(0).([]models.PartyRoleMongo)
	return partyRoles, nil
}

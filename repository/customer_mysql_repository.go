package repository

import (
	"errors"
	"learn-gorm/models"

	"github.com/jinzhu/gorm"
)

type CustomerMysqlRepository interface {
	FindCustomerDetail([]string) ([]models.CustomerMongo, error)
	FindContractDetail([]string) ([]models.ContractMongo, error)
	FindComponentDetail([]string) ([]models.ComponentMongo, error)
	FindPartyRoleDetail([]string) ([]models.PartyRoleMongo, error)
}

type customerMysqlRepository struct {
	mysqlConn *gorm.DB
}

func GetCustomerMysqlRepository(conn *gorm.DB) CustomerMysqlRepository {
	return &customerMysqlRepository{
		mysqlConn: conn,
	}
}

func (repo *customerMysqlRepository) FindCustomerDetail(ids []string) ([]models.CustomerMongo, error) {
	var customers []models.CustomerMongo
	errExec := repo.mysqlConn.
		Table("con_customer").
		Select("*").
		Where("cont_id IN (?)", ids).
		Scan(&customers).Error

	if errExec != nil {
		return nil, errExec
	}
	if len(customers) == 0 {
		return nil, errors.New("record not found")
	}

	return customers, nil
}

func (repo *customerMysqlRepository) FindContractDetail(ids []string) ([]models.ContractMongo, error) {
	var contracts []models.ContractMongo
	errExec := repo.mysqlConn.
		Table("con_contract").
		Select("*").
		Where("contract_id IN (?)", ids).
		Scan(&contracts).Error

	if errExec != nil {
		return nil, errExec
	}
	if len(contracts) == 0 {
		return nil, errors.New("record not found")
	}

	return contracts, nil
}

func (repo *customerMysqlRepository) FindComponentDetail(ids []string) ([]models.ComponentMongo, error) {
	var components []models.ComponentMongo
	errExec := repo.mysqlConn.
		Table("con_component").
		Select("*").
		Where("component_id IN (?)", ids).
		Scan(&components).Error

	if errExec != nil {
		return nil, errExec
	}
	if len(components) == 0 {
		return nil, errors.New("record not found")
	}

	return components, nil
}

func (repo *customerMysqlRepository) FindPartyRoleDetail(ids []string) ([]models.PartyRoleMongo, error) {
	var partyRoles []models.PartyRoleMongo
	errExec := repo.mysqlConn.
		Table("con_party_role").
		Select("*").
		Where("party_role_id IN (?)", ids).
		Scan(&partyRoles).Error

	if errExec != nil {
		return nil, errExec
	}
	if len(partyRoles) == 0 {
		return nil, errors.New("record not found")
	}

	return partyRoles, nil
}

package repository

import (
	"errors"
	"learn-gorm/models"

	"github.com/jinzhu/gorm"
)

type ContractRepository interface {
	FindByName(string) ([]models.ContractDetails, error)
}

type contractRepository struct {
	mysqlConnection *gorm.DB
}

func GetContractRepository(conn *gorm.DB) ContractRepository {
	return &contractRepository{
		mysqlConnection: conn,
	}
}

func (repo *contractRepository) FindByName(name string) ([]models.ContractDetails, error) {
	var contracts []models.ContractDetails
	errExec := repo.mysqlConnection.
		Model(&contracts).
		Select("*").
		Where("SOUNDEX(full_name) LIKE SOUNDEX(?)", name).
		Scan(&contracts).
		Error

	if errExec != nil {
		return nil, errExec
	}
	if len(contracts) == 0 {
		return nil, errors.New("record not found")
	}

	return contracts, nil
}

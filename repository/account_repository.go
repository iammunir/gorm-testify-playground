package repository

import (
	"errors"
	"learn-gorm/models"
	"time"

	"github.com/jinzhu/gorm"
)

type AccountRepository interface {
	FindByNik(string) ([]models.Account, error)
}

type accountRepository struct {
	mysqlConnection *gorm.DB
}

func GetAccountRepository(conn *gorm.DB) AccountRepository {
	return &accountRepository{
		mysqlConnection: conn,
	}
}

func (repo *accountRepository) FindByNik(nik string) ([]models.Account, error) {
	var accounts []models.Account
	errExec := repo.mysqlConnection.
		Model(&accounts).
		Select("*").
		Where("nik = ?", nik).
		Scan(&accounts).
		Error

	if errExec != nil {
		return nil, errExec
	}

	if len(accounts) == 0 {
		return nil, errors.New("record not found")
	}

	time.Sleep(time.Millisecond * 2000)
	return accounts, nil
}

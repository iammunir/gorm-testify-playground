package repository

import (
	"errors"
	"learn-gorm/models"
	"time"

	"github.com/jinzhu/gorm"
)

type CustomerRepository interface {
	FindById(int) (*models.Customer, error)
	FindByName(string) ([]models.Customer, error)
	FindByNik(string) ([]models.Customer, error)
}

type customerRepository struct {
	mysqlConnection *gorm.DB
}

func GetCustomerRepository(conn *gorm.DB) CustomerRepository {
	return &customerRepository{
		mysqlConnection: conn,
	}
}

func (repo *customerRepository) FindById(id int) (*models.Customer, error) {
	var customer models.Customer
	errExec := repo.mysqlConnection.
		Model(&customer).
		Select("*").
		Where("id = ?", id).
		Scan(&customer).
		Error

	if errExec != nil {
		return nil, errExec
	}

	return &customer, nil
}

func (repo *customerRepository) FindByName(name string) ([]models.Customer, error) {
	var customers []models.Customer
	errExec := repo.mysqlConnection.
		Model(&customers).
		Select("*").
		Where("SOUNDEX(first_name) = SOUNDEX(?) OR last_name = SOUNDEX(?)", name, name).
		Scan(&customers).
		Error

	if errExec != nil {
		return nil, errExec
	}

	if len(customers) == 0 {
		return nil, errors.New("record not found")
	}

	return customers, nil
}

func (repo *customerRepository) FindByNik(nik string) ([]models.Customer, error) {
	var customers []models.Customer
	errExec := repo.mysqlConnection.
		Model(&customers).
		Select("*").
		Where("nik = ?", nik).
		Scan(&customers).
		Error

	if errExec != nil {
		return nil, errExec
	}

	if len(customers) == 0 {
		return nil, errors.New("record not found")
	}

	time.Sleep(time.Millisecond * 4000)
	return customers, nil
}

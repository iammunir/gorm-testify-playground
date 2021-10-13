package repository

import (
	"database/sql"
	"learn-gorm/models"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

var repo AccountRepository
var mockSql sqlmock.Sqlmock
var db *sql.DB

func TestMain(m *testing.M) {
	var err error
	db, mockSql, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic("error sql mock")
	}
	defer db.Close()

	gdb, errG := gorm.Open("mysql", db) // open gorm db
	if errG != nil {
		panic("error gorm")
	}

	repo = GetAccountRepository(gdb)

	m.Run()
}

func TestFindByNikSuccess(t *testing.T) {

	dummy := models.Account{
		NIK:         "4720329949345320",
		AccountNum:  "87880966",
		AccountName: "TabKeren",
		Balance:     "72200764",
	}
	dummy2 := models.Account{
		NIK:         "4720329949345320",
		AccountNum:  "87880967",
		AccountName: "Gaul",
		Balance:     "5132412",
	}

	query := "SELECT * FROM `customer_accounts` WHERE (nik = ?)"

	rows := sqlmock.NewRows([]string{"nik", "account_num", "account_name", "account_balance"}).
		AddRow(dummy.NIK, dummy.AccountNum, dummy.AccountName, dummy.Balance).
		AddRow(dummy2.NIK, dummy2.AccountNum, dummy2.AccountName, dummy2.Balance)

	mockSql.ExpectQuery(query).WithArgs(dummy.NIK).WillReturnRows(rows)

	accounts, errdb := repo.FindByNik(dummy.NIK)

	assert.Nil(t, errdb)
	assert.NotNil(t, accounts)
	assert.EqualValues(t, dummy.Balance, accounts[0].Balance, "balance of the first row should match")
}

func TestFindByNikDataNotFound(t *testing.T) {

	query := "SELECT * FROM `customer_accounts` WHERE (nik = ?)"

	rows := sqlmock.NewRows([]string{"nik", "account_num", "account_name", "account_balance"})

	mockSql.ExpectQuery(query).WithArgs("23952039").WillReturnRows(rows)

	accounts, errdb := repo.FindByNik("23952039")

	assert.Nil(t, accounts)
	assert.NotNil(t, errdb)
	assert.EqualError(t, errdb, "record not found")
}

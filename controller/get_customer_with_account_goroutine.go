package controller

import (
	"learn-gorm/models"

	"github.com/jinzhu/gorm"
)

func (ctrl *controller) GetCustomerWithAccountGoroutine(reqData models.CustomerRequest) ([]models.Customer, *models.ErrorResponse) {

	nik := reqData.Nik

	customersChan := make(chan []models.Customer, 2)
	accountsChan := make(chan []models.Account, 2)
	custErr := make(chan error, 2)
	accErr := make(chan error, 2)
	defer close(customersChan)
	defer close(accountsChan)
	defer close(custErr)
	defer close(accErr)

	go func() {
		customers, errCust := ctrl.CustomerRepo.FindByNik(nik)
		if errCust != nil {
			customersChan <- nil
			custErr <- errCust
		} else {
			customersChan <- customers
			custErr <- nil
		}

		for i, customer := range customers {
			customer.DefineRequestDetails(&reqData.RequestDetails)
			customers[i] = customer
		}
	}()

	go func() {
		accounts, errAcc := ctrl.AccountRepo.FindByNik(nik)
		if errAcc != nil {
			accountsChan <- nil
			accErr <- errAcc
		} else {
			accountsChan <- accounts
			accErr <- nil
		}
	}()

	errCust := <-custErr
	customers := <-customersChan
	if errCust != nil && gorm.IsRecordNotFoundError(errCust) {
		return nil, models.BuildError("data not found", errCust.Error())
	} else if errCust != nil {
		return nil, models.BuildError("error when executing db", errCust.Error())
	}

	errAcc := <-accErr
	accounts := <-accountsChan
	accountFound := true
	if errAcc != nil && gorm.IsRecordNotFoundError(errAcc) {
		accountFound = false
	} else if errAcc != nil {
		return nil, models.BuildError("error when executing db", errAcc.Error())
	}

	if accountFound {
		for i, customer := range customers {
			for _, account := range accounts {
				if customer.NIK == account.NIK {
					customer.Account = append(customer.Account, account)
				}
			}
			customers[i] = customer
		}
	}

	return customers, nil
}

package controller

import (
	"learn-gorm/models"

	"github.com/jinzhu/gorm"
)

func (ctrl *controller) GetCustomerWithAccount(reqData models.CustomerRequest) ([]models.Customer, *models.ErrorResponse) {

	nik := reqData.Nik

	customers, errCust := ctrl.CustomerRepo.FindByNik(nik)
	if errCust != nil && gorm.IsRecordNotFoundError(errCust) {
		return nil, models.BuildError("data not found", errCust.Error())
	} else if errCust != nil {
		return nil, models.BuildError("error when executing db", errCust.Error())
	}

	accountFound := true
	accounts, errAcc := ctrl.AccountRepo.FindByNik(nik)
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

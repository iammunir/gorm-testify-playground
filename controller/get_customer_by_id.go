package controller

import (
	"learn-gorm/models"
	"strconv"

	"github.com/jinzhu/gorm"
)

func (ctrl *controller) GetCustomerById(reqData models.CustomerRequest) (*models.Customer, *models.ErrorResponse) {

	customerIdStr := reqData.CustomerId
	customerId, err := strconv.Atoi(customerIdStr)
	if err != nil {
		return nil, models.BuildError("error when converting customer id", err.Error())
	}

	customer, errCust := ctrl.CustomerRepo.FindById(customerId)
	if errCust != nil && gorm.IsRecordNotFoundError(errCust) {
		return nil, models.BuildError("data not found", err.Error())
	} else if errCust != nil {
		return nil, models.BuildError("error when executing db", errCust.Error())
	}

	customer.DefineRequestDetails(&reqData.RequestDetails)

	return customer, nil
}

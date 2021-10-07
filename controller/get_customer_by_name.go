package controller

import (
	"learn-gorm/models"
)

func (ctrl *controller) GetCustomerByName(reqData models.CustomerRequest) ([]models.Customer, *models.ErrorResponse) {

	name := reqData.CustomerName

	customers, errCust := ctrl.CustomerRepo.FindByName(name)
	if errCust != nil && errCust.Error() == "record not found" {
		return nil, models.BuildError("data not found", errCust.Error())
	} else if errCust != nil {
		return nil, models.BuildError("error when executing db", errCust.Error())
	}

	for i, customer := range customers {
		customer.DefineRequestDetails(&reqData.RequestDetails)
		customers[i] = customer
	}

	return customers, nil
}

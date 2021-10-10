package controller

import (
	"learn-gorm/models"
)

func (ctrl *controller) GetCustomerByNameMongo(reqData models.ContractRequest) ([]models.CustomerMongo, *models.ErrorResponse) {

	name := reqData.CustomerName

	customers, errCust := ctrl.CustMongo.FindByName(name)
	if errCust != nil && errCust.Error() == "record not found" {
		return nil, models.BuildError("data not found", errCust.Error())
	} else if errCust != nil {
		return nil, models.BuildError("error when executing db", errCust.Error())
	}

	// get each model ids
	var contIds []string
	var contractIds []string
	var componentIds []string
	var partyRolesIds []string
	for _, customer := range customers {
		contIds = append(contIds, customer.ContId)
		if len(customer.Contracts) > 0 {
			for _, cont := range customer.Contracts {
				contractIds = append(contractIds, cont.ContractId)
				if len(cont.Components) > 0 {
					for _, comp := range cont.Components {
						componentIds = append(componentIds, comp.ComponentId)
						if len(comp.PartyRoles) > 0 {
							for _, role := range comp.PartyRoles {
								partyRolesIds = append(partyRolesIds, role.PartyRoleId)
							}
						}
					}
				}
			}
		}
	}

	customerChan := make(chan []models.CustomerMongo, 1)
	custErrChan := make(chan error, 1)
	defer close(customerChan)
	defer close(custErrChan)

	contractChan := make(chan []models.ContractMongo, 1)
	contErrChan := make(chan error, 1)
	defer close(contractChan)
	defer close(contErrChan)

	componentChan := make(chan []models.ComponentMongo, 1)
	compErrChan := make(chan error, 1)
	defer close(componentChan)
	defer close(compErrChan)

	partyRoleChan := make(chan []models.PartyRoleMongo, 1)
	partyErrChan := make(chan error, 1)
	defer close(partyRoleChan)
	defer close(partyErrChan)

	go ctrl.ConcurrentlyGetCustomerDetails(contIds, customerChan, custErrChan)
	go ctrl.ConcurrentlyGetContractDetails(contractIds, contractChan, contErrChan)
	go ctrl.ConcurrentlyGetComponentDetails(componentIds, componentChan, compErrChan)
	go ctrl.ConcurrentlyGetPartyRoleDetails(partyRolesIds, partyRoleChan, partyErrChan)

	customersDetail := <-customerChan
	custErr := <-custErrChan
	if custErr != nil && custErr.Error() == "record not found" {
		return nil, models.BuildError("customer details not found", custErr.Error())
	} else if custErr != nil {
		return nil, models.BuildError("error getting customer details", custErr.Error())
	}

	isContractFound := true
	isComponentFound := true
	isPartyRoleFound := true

	contractsDetail := <-contractChan
	contErr := <-contErrChan
	if contErr != nil && contErr.Error() == "record not found" {
		isContractFound = false
	} else if contErr != nil {
		return nil, models.BuildError("error getting contract details", contErr.Error())
	}

	componentsDetail := <-componentChan
	compErr := <-compErrChan
	if compErr != nil && compErr.Error() == "record not found" {
		isComponentFound = false
	} else if compErr != nil {
		return nil, models.BuildError("error getting component details", compErr.Error())
	}

	partyRolesDetail := <-partyRoleChan
	partyErr := <-partyErrChan
	if partyErr != nil && partyErr.Error() == "record not found" {
		isPartyRoleFound = false
	} else if partyErr != nil {
		return nil, models.BuildError("error getting party roles details", partyErr.Error())
	}

	for i, customer := range customers {
		for _, cust := range customersDetail {

			// merge customer details
			if customer.ContId == cust.ContId {
				customer.FullName = cust.FullName
				customer.NamaIbu = cust.NamaIbu

				if isContractFound {
					for i, custCont := range customer.Contracts {
						for _, cont := range contractsDetail {

							// merge contract details
							if custCont.ContractId == cont.ContractId {
								custCont.ContractName = cont.ContractName

								if isComponentFound {
									for i, custComp := range custCont.Components {
										for _, comp := range componentsDetail {

											// merge component details
											if custComp.ComponentId == comp.ComponentId {
												custComp.ComponentName = comp.ComponentName

												if isPartyRoleFound {
													for i, custParty := range custComp.PartyRoles {
														for _, party := range partyRolesDetail {

															// merge party roles
															if custParty.PartyRoleId == party.PartyRoleId {
																custParty.PartyRoleName = party.PartyRoleName
															}
														}
														custComp.PartyRoles[i] = custParty
													}
												}
											}

										}
										custCont.Components[i] = custComp
									}
								}
							}
						}
						customer.Contracts[i] = custCont
					}
				}
			}
		}

		customers[i] = customer
	}

	return customers, nil
}

func (ctrl *controller) ConcurrentlyGetCustomerDetails(ids []string, dataChan chan<- []models.CustomerMongo, errChan chan<- error) {
	data, err := ctrl.CustMysql.FindCustomerDetail(ids)
	if err != nil {
		dataChan <- nil
		errChan <- err
	} else {
		dataChan <- data
		errChan <- nil
	}
}

func (ctrl *controller) ConcurrentlyGetContractDetails(ids []string, dataChan chan<- []models.ContractMongo, errChan chan<- error) {
	data, err := ctrl.CustMysql.FindContractDetail(ids)
	if err != nil {
		dataChan <- nil
		errChan <- err
	} else {
		dataChan <- data
		errChan <- nil
	}
}

func (ctrl *controller) ConcurrentlyGetComponentDetails(ids []string, dataChan chan<- []models.ComponentMongo, errChan chan<- error) {
	data, err := ctrl.CustMysql.FindComponentDetail(ids)
	if err != nil {
		dataChan <- nil
		errChan <- err
	} else {
		dataChan <- data
		errChan <- nil
	}
}

func (ctrl *controller) ConcurrentlyGetPartyRoleDetails(ids []string, dataChan chan<- []models.PartyRoleMongo, errChan chan<- error) {
	data, err := ctrl.CustMysql.FindPartyRoleDetail(ids)
	if err != nil {
		dataChan <- nil
		errChan <- err
	} else {
		dataChan <- data
		errChan <- nil
	}
}

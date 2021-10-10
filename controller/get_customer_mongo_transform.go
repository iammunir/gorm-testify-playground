package controller

import (
	"learn-gorm/models"
)

func (ctrl *controller) GetCustomerByNameMongoWithSwitchObj(reqData models.ContractRequest) ([]models.CustomerMongo, *models.ErrorResponse) {

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
		var newCustObj models.CustomerMongo
		for _, cust := range customersDetail {

			// merge customer details
			if customer.ContId == cust.ContId {
				newCustObj = cust

				if isContractFound {
					for _, custCont := range customer.Contracts {
						var newContObj models.ContractMongo
						for _, cont := range contractsDetail {

							// merge contract details
							if custCont.ContractId == cont.ContractId {
								newContObj = cont

								if isComponentFound {
									for _, custComp := range custCont.Components {
										var newCompObj models.ComponentMongo
										for _, comp := range componentsDetail {

											// merge component details
											if custComp.ComponentId == comp.ComponentId {
												newCompObj = comp

												if isPartyRoleFound {
													for _, custParty := range custComp.PartyRoles {
														var newPartyObj models.PartyRoleMongo
														for _, party := range partyRolesDetail {

															// merge party roles
															if custParty.PartyRoleId == party.PartyRoleId {
																newPartyObj = party
															}
														}
														newCompObj.PartyRoles = append(newCompObj.PartyRoles, newPartyObj)
													}
												}

											}
										}
										newContObj.Components = append(newContObj.Components, newCompObj)
									}
								}
							}
						}
						newCustObj.Contracts = append(newCustObj.Contracts, newContObj)
					}
				}
			}
		}
		customers[i] = newCustObj
	}

	return customers, nil
}

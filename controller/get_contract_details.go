package controller

import (
	"learn-gorm/models"
)

func (ctrl *controller) GetContractDetails(reqData models.ContractRequest) ([]models.ContractDetails, *models.ErrorResponse) {

	name := reqData.CustomerName

	contDetails, errCont := ctrl.ContractRepo.FindByName(name)
	if errCont != nil && errCont.Error() == "record not found" {
		return nil, models.BuildError("data not found", errCont.Error())
	} else if errCont != nil {
		return nil, models.BuildError("error when executing db", errCont.Error())
	}

	var resultContracts []models.ContractDetails
	for i, contDetail := range contDetails {

		var currPartyRole models.PartyRole
		var currComponent models.Component
		var currContract models.Contract
		var currDetail models.ContractDetails

		currPartyRole = models.PartyRole{
			PartyRoleId:   contDetail.Contract.Component.PartyRole.PartyRoleId,
			PartyRoleName: contDetail.Contract.Component.PartyRole.PartyRoleName,
		}

		if i == 0 || contDetail.ContId != contDetails[i-1].ContId {

			currComponent = models.Component{
				ComponentId:   contDetail.Contract.Component.ComponentId,
				ComponentName: contDetail.Contract.Component.ComponentName,
				PartyRoles:    []models.PartyRole{currPartyRole},
			}

			currContract = models.Contract{
				ContractId:   contDetail.Contract.ContractId,
				ContractName: contDetail.Contract.ContractName,
				Components:   []models.Component{currComponent},
			}

			currDetail = models.ContractDetails{
				ContId:    contDetail.ContId,
				FullName:  contDetail.FullName,
				Contracts: []models.Contract{currContract},
			}

			resultContracts = append(resultContracts, currDetail)
			continue
		}

		// check if the component is inherit
		if i != 0 && contDetail.Contract.Component.ComponentId == contDetails[i-1].Contract.Component.ComponentId {

			// check if the contract is inherit
			if contDetail.Contract.ContractId == contDetails[i-1].Contract.ContractId {
				// append
				i_detail := len(resultContracts) - 1
				i_contr := len(resultContracts[i_detail].Contracts) - 1
				i_comp := len(resultContracts[i_detail].Contracts[i_contr].Components) - 1
				resultContracts[i_detail].
					Contracts[i_contr].
					Components[i_comp].
					PartyRoles = append(resultContracts[i_detail].
					Contracts[i_contr].
					Components[i_comp].
					PartyRoles, currPartyRole)
			} else {
				currContract = models.Contract{
					ContractId:   contDetail.Contract.ContractId,
					ContractName: contDetail.Contract.ContractName,
					Components:   []models.Component{currComponent},
				}
				// append
				i_detail := len(resultContracts) - 1
				resultContracts[i_detail].
					Contracts = append(resultContracts[i_detail].
					Contracts, currContract)
			}

		} else { // if not, create new component
			currComponent = models.Component{
				ComponentId:   contDetail.Contract.Component.ComponentId,
				ComponentName: contDetail.Contract.Component.ComponentName,
				PartyRoles:    []models.PartyRole{currPartyRole},
			}

			// check if the contract is inherit
			if contDetail.Contract.ContractId == contDetails[i-1].Contract.ContractId {
				// append
				i_detail := len(resultContracts) - 1
				i_contr := len(resultContracts[i_detail].Contracts) - 1
				resultContracts[i_detail].
					Contracts[i_contr].
					Components = append(resultContracts[i_detail].
					Contracts[i_contr].
					Components, currComponent)
			} else {
				currContract = models.Contract{
					ContractId:   contDetail.Contract.ContractId,
					ContractName: contDetail.Contract.ContractName,
					Components:   []models.Component{currComponent},
				}
				// append
				i_detail := len(resultContracts) - 1
				resultContracts[i_detail].
					Contracts = append(resultContracts[i_detail].
					Contracts, currContract)
			}
		}
	}

	return resultContracts, nil

}

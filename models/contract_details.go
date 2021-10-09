package models

type ContractDetails struct {
	ContId    string     `json:"cont_id" gorm:"column:cont_id"`
	FullName  string     `json:"full_name" gorm:"column:full_name"`
	Contracts []Contract `json:"contracts" gorm:"-"`
	Contract  Contract   `json:"-" gorm:"embedded"`
}

type Contract struct {
	ContractId   string      `json:"contract_id" gorm:"column:contract_id"`
	ContractName string      `json:"contract_name" gorm:"column:contract_name"`
	Components   []Component `json:"components" gorm:"-"`
	Component    Component   `json:"-" gorm:"embedded"`
}

type Component struct {
	ComponentId   string      `json:"component_id" gorm:"column:component_id"`
	ComponentName string      `json:"component_name" gorm:"column:component_name"`
	PartyRoles    []PartyRole `json:"party_roles" gorm:"-"`
	PartyRole     PartyRole   `json:"-" gorm:"embedded"`
}

type PartyRole struct {
	PartyRoleId   string `json:"party_role_id" gorm:"column:party_role_id"`
	PartyRoleName string `json:"party_role_name" gorm:"column:party_role_name"`
}

func (ContractDetails) TableName() string {
	return "contract_details"
}

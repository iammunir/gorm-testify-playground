package models

import "reflect"

type CustomerMongo struct {
	ContId    string          `json:"cont_id" bson:"cont_id" gorm:"cont_id"`
	FullName  string          `json:"full_name" bson:"full_name" gorm:"full_name"`
	NamaIbu   string          `json:"nama_ibu" bson:"omitempty" gorm:"nama_ibu"`
	Contracts []ContractMongo `json:"contracts" bson:"contracts,omitempty" gorm:"-"`
}

func (cust CustomerMongo) IsEmpty() bool {
	return reflect.DeepEqual(cust, CustomerMongo{})
}

type ContractMongo struct {
	ContractId   string           `json:"contract_id" bson:"contract_id" gorm:"column:contract_id"`
	ContractName string           `json:"contract_name" bson:"omitempty" gorm:"column:contract_name"`
	Components   []ComponentMongo `json:"components" bson:"components,omitempty" gorm:"-"`
}

func (cont ContractMongo) IsEmpty() bool {
	return reflect.DeepEqual(cont, ContractMongo{})
}

type ComponentMongo struct {
	ComponentId   string           `json:"component_id" bson:"component_id" gorm:"column:component_id"`
	ComponentName string           `json:"component_name" bson:"omitempty" gorm:"column:component_name"`
	PartyRoles    []PartyRoleMongo `json:"party_roles" bson:"party_roles,omitempty" gorm:"-"`
}

func (comp ComponentMongo) IsEmpty() bool {
	return reflect.DeepEqual(comp, ComponentMongo{})
}

type PartyRoleMongo struct {
	PartyRoleId   string `json:"party_role_id" bson:"party_role_id" gorm:"column:party_role_id"`
	PartyRoleName string `json:"party_role_name" bson:"party_role_name" gorm:"column:party_role_name"`
}

func (part PartyRoleMongo) IsEmpty() bool {
	return reflect.DeepEqual(part, PartyRoleMongo{})
}

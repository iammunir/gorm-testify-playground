package models

type CustomerRequest struct {
	Nik            string         `json:"nik,omitempty" validate:"required_without=customer_id"`
	CustomerId     string         `json:"customer_id,omitempty" validate:"required_without=customer_nik"`
	CustomerName   string         `json:"customer_name,omitempty"`
	RequestDetails RequestDetails `json:"request_details" binding:"required"`
}

type RequestDetails struct {
	Address int8 `json:"address,omitempty"`
	Company int8 `json:"company,omitempty"`
	Contact int8 `json:"contact,omitempty"`
}

type ContractRequest struct {
	CustomerName string `json:"customer_name" binding:"required"`
}

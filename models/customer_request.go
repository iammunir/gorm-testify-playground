package models

type CustomerRequest struct {
	CustomerId     string         `json:"customer_id,omitempty" validate:"required_without=customer_name"`
	CustomerName   string         `json:"customer_name,omitempty" validate:"required_without=customer_id"`
	RequestDetails RequestDetails `json:"request_details" binding:"required"`
}

type RequestDetails struct {
	Address int8 `json:"address,omitempty"`
	Company int8 `json:"company,omitempty"`
	Contact int8 `json:"contact,omitempty"`
}

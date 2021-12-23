package requests

import (
	"domain-driven-design-layout/domain/entities"
)

type CreateAddressRequest struct {
	Street string  `json:"street" binding:"required"`
	Number int32   `json:"number" binding:"required"`
	City   *string `json:"city"`
}

func (req *CreateAddressRequest) ToAddressPrototype() entities.AddressPrototype {
	return entities.AddressPrototype{
		Street: req.Street,
		Number: req.Number,
		City:   req.City,
	}
}

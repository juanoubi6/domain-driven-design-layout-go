package requests

import (
	"domain-driven-design-layout/domain/entities"
	"time"
)

type CreateUserRequest struct {
	FirstName string    `json:"first_name" binding:"required"`
	LastName  string    `json:"last_name" binding:"required"`
	BirthDate time.Time `json:"birth_date" binding:"required"`
	Addresses []struct {
		Street string  `json:"street" binding:"required"`
		Number int32   `json:"number" binding:"required"`
		City   *string `json:"city"`
	} `json:"addresses" binding:"required"`
}

func (cur *CreateUserRequest) ToUserPrototype() entities.UserPrototype {
	var addressPrototypes []entities.AddressPrototype

	for _, address := range cur.Addresses {
		addressPrototypes = append(addressPrototypes, entities.AddressPrototype{
			Street: address.Street,
			Number: address.Number,
			City:   address.City,
		})
	}

	return entities.UserPrototype{
		FirstName:           cur.FirstName,
		LastName:            cur.LastName,
		BirthDate:           cur.BirthDate,
		AddressesPrototypes: addressPrototypes,
	}
}

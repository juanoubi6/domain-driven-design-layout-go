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

func (req *CreateUserRequest) ToUserPrototype() entities.UserPrototype {
	var addressPrototypes []entities.AddressPrototype

	for _, address := range req.Addresses {
		addressPrototypes = append(addressPrototypes, entities.AddressPrototype{
			Street: address.Street,
			Number: address.Number,
			City:   address.City,
		})
	}

	return entities.UserPrototype{
		FirstName:           req.FirstName,
		LastName:            req.LastName,
		BirthDate:           req.BirthDate,
		AddressesPrototypes: addressPrototypes,
	}
}

type FindUsersByIdListRequest struct {
	UserIDs []int64 `json:"user_ids" binding:"required"`
}

type UpdateUserRequest struct {
	ID        int64     `json:"id" binding:"required"`
	FirstName string    `json:"first_name" binding:"required"`
	LastName  string    `json:"last_name" binding:"required"`
	BirthDate time.Time `json:"birth_date" binding:"required"`
}

func (req *UpdateUserRequest) ToUser() entities.User {
	return entities.User{
		ID:        req.ID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		BirthDate: req.BirthDate,
		Addresses: []entities.Address{},
	}
}

package domain

import (
	"domain-driven-design-layout/domain/entities"
	"time"
)

func CreateUserPrototype() entities.UserPrototype {
	country := "Argentina"

	return entities.UserPrototype{
		FirstName: "test",
		LastName:  "name",
		BirthDate: time.Now(),
		AddressesPrototypes: []entities.AddressPrototype{
			{
				Street: "street 1",
				Number: 1,
				City:   &country,
			},
			{
				Street: "street 2",
				Number: 2,
				City:   nil,
			},
		},
	}
}

func CreateUser() entities.User {
	country := "Argentina"

	return entities.User{
		ID:        1,
		FirstName: "test",
		LastName:  "name",
		BirthDate: time.Now(),
		Addresses: []entities.Address{
			{
				ID:     1,
				UserID: 1,
				Street: "street 1",
				Number: 1,
				City:   &country,
			},
			{
				ID:     2,
				UserID: 2,
				Street: "street 1",
				Number: 1,
				City:   nil,
			},
		},
	}
}

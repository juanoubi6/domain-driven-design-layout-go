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
	prototype := CreateUserPrototype()

	return entities.User{
		ID:        1,
		FirstName: prototype.FirstName,
		LastName:  prototype.LastName,
		BirthDate: prototype.BirthDate,
		Addresses: []entities.Address{
			{
				ID:     1,
				UserID: 1,
				Street: prototype.AddressesPrototypes[0].Street,
				Number: prototype.AddressesPrototypes[0].Number,
				City:   prototype.AddressesPrototypes[0].City,
			},
			{
				ID:     2,
				UserID: 1,
				Street: prototype.AddressesPrototypes[1].Street,
				Number: prototype.AddressesPrototypes[1].Number,
				City:   prototype.AddressesPrototypes[1].City,
			},
		},
	}
}

func CreateAddressPrototype() entities.AddressPrototype {
	country := "Argentina"

	return entities.AddressPrototype{
		Street: "street 1",
		Number: 1,
		City:   &country,
	}
}

func CreateAddress() entities.Address {
	prototype := CreateAddressPrototype()

	return entities.Address{
		ID:     1,
		UserID: 1,
		Street: prototype.Street,
		Number: prototype.Number,
		City:   prototype.City,
	}
}

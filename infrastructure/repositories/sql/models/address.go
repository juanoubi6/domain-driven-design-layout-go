package models

import (
	"domain-driven-design-layout/domain/entities"
)

type AddressModel struct {
	ID     int64   `db:"id"`
	UserID int64   `db:"user_id"`
	Street string  `db:"street"`
	Number int32   `db:"number"`
	City   *string `db:"city"`
}

func (am *AddressModel) ToAddress() entities.Address {
	return entities.Address{
		ID:     am.ID,
		UserID: am.UserID,
		Street: am.Street,
		Number: am.Number,
		City:   am.City,
	}
}

func CreateAddressModelFromPrototype(prototype entities.AddressPrototype, userId int64) AddressModel {
	var addressModel AddressModel

	addressModel.UserID = userId
	addressModel.Street = prototype.Street
	addressModel.Number = prototype.Number
	addressModel.City = prototype.City

	return addressModel
}

package models

import (
	"database/sql"
	"domain-driven-design-layout/domain/entities"
)

type AddressModel struct {
	ID     int64          `db:"id"`
	UserID int64          `db:"user_id"`
	Street string         `db:"street"`
	Number int32          `db:"number"`
	City   sql.NullString `db:"city"`
}

func CreateAddressModelFromPrototype(prototype entities.AddressPrototype, userId int64) AddressModel {
	var addressModel AddressModel

	addressModel.UserID = userId
	addressModel.Street = prototype.Street
	addressModel.Number = prototype.Number

	if prototype.City != nil {
		addressModel.City = sql.NullString{String: *prototype.City, Valid: true}
	} else {
		addressModel.City = sql.NullString{String: "", Valid: false}
	}

	return addressModel
}

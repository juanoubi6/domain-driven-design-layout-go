package models

import (
	"domain-driven-design-layout/domain/entities"
	"time"
)

type UserModel struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDate time.Time `json:"birth_date"`
}

func UserPrototypeToModel(prototype entities.UserPrototype) UserModel {
	return UserModel{
		ID:        0,
		FirstName: prototype.FirstName,
		LastName:  prototype.LastName,
		BirthDate: prototype.BirthDate,
	}
}

type AddressModel struct {
	ID     int64   `json:"id"`
	UserID int64   `json:"user_id"`
	Street string  `json:"street"`
	Number int32   `json:"number"`
	City   *string `json:"city"`
}

func AddressPrototypeToModel(prototype entities.AddressPrototype, userID int64) AddressModel {
	return AddressModel{
		ID:     0,
		UserID: userID,
		Street: prototype.Street,
		Number: prototype.Number,
		City:   prototype.City,
	}
}

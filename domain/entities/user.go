package entities

import "time"

type User struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDate time.Time `json:"birth_date"`
	Addresses []Address `json:"addresses"`
}

type Address struct {
	ID     int64   `json:"id"`
	UserID int64   `json:"-"`
	Street string  `json:"street"`
	Number int32   `json:"number"`
	City   *string `json:"city"`
}

type UserPrototype struct {
	FirstName           string
	LastName            string
	BirthDate           time.Time
	AddressesPrototypes []AddressPrototype
}

type AddressPrototype struct {
	Street string
	Number int32
	City   *string
}

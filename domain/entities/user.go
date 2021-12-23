package entities

import "time"

type User struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDate time.Time `json:"birth_date"`
	Addresses []Address `json:"addresses"`
}

type UserPrototype struct {
	FirstName           string
	LastName            string
	BirthDate           time.Time
	AddressesPrototypes []AddressPrototype
}

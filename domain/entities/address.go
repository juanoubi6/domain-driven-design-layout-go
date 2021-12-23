package entities

type Address struct {
	ID     int64   `json:"id"`
	UserID int64   `json:"-"`
	Street string  `json:"street"`
	Number int32   `json:"number"`
	City   *string `json:"city"`
}

type AddressPrototype struct {
	Street string
	Number int32
	City   *string
}

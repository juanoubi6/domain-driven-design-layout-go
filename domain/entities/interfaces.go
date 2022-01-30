package entities

type UserRepository interface {
	GetUser(int64) (*User, error)
	GetUsers([]int64) ([]User, error)
	CreateUser(UserPrototype) (User, error)
	UpdateUser(User) (User, error)
	DeleteUser(int64) error
}

type AddressRepository interface {
	CreateAddress(int64, AddressPrototype) (Address, error)
	DeleteAddress(int64) error
	GetAddress(int64) (*Address, error)
}

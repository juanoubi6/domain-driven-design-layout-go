package entities

type UserRepository interface {
	GetUser(userID int64) (*User, error)
	GetUsers(userIDs []int64) ([]User, error)
	CreateUser(userPrototype UserPrototype) (User, error)
	UpdateUser(user User) (User, error)
	DeleteUser(userID int64) error
}

type AddressRepository interface {
	CreateAddress(userID int64, addressPrototype AddressPrototype) (Address, error)
	DeleteAddress(addressID int64) error
	GetAddress(addressID int64) (*Address, error)
	DeleteUserAddresses(userID int64) error
}

type MainDatabase interface {
	UserRepository
	AddressRepository
}

type TxRepositoryCreator interface {
	CreateMainDatabase() (MainDatabase, error)
}

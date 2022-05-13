package entities

import "context"

type UserRepository interface {
	GetUser(ctx ApplicationContext, userID int64) (*User, error)
	GetUsers(ctx ApplicationContext, userIDs []int64) ([]User, error)
	CreateUser(ctx ApplicationContext, userPrototype UserPrototype) (User, error)
	UpdateUser(ctx ApplicationContext, user User) (User, error)
	DeleteUser(ctx ApplicationContext, userID int64) error
}

type AddressRepository interface {
	CreateAddress(ctx ApplicationContext, userID int64, addressPrototype AddressPrototype) (Address, error)
	DeleteAddress(ctx ApplicationContext, addressID int64) error
	GetAddress(ctx ApplicationContext, addressID int64) (*Address, error)
	DeleteUserAddresses(ctx ApplicationContext, userID int64) error
}

type MainDatabase interface {
	UserRepository
	AddressRepository
	CommitTx() error
	RollbackTx() error
}

type TxRepositoryCreator interface {
	CreateTxMainDatabase(ctx context.Context) (MainDatabase, error)
}

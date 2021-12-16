package entities

type UserRepository interface {
	GetUser(id int64) (User, error)
	GetUsers(id []int64) ([]User, error)
	CreateUser(prototype UserPrototype) (User, error)
	UpdateUser(entity User) (User, error)
	DeleteUser(id int64) error
}

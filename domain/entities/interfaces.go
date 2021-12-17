package entities

type UserRepository interface {
	GetUser(int64) (*User, error)
	GetUsers([]int64) ([]User, error)
	CreateUser(UserPrototype) (User, error)
	UpdateUser(User) (User, error)
	DeleteUser(int64) error
}

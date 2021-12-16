package users

import (
	"domain-driven-design-layout/domain/entities"
)

type CreateUser struct {
	userRepository entities.UserRepository
}

func NewCreateUser(repository entities.UserRepository) (*CreateUser, error) {
	result := &CreateUser{
		userRepository: repository,
	}

	return result, nil
}

func (act *CreateUser) Execute(prototype entities.UserPrototype) (entities.User, error) {
	//Execute any business logic or validations you need
	var user entities.User
	user, err := act.userRepository.CreateUser(prototype)
	if err != nil {
		return user, err
	}

	return user, nil
}

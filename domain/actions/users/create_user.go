package users

import (
	"domain-driven-design-layout/domain/entities"
)

type CreateUser interface {
	Execute(entities.UserPrototype) (entities.User, error)
}

type CreateUserAction struct {
	userRepository entities.UserRepository
}

func NewCreateUserAction(repository entities.UserRepository) (CreateUser, error) {
	result := CreateUserAction{
		userRepository: repository,
	}

	return &result, nil
}

func (act *CreateUserAction) Execute(prototype entities.UserPrototype) (entities.User, error) {
	//Execute any business logic or validations you need
	var user entities.User
	user, err := act.userRepository.CreateUser(prototype)
	if err != nil {
		return user, err
	}

	return user, nil
}

package users

import (
	"domain-driven-design-layout/domain/entities"
	"fmt"
)

type CreateUser interface {
	Execute(ctx entities.ApplicationContext, proto entities.UserPrototype) (entities.User, error)
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

func (act *CreateUserAction) Execute(ctx entities.ApplicationContext, prototype entities.UserPrototype) (entities.User, error) {
	//Execute any business logic or validations you need
	var user entities.User
	user, err := act.userRepository.CreateUser(ctx, prototype)
	if err != nil {
		return user, fmt.Errorf("user could not be created. Error: %v", err)
	}

	return user, nil
}

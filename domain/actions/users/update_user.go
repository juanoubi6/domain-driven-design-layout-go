package users

import (
	"domain-driven-design-layout/domain/entities"
	"fmt"
)

type UpdateUser interface {
	Execute(ctx entities.ApplicationContext, user entities.User) (entities.User, error)
}

type UpdateUserAction struct {
	userRepository entities.UserRepository
}

func NewUpdateUserAction(repository entities.UserRepository) (UpdateUser, error) {
	result := UpdateUserAction{
		userRepository: repository,
	}

	return &result, nil
}

func (act *UpdateUserAction) Execute(ctx entities.ApplicationContext, user entities.User) (entities.User, error) {
	//Execute any business logic or validations you need
	user, err := act.userRepository.UpdateUser(ctx, user)
	if err != nil {
		return user, fmt.Errorf("user could not be updated. Error: %v", err)
	}

	return user, nil
}

package users

import (
	"domain-driven-design-layout/domain/entities"
)

type UpdateUser interface {
	Execute(user entities.User) (entities.User, error)
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

func (act *UpdateUserAction) Execute(user entities.User) (entities.User, error) {
	//Execute any business logic or validations you need
	user, err := act.userRepository.UpdateUser(user)
	if err != nil {
		return user, err
	}

	return user, nil
}
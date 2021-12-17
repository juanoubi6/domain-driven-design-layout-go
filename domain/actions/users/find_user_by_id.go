package users

import (
	"domain-driven-design-layout/domain/entities"
)

type FindUserById interface {
	Execute(int64) (*entities.User, error)
}

type FindUserByIdAction struct {
	userRepository entities.UserRepository
}

func NewFindUserByIdAction(repository entities.UserRepository) (FindUserById, error) {
	result := FindUserByIdAction{
		userRepository: repository,
	}

	return &result, nil
}

func (act *FindUserByIdAction) Execute(id int64) (*entities.User, error) {
	//Execute any business logic or validations you need
	var user *entities.User
	user, err := act.userRepository.GetUser(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

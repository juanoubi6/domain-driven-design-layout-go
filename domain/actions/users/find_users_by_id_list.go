package users

import (
	"domain-driven-design-layout/domain/entities"
)

type FindUsersByIdList interface {
	Execute([]int64) ([]entities.User, error)
}

type FindUsersByIdListAction struct {
	userRepository entities.UserRepository
}

func NewFindUsersByIdListAction(repository entities.UserRepository) (FindUsersByIdList, error) {
	result := FindUsersByIdListAction{
		userRepository: repository,
	}

	return &result, nil
}

func (act *FindUsersByIdListAction) Execute(ids []int64) ([]entities.User, error) {
	//Execute any business logic or validations you need
	var users []entities.User
	users, err := act.userRepository.GetUsers(ids)
	if err != nil {
		return users, err
	}

	return users, nil
}

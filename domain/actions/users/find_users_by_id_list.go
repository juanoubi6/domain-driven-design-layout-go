package users

import (
	"domain-driven-design-layout/domain/entities"
	"fmt"
)

type FindUsersByIdList interface {
	Execute(ctx entities.ApplicationContext, userIDs []int64) ([]entities.User, error)
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

func (act *FindUsersByIdListAction) Execute(ctx entities.ApplicationContext, ids []int64) ([]entities.User, error) {
	//Execute any business logic or validations you need
	var users []entities.User
	users, err := act.userRepository.GetUsers(ctx, ids)
	if err != nil {
		return users, fmt.Errorf("users could not be found. Error: %v", err)
	}

	return users, nil
}

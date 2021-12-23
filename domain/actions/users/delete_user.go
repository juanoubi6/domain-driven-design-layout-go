package users

import (
	"domain-driven-design-layout/domain/entities"
	"fmt"
)

type DeleteUser interface {
	Execute(int64) error
}

type DeleteUserAction struct {
	userRepository entities.UserRepository
}

func NewDeleteUserAction(repository entities.UserRepository) (DeleteUser, error) {
	result := DeleteUserAction{
		userRepository: repository,
	}

	return &result, nil
}

func (act *DeleteUserAction) Execute(id int64) error {
	//Execute any business logic or validations you need
	if err := act.userRepository.DeleteUser(id); err != nil {
		return fmt.Errorf("user could not be deleted. Error: %v", err)
	}

	return nil
}

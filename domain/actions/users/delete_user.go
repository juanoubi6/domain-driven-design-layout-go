package users

import (
	"domain-driven-design-layout/domain/entities"
	"fmt"
)

type DeleteUser interface {
	Execute(int64) error
}

type DeleteUserAction struct {
	txRepositoryCreator entities.TxRepositoryCreator
}

func NewDeleteUserAction(txRepositoryCreator entities.TxRepositoryCreator) (DeleteUser, error) {
	result := DeleteUserAction{
		txRepositoryCreator: txRepositoryCreator,
	}

	return &result, nil
}

func (act *DeleteUserAction) Execute(id int64) error {
	//Execute any business logic or validations you need
	mainDatabase, err := act.txRepositoryCreator.CreateMainDatabase()
	if err != nil {
		return fmt.Errorf("could not create repository. Error: %w", err)
	}

	if err = mainDatabase.DeleteUser(id); err != nil {
		return fmt.Errorf("user could not be deleted. Error: %w", err)
	}

	if err = mainDatabase.DeleteUserAddresses(id); err != nil {
		return fmt.Errorf("user addresses could not be deleted. Error: %w", err)
	}

	return nil
}

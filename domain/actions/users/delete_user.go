package users

import (
	"context"
	"domain-driven-design-layout/domain/entities"
	"fmt"
)

type DeleteUser interface {
	Execute(ctx entities.ApplicationContext, userID int64) error
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

func (act *DeleteUserAction) Execute(ctx entities.ApplicationContext, id int64) error {
	mainDatabase, err := act.txRepositoryCreator.CreateTxMainDatabase(context.Background())
	if err != nil {
		return fmt.Errorf("could not create repository. Error: %w", err)
	}
	defer mainDatabase.RollbackTx()

	if err = mainDatabase.DeleteUser(ctx, id); err != nil {
		return fmt.Errorf("user could not be deleted. Error: %w", err)
	}

	if err = mainDatabase.DeleteUserAddresses(ctx, id); err != nil {
		return fmt.Errorf("user addresses could not be deleted. Error: %w", err)
	}

	// I could have 2 http calls here that do some stuff, and which are part of the logical
	// operation of "deleting a user". If any of them fails, I will be able to rollback all
	// database operations that were executed here.

	if err = mainDatabase.CommitTx(); err != nil {
		return fmt.Errorf("failed to execute operation. Error: %w", err)
	}

	return nil
}

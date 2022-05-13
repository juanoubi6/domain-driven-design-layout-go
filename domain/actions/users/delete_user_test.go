package users

import (
	"domain-driven-design-layout/domain/entities"
	"domain-driven-design-layout/domain/mocks"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestDeleteUser_Execute_Success(t *testing.T) {
	var mainDatabaseMock = new(mocks.MainDatabaseMock)
	var txRepositoryCreatorMock = new(mocks.TxRepositoryCreatorMock)
	var deleteUserAction, _ = NewDeleteUserAction(txRepositoryCreatorMock)

	var userId int64 = 1

	txRepositoryCreatorMock.On("CreateTxMainDatabase", mock.Anything).Return(mainDatabaseMock, nil)
	mainDatabaseMock.On("RollbackTx").Return(nil).Once()
	mainDatabaseMock.On("DeleteUser", mock.Anything, userId).Return(nil)
	mainDatabaseMock.On("DeleteUserAddresses", mock.Anything, userId).Return(nil)
	mainDatabaseMock.On("CommitTx").Return(nil).Once()

	err := deleteUserAction.Execute(entities.CreateEmptyAppContext(), userId)

	assert.Nil(t, err)
	txRepositoryCreatorMock.AssertExpectations(t)
	mainDatabaseMock.AssertExpectations(t)
}

func TestDeleteUser_Execute_FailsIfAnyRepositoryMethodFails(t *testing.T) {
	var mainDatabaseMock = new(mocks.MainDatabaseMock)
	var txRepositoryCreatorMock = new(mocks.TxRepositoryCreatorMock)
	var deleteUserAction, _ = NewDeleteUserAction(txRepositoryCreatorMock)

	txRepositoryCreatorMock.On("CreateTxMainDatabase", mock.Anything).Return(mainDatabaseMock, nil)
	mainDatabaseMock.On("RollbackTx").Return(nil).Once()
	mainDatabaseMock.On("DeleteUser", mock.Anything, int64(1)).Return(errors.New("some error"))

	err := deleteUserAction.Execute(entities.CreateEmptyAppContext(), 1)

	assert.NotNil(t, err)
	txRepositoryCreatorMock.AssertExpectations(t)
	mainDatabaseMock.AssertExpectations(t)
}

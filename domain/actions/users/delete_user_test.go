package users

import (
	"domain-driven-design-layout/domain"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestDeleteUser_Execute_Success(t *testing.T) {
	var mainDatabaseMock = new(domain.MainDatabaseMock)
	var txRepositoryCreatorMock = new(domain.TxRepositoryCreatorMock)
	var deleteUserAction, _ = NewDeleteUserAction(txRepositoryCreatorMock)

	var userId int64 = 1

	txRepositoryCreatorMock.On("CreateTxMainDatabase", mock.Anything).Return(mainDatabaseMock, nil)
	mainDatabaseMock.On("RollbackTx").Return(nil).Once()
	mainDatabaseMock.On("DeleteUser", userId).Return(nil)
	mainDatabaseMock.On("DeleteUserAddresses", userId).Return(nil)
	mainDatabaseMock.On("CommitTx").Return(nil).Once()

	err := deleteUserAction.Execute(userId)

	assert.Nil(t, err)
	txRepositoryCreatorMock.AssertExpectations(t)
	mainDatabaseMock.AssertExpectations(t)
}

func TestDeleteUser_Execute_FailsIfAnyRepositoryMethodFails(t *testing.T) {
	var mainDatabaseMock = new(domain.MainDatabaseMock)
	var txRepositoryCreatorMock = new(domain.TxRepositoryCreatorMock)
	var deleteUserAction, _ = NewDeleteUserAction(txRepositoryCreatorMock)

	txRepositoryCreatorMock.On("CreateTxMainDatabase", mock.Anything).Return(mainDatabaseMock, nil)
	mainDatabaseMock.On("RollbackTx").Return(nil).Once()
	mainDatabaseMock.On("DeleteUser", int64(1)).Return(errors.New("some error"))

	err := deleteUserAction.Execute(1)

	assert.NotNil(t, err)
	txRepositoryCreatorMock.AssertExpectations(t)
	mainDatabaseMock.AssertExpectations(t)
}

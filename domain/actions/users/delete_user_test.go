package users

import (
	"domain-driven-design-layout/domain"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteUser_Execute_Success(t *testing.T) {
	var mainDatabaseMock = new(domain.MainDatabaseMock)
	var txRepositoryCreatorMock = new(domain.TxRepositoryCreatorMock)
	var deleteUserAction, _ = NewDeleteUserAction(txRepositoryCreatorMock)

	var userId int64 = 1

	txRepositoryCreatorMock.On("CreateMainDatabase").Return(mainDatabaseMock, nil)
	mainDatabaseMock.On("DeleteUser", userId).Return(nil)
	mainDatabaseMock.On("DeleteUserAddresses", userId).Return(nil)

	err := deleteUserAction.Execute(userId)

	assert.Nil(t, err)
	txRepositoryCreatorMock.AssertExpectations(t)
	mainDatabaseMock.AssertExpectations(t)
}

func TestDeleteUser_Execute_FailsIfAnyRepositoryMethodFails(t *testing.T) {
	var mainDatabaseMock = new(domain.MainDatabaseMock)
	var txRepositoryCreatorMock = new(domain.TxRepositoryCreatorMock)
	var deleteUserAction, _ = NewDeleteUserAction(txRepositoryCreatorMock)

	txRepositoryCreatorMock.On("CreateMainDatabase").Return(mainDatabaseMock, nil)
	mainDatabaseMock.On("DeleteUser", int64(1)).Return(errors.New("some error"))

	err := deleteUserAction.Execute(1)

	assert.NotNil(t, err)
	txRepositoryCreatorMock.AssertExpectations(t)
	mainDatabaseMock.AssertExpectations(t)
}

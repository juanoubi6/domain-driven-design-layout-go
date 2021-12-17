package users

import (
	"domain-driven-design-layout/domain"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestDeleteUser_Execute_Success(t *testing.T) {
	var userRepositoryMock = new(domain.UserRepositoryMock)
	var deleteUserAction, _ = NewDeleteUserAction(userRepositoryMock)

	var userId int64 = 1

	userRepositoryMock.On("DeleteUser", userId).Return(nil)

	err := deleteUserAction.Execute(userId)

	assert.Nil(t, err)
	userRepositoryMock.AssertExpectations(t)
}

func TestDeleteUser_Execute_FailsIfUserRepositoryFails(t *testing.T) {
	var userRepositoryMock = new(domain.UserRepositoryMock)
	var deleteUserAction, _ = NewDeleteUserAction(userRepositoryMock)

	userRepositoryMock.On("DeleteUser", mock.Anything).Return(errors.New("error"))

	err := deleteUserAction.Execute(1)

	assert.NotNil(t, err)
	userRepositoryMock.AssertExpectations(t)
}

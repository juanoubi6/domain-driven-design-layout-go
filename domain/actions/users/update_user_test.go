package users

import (
	"domain-driven-design-layout/domain"
	"domain-driven-design-layout/domain/entities"
	"domain-driven-design-layout/domain/mocks"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestUpdateUser_Execute_Success(t *testing.T) {
	var userRepositoryMock = new(mocks.MainDatabaseMock)
	var updateUserAction, _ = NewUpdateUserAction(userRepositoryMock)

	var userToUpdate = domain.CreateUser()

	var expected = domain.CreateUser()
	expected.FirstName = "New name"

	userRepositoryMock.On("UpdateUser", mock.Anything, userToUpdate).Return(expected, nil)

	result, err := updateUserAction.Execute(entities.CreateEmptyAppContext(), userToUpdate)

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
	userRepositoryMock.AssertExpectations(t)
}

func TestUpdateUser_Execute_FailsIfUserRepositoryFails(t *testing.T) {
	var userRepositoryMock = new(mocks.MainDatabaseMock)
	var updateUserAction, _ = NewUpdateUserAction(userRepositoryMock)

	userRepositoryMock.On("UpdateUser", mock.Anything, mock.Anything).Return(entities.User{}, errors.New("error"))

	_, err := updateUserAction.Execute(entities.CreateEmptyAppContext(), domain.CreateUser())

	assert.NotNil(t, err)
	userRepositoryMock.AssertExpectations(t)
}

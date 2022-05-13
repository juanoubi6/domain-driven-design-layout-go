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

func TestCreateUser_Execute_Success(t *testing.T) {
	var userRepositoryMock = new(mocks.MainDatabaseMock)
	var createUserAction, _ = NewCreateUserAction(userRepositoryMock)

	var prototype = domain.CreateUserPrototype()
	var expected = domain.CreateUser()

	userRepositoryMock.On("CreateUser", mock.Anything, prototype).Return(expected, nil)

	result, err := createUserAction.Execute(entities.CreateEmptyAppContext(), prototype)

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
	userRepositoryMock.AssertExpectations(t)
}

func TestCreateUser_Execute_FailsIfUserRepositoryFails(t *testing.T) {
	var userRepositoryMock = new(mocks.MainDatabaseMock)
	var createUserAction, _ = NewCreateUserAction(userRepositoryMock)

	var prototype = domain.CreateUserPrototype()

	userRepositoryMock.On("CreateUser", mock.Anything, prototype).Return(entities.User{}, errors.New("error"))

	_, err := createUserAction.Execute(entities.CreateEmptyAppContext(), prototype)

	assert.NotNil(t, err)
	userRepositoryMock.AssertExpectations(t)

}

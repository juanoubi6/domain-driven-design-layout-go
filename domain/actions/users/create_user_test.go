package users

import (
	"domain-driven-design-layout/domain"
	"domain-driven-design-layout/domain/entities"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUser_Execute_Success(t *testing.T) {
	var userRepositoryMock = new(domain.MainDatabaseMock)
	var createUserAction, _ = NewCreateUserAction(userRepositoryMock)

	var prototype = domain.CreateUserPrototype()
	var expected = domain.CreateUser()

	userRepositoryMock.On("CreateUser", prototype).Return(expected, nil)

	result, err := createUserAction.Execute(prototype)

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
	userRepositoryMock.AssertExpectations(t)
}

func TestCreateUser_Execute_FailsIfUserRepositoryFails(t *testing.T) {
	var userRepositoryMock = new(domain.MainDatabaseMock)
	var createUserAction, _ = NewCreateUserAction(userRepositoryMock)

	var prototype = domain.CreateUserPrototype()

	userRepositoryMock.On("CreateUser", prototype).Return(entities.User{}, errors.New("error"))

	_, err := createUserAction.Execute(prototype)

	assert.NotNil(t, err)
	userRepositoryMock.AssertExpectations(t)

}

package users

import (
	"domain-driven-design-layout/domain"
	"domain-driven-design-layout/domain/entities"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestFindUserById_Execute_Success(t *testing.T) {
	var userRepositoryMock = new(domain.UserRepositoryMock)
	var findUserByIdAction, _ = NewFindUserByIdAction(userRepositoryMock)

	var userId int64 = 1
	var expected = domain.CreateUser()

	userRepositoryMock.On("GetUser", userId).Return(expected, nil)

	result, err := findUserByIdAction.Execute(userId)

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
	userRepositoryMock.AssertExpectations(t)
}

func TestFindUserById_Execute_FailsIfUserRepositoryFails(t *testing.T) {
	var userRepositoryMock = new(domain.UserRepositoryMock)
	var findUserByIdAction, _ = NewFindUserByIdAction(userRepositoryMock)

	userRepositoryMock.On("GetUser", mock.Anything).Return(entities.User{}, errors.New("error"))

	_, err := findUserByIdAction.Execute(1)

	assert.NotNil(t, err)
	userRepositoryMock.AssertExpectations(t)
}

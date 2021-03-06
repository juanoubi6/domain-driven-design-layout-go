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

func TestFindUsersByIdList_Execute_Success(t *testing.T) {
	var userRepositoryMock = new(mocks.MainDatabaseMock)
	var findUsersByIdListAction, _ = NewFindUsersByIdListAction(userRepositoryMock)

	var userIds = []int64{1, 2}
	var expected = []entities.User{domain.CreateUser(), domain.CreateUser()}

	userRepositoryMock.On("GetUsers", mock.Anything, userIds).Return(expected, nil)

	result, err := findUsersByIdListAction.Execute(entities.CreateEmptyAppContext(), userIds)

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
	userRepositoryMock.AssertExpectations(t)
}

func TestFindUsersByIdList_Execute_FailsIfUserRepositoryFails(t *testing.T) {
	var userRepositoryMock = new(mocks.MainDatabaseMock)
	var findUsersByIdListAction, _ = NewFindUsersByIdListAction(userRepositoryMock)

	userRepositoryMock.On("GetUsers", mock.Anything, mock.Anything).Return([]entities.User{}, errors.New("error"))

	_, err := findUsersByIdListAction.Execute(entities.CreateEmptyAppContext(), []int64{1, 2})

	assert.NotNil(t, err)
	userRepositoryMock.AssertExpectations(t)
}

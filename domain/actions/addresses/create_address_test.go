package addresses

import (
	"domain-driven-design-layout/domain"
	"domain-driven-design-layout/domain/entities"
	"domain-driven-design-layout/domain/mocks"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateAddress_Execute_Success(t *testing.T) {
	var userRepositoryMock = new(mocks.MainDatabaseMock)
	var addressRepositoryMock = new(mocks.MainDatabaseMock)
	var createAddressAction, _ = NewCreateAddressAction(addressRepositoryMock, userRepositoryMock)

	var prototype = domain.CreateAddressPrototype()
	var expected = domain.CreateAddress()

	var userID int64 = 1
	var user = domain.CreateUser()

	userRepositoryMock.On("GetUser", mock.Anything, userID).Return(&user, nil)
	addressRepositoryMock.On("CreateAddress", mock.Anything, userID, prototype).Return(expected, nil)

	result, err := createAddressAction.Execute(entities.CreateEmptyAppContext(), userID, prototype)

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
	userRepositoryMock.AssertExpectations(t)
	addressRepositoryMock.AssertExpectations(t)
}

func TestCreateAddress_Execute_FailsIfAddressRepositoryFails(t *testing.T) {
	var userRepositoryMock = new(mocks.MainDatabaseMock)
	var addressRepositoryMock = new(mocks.MainDatabaseMock)
	var createAddressAction, _ = NewCreateAddressAction(addressRepositoryMock, userRepositoryMock)

	var prototype = domain.CreateAddressPrototype()
	var userID int64 = 1
	var user = domain.CreateUser()

	userRepositoryMock.On("GetUser", mock.Anything, userID).Return(&user, nil)
	addressRepositoryMock.On("CreateAddress", mock.Anything, userID, prototype).Return(entities.Address{}, errors.New("err"))

	_, err := createAddressAction.Execute(entities.CreateEmptyAppContext(), userID, prototype)

	assert.NotNil(t, err)
	userRepositoryMock.AssertExpectations(t)
	addressRepositoryMock.AssertExpectations(t)
}

func TestCreateAddress_Execute_FailsIfUserDoesNotExist(t *testing.T) {
	var userRepositoryMock = new(mocks.MainDatabaseMock)
	var addressRepositoryMock = new(mocks.MainDatabaseMock)
	var createAddressAction, _ = NewCreateAddressAction(addressRepositoryMock, userRepositoryMock)

	var prototype = domain.CreateAddressPrototype()
	var userID int64 = 1

	userRepositoryMock.On("GetUser", mock.Anything, userID).Return(nil, nil)

	_, err := createAddressAction.Execute(entities.CreateEmptyAppContext(), userID, prototype)

	assert.NotNil(t, err)
	userRepositoryMock.AssertExpectations(t)
}

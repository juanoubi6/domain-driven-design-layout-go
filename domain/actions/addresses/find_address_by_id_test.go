package addresses

import (
	"domain-driven-design-layout/domain"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestFindAddressById_Execute_Success(t *testing.T) {
	var addressRepositoryMock = new(domain.MainDatabaseMock)
	var findAddressByIdAction, _ = NewFindAddressByIdAction(addressRepositoryMock)

	var addressId int64 = 1
	var expected = domain.CreateAddress()

	addressRepositoryMock.On("GetAddress", addressId).Return(&expected, nil)

	result, err := findAddressByIdAction.Execute(addressId)

	assert.Nil(t, err)
	assert.Equal(t, &expected, result)
	addressRepositoryMock.AssertExpectations(t)
}

func TestFindAddressById_Execute_FailsIfAddressRepositoryFails(t *testing.T) {
	var addressRepositoryMock = new(domain.MainDatabaseMock)
	var findAddressByIdAction, _ = NewFindAddressByIdAction(addressRepositoryMock)

	addressRepositoryMock.On("GetAddress", mock.Anything).Return(nil, errors.New("error"))

	_, err := findAddressByIdAction.Execute(1)

	assert.NotNil(t, err)
	addressRepositoryMock.AssertExpectations(t)
}

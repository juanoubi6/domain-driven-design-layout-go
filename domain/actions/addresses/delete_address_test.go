package addresses

import (
	"domain-driven-design-layout/domain"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteAddress_Execute_Success(t *testing.T) {
	var addressRepositoryMock = new(domain.MainDatabaseMock)
	var deleteAddressAction, _ = NewDeleteAddressAction(addressRepositoryMock)

	var addressID int64 = 1

	addressRepositoryMock.On("DeleteAddress", addressID).Return(nil)

	err := deleteAddressAction.Execute(addressID)

	assert.Nil(t, err)
	addressRepositoryMock.AssertExpectations(t)
}

func TestDeleteAddress_Execute_FailsIfAddressRepositoryFails(t *testing.T) {
	var addressRepositoryMock = new(domain.MainDatabaseMock)
	var deleteAddressAction, _ = NewDeleteAddressAction(addressRepositoryMock)

	addressRepositoryMock.On("DeleteAddress", int64(1)).Return(errors.New("error"))

	err := deleteAddressAction.Execute(int64(1))

	assert.NotNil(t, err)
	addressRepositoryMock.AssertExpectations(t)
}

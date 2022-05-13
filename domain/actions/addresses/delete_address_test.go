package addresses

import (
	"domain-driven-design-layout/domain/entities"
	"domain-driven-design-layout/domain/mocks"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestDeleteAddress_Execute_Success(t *testing.T) {
	var addressRepositoryMock = new(mocks.MainDatabaseMock)
	var deleteAddressAction, _ = NewDeleteAddressAction(addressRepositoryMock)

	var addressID int64 = 1

	addressRepositoryMock.On("DeleteAddress", mock.Anything, addressID).Return(nil)

	err := deleteAddressAction.Execute(entities.CreateEmptyAppContext(), addressID)

	assert.Nil(t, err)
	addressRepositoryMock.AssertExpectations(t)
}

func TestDeleteAddress_Execute_FailsIfAddressRepositoryFails(t *testing.T) {
	var addressRepositoryMock = new(mocks.MainDatabaseMock)
	var deleteAddressAction, _ = NewDeleteAddressAction(addressRepositoryMock)

	addressRepositoryMock.On("DeleteAddress", mock.Anything, int64(1)).Return(errors.New("error"))

	err := deleteAddressAction.Execute(entities.CreateEmptyAppContext(), int64(1))

	assert.NotNil(t, err)
	addressRepositoryMock.AssertExpectations(t)
}

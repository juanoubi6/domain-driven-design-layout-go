package repositories

import (
	"domain-driven-design-layout/domain/entities"
	"domain-driven-design-layout/infrastructure/repositories/sql/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AddressRepositoryTestSuite struct {
	suite.Suite
	addressRepository *AddressRepository
}

func (suite *AddressRepositoryTestSuite) SetupTest() {
	suite.addressRepository = &AddressRepository{db: db}
	generateSchema()
}

func TestAddressRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(AddressRepositoryTestSuite))
}

func (suite *AddressRepositoryTestSuite) TestAddressRepository_CreateAddress_SuccessfullyReturnsCreatedAddress() {
	var userID int64 = 1

	saveUserWithAddresses(userID)

	country := "Argentina"
	prototype := entities.AddressPrototype{
		Street: "New street 1",
		Number: 10,
		City:   &country,
	}

	address, err := suite.addressRepository.CreateAddress(userID, prototype)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}

	assert.Equal(suite.T(), "New street 1", address.Street)
	assert.Equal(suite.T(), "Argentina", *address.City)
}

func (suite *AddressRepositoryTestSuite) TestAddressRepository_DeleteAddress_SuccessfullyDeletesAddress() {
	var userId int64 = 10
	saveUserWithAddresses(userId)

	savedAddress := getSingleAddressFromUser(userId)

	err := suite.addressRepository.DeleteAddress(savedAddress.ID)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}

	assert.Nil(suite.T(), err)

	deletedAddress := getAddressById(savedAddress.ID)

	assert.Equal(suite.T(), int64(0), deletedAddress.ID)
}

func getSingleAddressFromUser(userId int64) models.AddressModel {
	addressModel := models.AddressModel{}
	_ = db.Get(&addressModel, "SELECT * FROM addresses WHERE user_id=$1 LIMIT 1", userId)

	return addressModel
}

func getAddressById(addressId int64) models.AddressModel {
	addressModel := models.AddressModel{}
	_ = db.Get(&addressModel, "SELECT * FROM addresses WHERE id=$1", addressId)

	return addressModel
}

package repositories

import (
	"domain-driven-design-layout/domain/entities"
	"domain-driven-design-layout/infrastructure/repositories/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

type AddressRepositoryTestSuite struct {
	suite.Suite
	addressRepository *AddressRepository
	sqlMock           sqlmock.Sqlmock
}

func (suite *AddressRepositoryTestSuite) SetupTest() {
	mockDb, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	suite.addressRepository = &AddressRepository{queryExecutor: QueryExecutor{db: sqlx.NewDb(mockDb, "postgres"), tx: nil}}
	suite.sqlMock = mock
}

func TestAddressRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(AddressRepositoryTestSuite))
}

func (suite *AddressRepositoryTestSuite) TestAddressRepository_CreateAddress_SuccessfullyReturnsCreatedAddress() {
	var userID int64 = 1

	country := "Argentina"
	prototype := entities.AddressPrototype{
		Street: "New street 1",
		Number: 10,
		City:   &country,
	}

	suite.sqlMock.ExpectQuery(sql.InsertAddress).WithArgs(userID, prototype.Street, prototype.Number, prototype.City).WillReturnRows(
		sqlmock.NewRows([]string{"id"}).AddRow(99),
	)

	address, err := suite.addressRepository.CreateAddress(userID, prototype)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}

	assert.Equal(suite.T(), int64(99), address.ID)
	assert.Equal(suite.T(), "New street 1", address.Street)
	assert.Equal(suite.T(), "Argentina", *address.City)
	if err := suite.sqlMock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

func (suite *AddressRepositoryTestSuite) TestAddressRepository_DeleteAddress_SuccessfullyDeletesAddress() {
	var addressId int64 = 10

	suite.sqlMock.ExpectExec(sql.DeleteAddress).WithArgs(addressId).WillReturnResult(sqlmock.NewResult(0, 1))

	err := suite.addressRepository.DeleteAddress(addressId)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}

	assert.Nil(suite.T(), err)
	if err := suite.sqlMock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

func (suite *AddressRepositoryTestSuite) TestAddressRepository_GetAddress_SuccessfullyReturnsAddress() {
	addressId := int64(15)

	suite.sqlMock.ExpectQuery(sql.GetAddressById).WithArgs(addressId).WillReturnRows(
		sqlmock.NewRows(
			[]string{"id", "user_id", "street", "number", "city"},
		).AddRow(
			addressId, 1, "Some street", 1, nil,
		),
	)

	address, err := suite.addressRepository.GetAddress(addressId)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}

	assert.Equal(suite.T(), addressId, address.ID)
	assert.Nil(suite.T(), address.City)
	if err := suite.sqlMock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

func (suite *AddressRepositoryTestSuite) TestAddressRepository_DeleteUserAddresses_SuccessfullyDeletesAddresses() {
	var userID int64 = 10

	suite.sqlMock.ExpectExec(sql.DeleteUserAddresses).WithArgs(userID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := suite.addressRepository.DeleteUserAddresses(userID)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}

	assert.Nil(suite.T(), err)
	if err := suite.sqlMock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

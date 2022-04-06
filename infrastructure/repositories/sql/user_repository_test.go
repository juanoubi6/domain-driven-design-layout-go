package sql

import (
	"domain-driven-design-layout/domain/entities"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
	"time"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	userRepository *QueryExecutor
	sqlMock        sqlmock.Sqlmock
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	mockDb, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	suite.userRepository = &QueryExecutor{db: sqlx.NewDb(mockDb, "postgres"), tx: nil}
	suite.sqlMock = mock
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

func (suite *UserRepositoryTestSuite) TestUserRepository_CreateUser_SuccessfullyReturnsCreatedUser() {
	country := "Argentina"
	prototype := entities.UserPrototype{
		FirstName: "test",
		LastName:  "name",
		BirthDate: time.Now(),
		AddressesPrototypes: []entities.AddressPrototype{
			{
				Street: "street 1",
				Number: 1,
				City:   &country,
			},
			{
				Street: "street 2",
				Number: 2,
				City:   nil,
			},
		},
	}

	expectedAddressInsertQuery := `INSERT INTO addresses (user_id, street, number, city) VALUES ($1, $2, $3, $4),($5, $6, $7, $8)`

	suite.sqlMock.ExpectBegin()
	suite.sqlMock.ExpectQuery(InsertUser).WithArgs(prototype.FirstName, prototype.LastName, prototype.BirthDate).WillReturnRows(
		sqlmock.NewRows([]string{"id"}).AddRow(99),
	)
	suite.sqlMock.ExpectExec(expectedAddressInsertQuery).WithArgs(
		99,
		prototype.AddressesPrototypes[0].Street,
		prototype.AddressesPrototypes[0].Number,
		prototype.AddressesPrototypes[0].City,
		99,
		prototype.AddressesPrototypes[1].Street,
		prototype.AddressesPrototypes[1].Number,
		prototype.AddressesPrototypes[1].City,
	).WillReturnResult(
		sqlmock.NewResult(0, 2),
	)
	suite.sqlMock.ExpectCommit()
	suite.sqlMock.ExpectQuery(GetUserWithAddressesById).WithArgs(99).WillReturnRows(
		sqlmock.NewRows(
			[]string{"id", "first_name", "last_name", "birth_date", "id", "user_id", "street", "number", "city"},
		).AddRow(
			99, "firstName", "lastName", time.Now(), 1, 99, "street 1", 1, country,
		).AddRow(
			99, "firstName", "lastName", time.Now(), 2, 99, "street 2", 11, nil,
		),
	)

	user, err := suite.userRepository.CreateUser(prototype)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}

	assert.Equal(suite.T(), "firstName", user.FirstName)
	assert.Equal(suite.T(), 2, len(user.Addresses))
	assert.Equal(suite.T(), user.Addresses[0].UserID, user.ID)
	if err = suite.sqlMock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

func (suite *UserRepositoryTestSuite) TestUserRepository_CreateUser_RollbacksTransactionOnInvalidAddress() {
	country := "Argentina"
	prototype := entities.UserPrototype{
		FirstName: "test",
		LastName:  "name",
		BirthDate: time.Now(),
		AddressesPrototypes: []entities.AddressPrototype{
			{
				Street: "street 1",
				Number: 1,
				City:   &country,
			},
			{
				Street: "street 2 exceeding max length ssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss",
				Number: 2,
				City:   nil,
			},
		},
	}

	expectedAddressInsertQuery := `INSERT INTO addresses (user_id, street, number, city) VALUES ($1, $2, $3, $4),($5, $6, $7, $8)`

	suite.sqlMock.ExpectBegin()
	suite.sqlMock.ExpectQuery(InsertUser).WithArgs(prototype.FirstName, prototype.LastName, prototype.BirthDate).WillReturnRows(
		sqlmock.NewRows([]string{"id"}).AddRow(99),
	)
	suite.sqlMock.ExpectExec(expectedAddressInsertQuery).WithArgs(
		99,
		prototype.AddressesPrototypes[0].Street,
		prototype.AddressesPrototypes[0].Number,
		prototype.AddressesPrototypes[0].City,
		99,
		prototype.AddressesPrototypes[1].Street,
		prototype.AddressesPrototypes[1].Number,
		prototype.AddressesPrototypes[1].City,
	).WillReturnError(
		errors.New("some error"),
	)
	suite.sqlMock.ExpectRollback()

	_, err := suite.userRepository.CreateUser(prototype)
	if err == nil {
		assert.FailNow(suite.T(), "method should have failed")
	}

	if err = suite.sqlMock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

func (suite *UserRepositoryTestSuite) TestUserRepository_GetUser_SuccessfullyReturnsUser() {
	var userId int64 = 10

	suite.sqlMock.ExpectQuery(GetUserWithAddressesById).WithArgs(userId).WillReturnRows(
		sqlmock.NewRows(
			[]string{"id", "first_name", "last_name", "birth_date", "id", "user_id", "street", "number", "city"},
		).AddRow(
			userId, "firstName", "lastName", time.Now(), 1, userId, "street 1", 1, "Argentina",
		).AddRow(
			userId, "firstName", "lastName", time.Now(), 2, userId, "street 2", 11, nil,
		),
	)

	user, err := suite.userRepository.GetUser(userId)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}
	if user == nil {
		assert.FailNow(suite.T(), "User should not be nil")
	}

	assert.Equal(suite.T(), userId, user.ID)
	assert.Equal(suite.T(), 2, len(user.Addresses))
	if err = suite.sqlMock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

func (suite *UserRepositoryTestSuite) TestUserRepository_GetUser_ReturnsNilWhenUserCouldNotBeFound() {
	var userId int64 = 999

	suite.sqlMock.ExpectQuery(GetUserWithAddressesById).WithArgs(userId).WillReturnRows(sqlmock.NewRows([]string{}))

	user, err := suite.userRepository.GetUser(999)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}

	assert.Nil(suite.T(), nil, user)
	if err = suite.sqlMock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

func (suite *UserRepositoryTestSuite) TestUserRepository_GetUsers_SuccessfullyReturnsUsers() {
	var userId1 int64 = 10
	var userId2 int64 = 20

	expectedQuery := `SELECT 
		u.id, 
		u.first_name, 
		u.last_name, 
		u.birth_date,
		a.id, 
		a.user_id, 
		a.street, 
		a.number, 
		a.city
	FROM users u
	LEFT JOIN addresses a ON u.id = a.user_id
	WHERE u.id IN ($1, $2)`

	suite.sqlMock.ExpectQuery(expectedQuery).WithArgs(userId1, userId2).WillReturnRows(
		sqlmock.NewRows(
			[]string{"id", "first_name", "last_name", "birth_date", "id", "user_id", "street", "number", "city"},
		).AddRow(
			userId1, "firstName", "lastName", time.Now(), 1, userId1, "street 1", 1, "Argentina",
		).AddRow(
			userId1, "firstName", "lastName", time.Now(), 2, userId1, "street 2", 11, nil,
		).AddRow(
			userId2, "firstName", "lastName", time.Now(), 3, userId2, "street 3", 11, nil,
		),
	)

	users, err := suite.userRepository.GetUsers([]int64{userId1, userId2})
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}
	if users == nil {
		assert.FailNow(suite.T(), "Users should not be nil")
	}

	assert.Equal(suite.T(), 2, len(users))
	assert.Equal(suite.T(), 2, len(users[0].Addresses))
	assert.Equal(suite.T(), 1, len(users[1].Addresses))
	if err = suite.sqlMock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

func (suite *UserRepositoryTestSuite) TestUserRepository_UpdateUser_SuccessfullyUpdatesUser() {
	var userId int64 = 10

	userWithUpdatedFields := entities.User{
		ID:        userId,
		FirstName: "newFirstName",
		LastName:  "newLastName",
		BirthDate: time.Now(),
		Addresses: nil,
	}

	suite.sqlMock.ExpectQuery(GetUserWithAddressesById).WithArgs(userId).WillReturnRows(
		sqlmock.NewRows(
			[]string{"id", "first_name", "last_name", "birth_date", "id", "user_id", "street", "number", "city"},
		).AddRow(
			userId, "firstName", "lastName", time.Now(), 1, userId, "street 1", 1, "Argentina",
		).AddRow(
			userId, "firstName", "lastName", time.Now(), 2, userId, "street 2", 11, nil,
		),
	)
	suite.sqlMock.ExpectPrepare(UpdateUser).ExpectExec().WithArgs(
		userWithUpdatedFields.FirstName,
		userWithUpdatedFields.LastName,
		userWithUpdatedFields.BirthDate,
		userWithUpdatedFields.ID,
	).WillReturnResult(
		sqlmock.NewResult(0, 1),
	)

	user, err := suite.userRepository.UpdateUser(userWithUpdatedFields)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}

	assert.Equal(suite.T(), "newFirstName", user.FirstName)
	if err = suite.sqlMock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

func (suite *UserRepositoryTestSuite) TestUserRepository_DeleteUser_SuccessfullyDeletesUser() {
	var userId int64 = 10

	suite.sqlMock.ExpectPrepare(DeleteUser).ExpectExec().WithArgs(userId).WillReturnResult(
		sqlmock.NewResult(0, 1),
	)

	err := suite.userRepository.DeleteUser(userId)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}

	assert.Nil(suite.T(), err)
	if err = suite.sqlMock.ExpectationsWereMet(); err != nil {
		suite.T().Errorf("there were unfulfilled expectations: %s", err)
	}
}

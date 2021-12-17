package repositories

import (
	"domain-driven-design-layout/domain/entities"
	"domain-driven-design-layout/infrastructure/config"
	"domain-driven-design-layout/infrastructure/repositories/sql"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"testing"
	"time"
)

var db = sql.CreateDatabaseConnection(config.LoadAppConfig().SQLConfig)

type UserRepositoryTestSuite struct {
	suite.Suite
	userRepository *UserRepository
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	suite.userRepository = &UserRepository{db: db}
	generateSchema()
}

func TestUserHandlersTestSuite(t *testing.T) {
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

	user, err := suite.userRepository.CreateUser(prototype)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}

	assert.Equal(suite.T(), "test", user.FirstName)
	assert.Equal(suite.T(), 2, len(user.Addresses))
	assert.Equal(suite.T(), user.Addresses[0].UserID, user.ID)
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

	_, err := suite.userRepository.CreateUser(prototype)
	if err == nil {
		assert.FailNow(suite.T(), "method should have failed")
	}

	createdUsers := 0
	err = db.QueryRow("SELECT count(*) FROM users").Scan(&createdUsers)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}

	assert.Equal(suite.T(), 0, createdUsers)
}

func (suite *UserRepositoryTestSuite) TestUserRepository_GetUser_SuccessfullyReturnsUser() {
	var userId int64 = 10
	saveUserWithAddresses(userId)

	user, err := suite.userRepository.GetUser(userId)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}
	if user == nil {
		assert.FailNow(suite.T(), "User should not be nil")
	}

	assert.Equal(suite.T(), userId, user.ID)
	assert.Equal(suite.T(), 2, len(user.Addresses))
}

func (suite *UserRepositoryTestSuite) TestUserRepository_GetUser_ReturnsNilWhenUserCouldNotBeFound() {
	user, err := suite.userRepository.GetUser(999)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}

	assert.Nil(suite.T(), nil, user)
}

func (suite *UserRepositoryTestSuite) TestUserRepository_GetUsers_SuccessfullyReturnsUsers() {
	var userId1 int64 = 10
	var userId2 int64 = 20
	saveUserWithAddresses(userId1)
	saveUserWithAddresses(userId2)

	users, err := suite.userRepository.GetUsers([]int64{userId1, userId2})
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}
	if users == nil {
		assert.FailNow(suite.T(), "Users should not be nil")
	}

	assert.Equal(suite.T(), 2, len(users))
	assert.Equal(suite.T(), 2, len(users[0].Addresses))
	assert.Equal(suite.T(), 2, len(users[1].Addresses))
}

func (suite *UserRepositoryTestSuite) TestUserRepository_UpdateUser_SuccessfullyUpdatesUser() {
	var userId int64 = 10
	saveUserWithAddresses(userId)

	userWithUpdatedFields := entities.User{
		ID:        userId,
		FirstName: "newFirstName",
		LastName:  "newLastName",
		BirthDate: time.Now(),
		Addresses: nil,
	}

	user, err := suite.userRepository.UpdateUser(userWithUpdatedFields)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}

	assert.Equal(suite.T(), "newFirstName", user.FirstName)

	updatedUser, err := suite.userRepository.GetUser(userId)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}

	assert.Equal(suite.T(), "newFirstName", updatedUser.FirstName)
}

func (suite *UserRepositoryTestSuite) TestUserRepository_DeleteUser_SuccessfullyDeletesUser() {
	var userId int64 = 10
	saveUserWithAddresses(userId)

	err := suite.userRepository.DeleteUser(userId)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}

	assert.Nil(suite.T(), err)

	deletedUser, err := suite.userRepository.GetUser(userId)
	if err != nil {
		assert.FailNow(suite.T(), err.Error())
	}

	assert.Nil(suite.T(), deletedUser)
}

func generateSchema() {
	content, err := ioutil.ReadFile("../../schema.sql")
	if err != nil {
		panic("Could not read schema file")
	}

	_, err = db.Exec(string(content))
	if err != nil {
		panic("Could not execute schema.sql file")
	}
}

func saveUserWithAddresses(userId int64) {
	insertUsersQuery := fmt.Sprintf(
		`INSERT INTO users (id, first_name, last_name, birth_date) VALUES 
				(%v,'test', 'user', '1995-07-20T00:00:00.000Z')`,
		userId,
	)

	insertAddressesQuery := fmt.Sprintf(
		`INSERT INTO addresses (street, number, user_id, city) VALUES 
			('Street 1', 1, %v, NULL), 
			('Street 2', 2, %v, 'Argentina')`,
		userId, userId,
	)

	_, _ = db.Exec(insertUsersQuery)
	_, _ = db.Exec(insertAddressesQuery)

}

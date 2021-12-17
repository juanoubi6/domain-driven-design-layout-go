package repositories

import (
	"context"
	"domain-driven-design-layout/domain/entities"
	"domain-driven-design-layout/infrastructure/config"
	"domain-driven-design-layout/infrastructure/repositories/sql"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

var connectionPool = sql.CreateConnectionPool(config.LoadAppConfig().SQLConfig)

type UserRepositoryTestSuite struct {
	suite.Suite
	userRepository *UserRepository
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	suite.userRepository = &UserRepository{connectionPool: connectionPool}
	truncateTables(connectionPool)
}

func TestUserHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

func (suite *UserRepositoryTestSuite) TestUserRepository_CreateUser_SuccessReturnsCreatedUser() {
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
		assert.Fail(suite.T(), err.Error())
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
		assert.Fail(suite.T(), "method should have failed")
	}

	createdUsers := 0
	err = connectionPool.QueryRow(context.Background(), "SELECT count(*) FROM users").Scan(&createdUsers)
	if err != nil {
		assert.Fail(suite.T(), err.Error())
	}

	assert.Equal(suite.T(), 0, createdUsers)
}

func truncateTables(pool *pgxpool.Pool) {
	_, _ = pool.Exec(context.Background(), "TRUNCATE TABLE users CASCADE")
	_, _ = pool.Exec(context.Background(), "TRUNCATE TABLE addresses CASCADE")
}

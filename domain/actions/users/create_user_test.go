package users

import (
	"domain-driven-design-layout/domain"
	"domain-driven-design-layout/domain/entities"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var userRepositoryMock = new(domain.UserRepositoryMock)
var createUserAction, _ = NewCreateUserAction(userRepositoryMock)

func TestCreateUser_Execute_Success(t *testing.T) {
	var prototype = createUserPrototype()
	var expected = createUser()

	userRepositoryMock.On("CreateUser", prototype).Return(expected, nil)

	result, err := createUserAction.Execute(prototype)

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
	userRepositoryMock.AssertExpectations(t)
}

func TestCreateUser_Execute_FailsIfUserRepositoryFails(t *testing.T) {
	var prototype = createUserPrototype()

	userRepositoryMock.On("CreateUser", prototype).Return(entities.User{}, errors.New("error"))

	_, err := createUserAction.Execute(prototype)

	assert.NotNil(t, err)
	userRepositoryMock.AssertExpectations(t)

}

func createUserPrototype() entities.UserPrototype {
	country := "Argentina"

	return entities.UserPrototype{
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
				Street: "street 1",
				Number: 1,
				City:   nil,
			},
		},
	}
}

func createUser() entities.User {
	country := "Argentina"

	return entities.User{
		ID:        1,
		FirstName: "test",
		LastName:  "name",
		BirthDate: time.Now(),
		Addresses: []entities.Address{
			{
				ID:     1,
				UserID: 1,
				Street: "street 1",
				Number: 1,
				City:   &country,
			},
			{
				ID:     2,
				UserID: 2,
				Street: "street 1",
				Number: 1,
				City:   nil,
			},
		},
	}
}

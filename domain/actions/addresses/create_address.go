package users

import (
	"domain-driven-design-layout/domain/entities"
	"errors"
	"fmt"
)

type CreateAddress interface {
	Execute(userID int64, prototype entities.AddressPrototype) (entities.Address, error)
}

type CreateAddressAction struct {
	addressRepository entities.AddressRepository
	userRepository    entities.UserRepository
}

func NewCreateAddressAction(
	addressRepository entities.AddressRepository,
	userRepository entities.UserRepository,
) (CreateAddress, error) {
	result := CreateAddressAction{
		addressRepository: addressRepository,
		userRepository:    userRepository,
	}

	return &result, nil
}

func (act *CreateAddressAction) Execute(userID int64, prototype entities.AddressPrototype) (entities.Address, error) {
	//Execute any business logic or validations you need, for example, validating user existence
	user, err := act.userRepository.GetUser(userID)
	if err != nil {
		return entities.Address{}, fmt.Errorf("user could not be found. Error: %v", err)
	}

	if user == nil {
		return entities.Address{}, errors.New("user does not exist")
	}

	createdAddress, err := act.addressRepository.CreateAddress(userID, prototype)
	if err != nil {
		return entities.Address{}, fmt.Errorf("could not create address. Error: %v", err)
	}

	return createdAddress, nil
}

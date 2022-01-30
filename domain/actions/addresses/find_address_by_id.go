package addresses

import (
	"domain-driven-design-layout/domain/entities"
	"fmt"
)

type FindAddressById interface {
	Execute(int64) (*entities.Address, error)
}

type FindAddressByIdAction struct {
	addressRepository entities.AddressRepository
}

func NewFindAddressByIdAction(repository entities.AddressRepository) (FindAddressById, error) {
	result := FindAddressByIdAction{
		addressRepository: repository,
	}

	return &result, nil
}

func (act *FindAddressByIdAction) Execute(id int64) (*entities.Address, error) {
	//Execute any business logic or validations you need
	address, err := act.addressRepository.GetAddress(id)
	if err != nil {
		return nil, fmt.Errorf("user could not be found. Error: %v", err)
	}

	return address, nil
}

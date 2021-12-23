package addresses

import (
	"domain-driven-design-layout/domain/entities"
)

type DeleteAddress interface {
	Execute(int64) error
}

type DeleteAddressAction struct {
	addressRepository entities.AddressRepository
}

func NewDeleteAddressAction(repository entities.AddressRepository) (DeleteAddress, error) {
	result := DeleteAddressAction{
		addressRepository: repository,
	}

	return &result, nil
}

func (act *DeleteAddressAction) Execute(id int64) error {
	//Execute any business logic or validations you need
	if err := act.addressRepository.DeleteAddress(id); err != nil {
		return err
	}

	return nil
}

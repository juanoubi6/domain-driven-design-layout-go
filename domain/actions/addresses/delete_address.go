package addresses

import (
	"domain-driven-design-layout/domain/entities"
)

type DeleteAddress interface {
	Execute(ctx entities.ApplicationContext, userID int64) error
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

func (act *DeleteAddressAction) Execute(ctx entities.ApplicationContext, id int64) error {
	//Execute any business logic or validations you need
	if err := act.addressRepository.DeleteAddress(ctx, id); err != nil {
		return err
	}

	return nil
}

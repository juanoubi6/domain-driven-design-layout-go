package builder

import (
	"domain-driven-design-layout/domain/actions/addresses"
	"domain-driven-design-layout/domain/actions/users"
	"domain-driven-design-layout/domain/entities"
)

type Actions struct {
	CreateUser        users.CreateUser
	FindUserById      users.FindUserById
	FindUsersByIdList users.FindUsersByIdList
	UpdateUser        users.UpdateUser
	DeleteUser        users.DeleteUser
	CreateAddress     addresses.CreateAddress
	DeleteAddress     addresses.DeleteAddress
	FindAddressById   addresses.FindAddressById
}

func CreateActions(repositories *Repositories, txRepositoryCreator entities.TxRepositoryCreator) (*Actions, error) {
	createUser, err := users.NewCreateUserAction(repositories.UserRepository)
	if err != nil {
		return nil, err
	}

	findUserById, err := users.NewFindUserByIdAction(repositories.UserRepository)
	if err != nil {
		return nil, err
	}

	findUsersByIdList, err := users.NewFindUsersByIdListAction(repositories.UserRepository)
	if err != nil {
		return nil, err
	}

	updateUser, err := users.NewUpdateUserAction(repositories.UserRepository)
	if err != nil {
		return nil, err
	}

	deleteUser, err := users.NewDeleteUserAction(txRepositoryCreator)
	if err != nil {
		return nil, err
	}

	createAddress, err := addresses.NewCreateAddressAction(repositories.AddressRepository, repositories.UserRepository)
	if err != nil {
		return nil, err
	}

	deleteAddress, err := addresses.NewDeleteAddressAction(repositories.AddressRepository)
	if err != nil {
		return nil, err
	}

	findAddressById, err := addresses.NewFindAddressByIdAction(repositories.AddressRepository)
	if err != nil {
		return nil, err
	}

	return &Actions{
		CreateUser:        createUser,
		FindUserById:      findUserById,
		FindUsersByIdList: findUsersByIdList,
		UpdateUser:        updateUser,
		DeleteUser:        deleteUser,
		CreateAddress:     createAddress,
		DeleteAddress:     deleteAddress,
		FindAddressById:   findAddressById,
	}, nil
}

package builder

import (
	"domain-driven-design-layout/domain/entities"
	"domain-driven-design-layout/infrastructure/config"
	"domain-driven-design-layout/infrastructure/repositories"
)

type Repositories struct {
	UserRepository    entities.UserRepository
	AddressRepository entities.AddressRepository
}

func CreateRepositories(config config.RepositoriesConfig) (*Repositories, error) {
	userRepository, err := repositories.NewUserRepository(config.SQLConfig)
	if err != nil {
		return nil, err
	}

	addressRepository, err := repositories.NewAddressRepository(config.SQLConfig)
	if err != nil {
		return nil, err
	}

	return &Repositories{
		UserRepository:    userRepository,
		AddressRepository: addressRepository,
	}, nil
}

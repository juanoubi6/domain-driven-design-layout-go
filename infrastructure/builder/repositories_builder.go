package builder

import (
	"domain-driven-design-layout/domain/entities"
	"domain-driven-design-layout/infrastructure/config"
	"domain-driven-design-layout/infrastructure/repositories"
	"domain-driven-design-layout/infrastructure/repositories/sql"
)

type Repositories struct {
	UserRepository    entities.UserRepository
	AddressRepository entities.AddressRepository
}

func CreateRepositories(config config.RepositoriesConfig) (*Repositories, error) {
	db := sql.CreateDatabaseConnection(config.SQLConfig)

	userRepository, err := repositories.NewUserRepository(db)
	if err != nil {
		return nil, err
	}

	addressRepository, err := repositories.NewAddressRepository(db)
	if err != nil {
		return nil, err
	}

	return &Repositories{
		UserRepository:    userRepository,
		AddressRepository: addressRepository,
	}, nil
}

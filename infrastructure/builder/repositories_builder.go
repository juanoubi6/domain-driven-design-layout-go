package builder

import (
	"domain-driven-design-layout/domain/entities"
	"domain-driven-design-layout/infrastructure/config"
	"domain-driven-design-layout/infrastructure/repositories/sql"
)

type Repositories struct {
	UserRepository    entities.UserRepository
	AddressRepository entities.AddressRepository
}

func CreateRepositories(config config.RepositoriesConfig) (*Repositories, error) {
	db := sql.CreateDatabaseConnection(config.SQLConfig)

	mainDatabase := sql.CreateQueryExecutor(db, nil)

	return &Repositories{
		UserRepository:    mainDatabase,
		AddressRepository: mainDatabase,
	}, nil
}

func CreateTxRepositoryFactory(config config.RepositoriesConfig) entities.TxRepositoryCreator {
	db := sql.CreateDatabaseConnection(config.SQLConfig)

	return sql.NewTxRepositoryFactory(db)
}

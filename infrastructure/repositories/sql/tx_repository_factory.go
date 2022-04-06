package sql

import (
	"domain-driven-design-layout/domain/entities"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type TxRepositoryFactory struct {
	db *sqlx.DB
}

func NewTxRepositoryFactory(db *sqlx.DB) entities.TxRepositoryCreator {
	return &TxRepositoryFactory{db: db}
}

func (txRF *TxRepositoryFactory) CreateMainDatabase() (entities.MainDatabase, error) {
	newTx, err := txRF.db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("unable to create transaction: %w", err)
	}

	return &QueryExecutor{db: txRF.db, tx: newTx}, nil
}

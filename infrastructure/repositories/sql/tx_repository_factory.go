package sql

import (
	"context"
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

func (txRF *TxRepositoryFactory) CreateTxMainDatabase(ctx context.Context) (entities.MainDatabase, error) {
	newTx, err := txRF.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create transaction for mainDatabase: %w", err)
	}

	return &QueryExecutor{db: txRF.db, tx: newTx}, nil
}

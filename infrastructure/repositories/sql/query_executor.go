package sql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type QueryExecutor struct {
	db *sqlx.DB
	tx *sqlx.Tx
}

func CreateQueryExecutor(db *sqlx.DB, tx *sqlx.Tx) *QueryExecutor {
	return &QueryExecutor{db: db, tx: tx}
}

func (qe *QueryExecutor) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	var err error
	var res sql.Result
	var prepStmt *sqlx.Stmt

	if qe.tx != nil {
		prepStmt, err = qe.tx.PreparexContext(ctx, query)
	} else {
		prepStmt, err = qe.db.PreparexContext(ctx, query)
	}

	if err != nil {
		return nil, fmt.Errorf("error preparing statement: %w", err)
	}

	res, err = prepStmt.ExecContext(ctx, args...)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}

	return res, nil
}

func (qe *QueryExecutor) CommitTx() error {
	if qe.tx == nil {
		return nil
	}

	if err := qe.tx.Commit(); err != nil {
		return fmt.Errorf("tx commit failed: %w", err)
	}

	qe.tx = nil

	return nil
}

func (qe *QueryExecutor) RollbackTx() error {
	if qe.tx == nil {
		return nil
	}

	if err := qe.tx.Rollback(); err != nil {
		return fmt.Errorf("tx rollback failed: %w", err)
	}

	qe.tx = nil

	return nil
}

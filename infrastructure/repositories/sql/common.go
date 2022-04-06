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

	if qe.tx != nil {
		res, err = qe.tx.ExecContext(ctx, query, args...)
	} else {
		res, err = qe.db.ExecContext(ctx, query, args...)
	}

	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}

	return res, nil
}

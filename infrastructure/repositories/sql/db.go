package sql

import (
	"context"
	"domain-driven-design-layout/infrastructure/config"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

const connStringTemplate = "host=%s user=%s dbname=%s password=%s port=%d sslmode=disable"

func CreateConnectionPool(config config.SQLConfig) *pgxpool.Pool {
	uri := fmt.Sprintf(
		connStringTemplate,
		config.Host,
		config.UserName,
		config.DbName,
		config.Password,
		config.Port,
	)

	dbCfg, err := pgxpool.ParseConfig(uri)
	if err != nil {
		panic("DB | Could not parse initial configuration: " + err.Error())
	}

	dbCfg.MaxConns = int32(config.PoolSize)
	dbCfg.MaxConnLifetime = 120 * time.Second

	connection, err := pgxpool.ConnectConfig(context.Background(), dbCfg)
	if err != nil {
		panic("DB | Could not connect to the database: " + err.Error())
	}

	if err = connection.Ping(context.TODO()); err != nil {
		panic("DB | Could not ping the database: " + err.Error())
	}

	return connection
}

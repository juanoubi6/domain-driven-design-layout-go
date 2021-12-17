package sql

import (
	"domain-driven-design-layout/infrastructure/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

const connStringTemplate = "host=%s user=%s dbname=%s password=%s port=%d sslmode=disable"

func CreateDatabaseConnection(config config.SQLConfig) *sqlx.DB {
	uri := fmt.Sprintf(
		connStringTemplate,
		config.Host,
		config.UserName,
		config.DbName,
		config.Password,
		config.Port,
	)

	db, err := sqlx.Connect("postgres", uri)
	if err != nil {
		log.Fatalln("DB | Could not connect to the database: " + err.Error())
	}

	return db
}

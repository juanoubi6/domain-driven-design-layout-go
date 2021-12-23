package repositories

import (
	"domain-driven-design-layout/infrastructure/config"
	"domain-driven-design-layout/infrastructure/repositories/sql"
	"fmt"
	"io/ioutil"
)

var db = sql.CreateDatabaseConnection(config.LoadAppConfig().SQLConfig)

func generateSchema() {
	content, err := ioutil.ReadFile("../../schema.sql")
	if err != nil {
		panic("Could not read schema file")
	}

	_, err = db.Exec(string(content))
	if err != nil {
		panic("Could not execute schema.sql file")
	}
}

func saveUserWithAddresses(userId int64) {
	insertUsersQuery := fmt.Sprintf(
		`INSERT INTO users (id, first_name, last_name, birth_date) VALUES 
				(%v,'test', 'user', '1995-07-20T00:00:00.000Z')`,
		userId,
	)

	insertAddressesQuery := fmt.Sprintf(
		`INSERT INTO addresses (street, number, user_id, city) VALUES 
			('Street 1', 1, %v, NULL), 
			('Street 2', 2, %v, 'Argentina')`,
		userId, userId,
	)

	_, _ = db.Exec(insertUsersQuery)
	_, _ = db.Exec(insertAddressesQuery)

}

package sql

import (
	"context"
	"domain-driven-design-layout/domain/entities"
	"domain-driven-design-layout/infrastructure/repositories/sql/models"
	"errors"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

func (qe *QueryExecutor) GetUser(ctx entities.ApplicationContext, id int64) (*entities.User, error) {
	var user entities.User

	start := time.Now()

	rows, err := qe.db.QueryxContext(context.TODO(), GetUserWithAddressesById, id)
	if err != nil {
		log.Printf("Error retrieving user of id %v: %v", id, err.Error())
		return nil, err
	}
	defer rows.Close()

	QueryTimeHistogram.WithLabelValues("GetUserWithAddressesById").Observe(time.Since(start).Seconds())

	for rows.Next() {
		var address entities.Address
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.BirthDate, &address.ID, &address.UserID, &address.Street, &address.Number, &address.City); err != nil {
			log.Printf("Error scanning row: %v", err.Error())
			return nil, err
		}

		user.Addresses = append(user.Addresses, address)
	}

	// User with ID 0 indicates no user was found
	if user.ID == 0 {
		return nil, nil
	}

	return &user, nil
}

func (qe *QueryExecutor) GetUsers(ctx entities.ApplicationContext, ids []int64) ([]entities.User, error) {
	query, args, err := sqlx.In(GetUsersWithAddressesByIds, ids)
	query = qe.db.Rebind(query)

	start := time.Now()

	rows, err := qe.db.QueryxContext(context.TODO(), query, args...)
	if err != nil {
		log.Printf("Error retrieving users of ids %v: %v", ids, err.Error())
		return nil, err
	}
	defer rows.Close()

	QueryTimeHistogram.WithLabelValues("GetUsersWithAddressesByIds").Observe(time.Since(start).Seconds())

	var userMap = make(map[int64]entities.User)

	for rows.Next() {
		var user entities.User
		var address entities.Address
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.BirthDate, &address.ID, &address.UserID, &address.Street, &address.Number, &address.City); err != nil {
			log.Printf("Error scanning row: %v", err.Error())
			return nil, err
		}

		if _, ok := userMap[user.ID]; ok {
			user := userMap[user.ID]
			user.Addresses = append(user.Addresses, address)
			userMap[user.ID] = user
		} else {
			user.Addresses = append(user.Addresses, address)
			userMap[user.ID] = user
		}
	}

	var users []entities.User
	for _, user := range userMap {
		users = append(users, user)
	}

	return users, nil
}

func (qe *QueryExecutor) CreateUser(ctx entities.ApplicationContext, prototype entities.UserPrototype) (entities.User, error) {
	// Start transaction to insert the user and it's addresses
	tx, err := qe.db.Beginx()
	if err != nil {
		log.Printf("Error creating transaction: %v", err.Error())
		return entities.User{}, err
	}

	var userId int64
	err = tx.QueryRowContext(context.Background(), InsertUser, prototype.FirstName, prototype.LastName, prototype.BirthDate).Scan(&userId)
	if err != nil {
		log.Printf("Error creating user: %v", err.Error())
		return entities.User{}, err
	}

	var addressModels []models.AddressModel
	for _, addressPrototype := range prototype.AddressesPrototypes {
		addressModels = append(addressModels, models.CreateAddressModelFromPrototype(addressPrototype, userId))
	}

	_, err = tx.NamedExecContext(context.TODO(), InsertAddresses, addressModels)
	if err != nil {
		tx.Rollback()
		log.Printf("Error when trying to insert addresses: %v", err.Error())
		return entities.User{}, err
	}

	// Finish transaction
	if err = tx.Commit(); err != nil {
		log.Printf("Error when committing tx: %v", err.Error())
		return entities.User{}, err
	}

	createdUser, err := qe.GetUser(ctx, userId)
	if err != nil {
		log.Printf("Error retrieving created user: %v", err.Error())
		return entities.User{}, err
	}

	return *createdUser, nil
}

func (qe *QueryExecutor) UpdateUser(ctx entities.ApplicationContext, user entities.User) (entities.User, error) {
	originalUser, err := qe.GetUser(ctx, user.ID)
	if err != nil {
		return user, err
	}
	if originalUser == nil {
		return user, errors.New("user not found")
	}

	originalUser.FirstName = user.FirstName
	originalUser.LastName = user.LastName
	originalUser.BirthDate = user.BirthDate

	_, err = qe.Exec(context.Background(), UpdateUser, originalUser.FirstName, originalUser.LastName, originalUser.BirthDate, originalUser.ID)
	if err != nil {
		log.Printf("Error updating user: %v", err.Error())
		return user, err
	}

	return *originalUser, nil
}

func (qe *QueryExecutor) DeleteUser(ctx entities.ApplicationContext, id int64) error {
	_, err := qe.Exec(context.Background(), DeleteUser, id)
	if err != nil {
		log.Printf("Error deleting user: %v", err.Error())
		return err
	}

	return nil
}

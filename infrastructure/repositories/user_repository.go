package repositories

import (
	"context"
	"domain-driven-design-layout/domain/entities"
	"domain-driven-design-layout/infrastructure/config"
	"domain-driven-design-layout/infrastructure/repositories/sql"
	"domain-driven-design-layout/infrastructure/repositories/sql/models"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type UserRepository struct {
	connectionPool *pgxpool.Pool
}

func NewUserRepository(config config.SQLConfig) (*UserRepository, error) {
	return &UserRepository{connectionPool: sql.CreateConnectionPool(config)}, nil
}

func (ur *UserRepository) GetUser(id int64) (*entities.User, error) {
	var user entities.User

	rows, err := ur.connectionPool.Query(context.TODO(), sql.GetUserWithAddressesById, id)
	if err != nil {
		log.Printf("Error retrieving user of id %v: %v", id, err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var address entities.Address
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.BirthDate, &address.ID, &address.UserID, &address.Street, &address.Number, &address.City); err != nil {
			log.Printf("Error scanning row: %v", err.Error())
			return nil, err
		}

		user.Addresses = append(user.Addresses, address)
	}

	return &user, nil
}

func (ur *UserRepository) GetUsers(ids []int64) ([]entities.User, error) {
	rows, err := ur.connectionPool.Query(context.TODO(), sql.GetUsersWithAddressesByIds, ids)
	if err != nil {
		log.Printf("Error retrieving users of ids %v: %v", ids, err.Error())
		return nil, err
	}
	defer rows.Close()

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

func (ur *UserRepository) CreateUser(prototype entities.UserPrototype) (entities.User, error) {
	var createdUser entities.User

	userModel := models.UserPrototypeToModel(prototype)

	// Start transaction to insert the user and it's addresses
	tx, err := ur.connectionPool.Begin(context.TODO())
	if err != nil {
		log.Printf("Error creating transaction: %v", err.Error())
		return createdUser, err
	}

	var userId int64
	err = tx.QueryRow(context.Background(), sql.InsertUser, userModel.FirstName, userModel.LastName, userModel.BirthDate).Scan(&userId)
	if err != nil {
		log.Printf("Error creating user: %v", err.Error())
		return createdUser, err
	}

	createdUser.ID = userId

	batch := pgx.Batch{}
	for _, addressPrototype := range prototype.AddressesPrototypes {
		addressModel := models.AddressPrototypeToModel(addressPrototype, createdUser.ID)
		batch.Queue(sql.InsertAddress, addressModel.UserID, addressModel.Street, addressModel.Number, addressModel.City)
	}

	batchResult := tx.SendBatch(context.TODO(), &batch)
	for i := 0; i < batch.Len(); i++ {
		_, err := batchResult.Exec()
		if err != nil {
			log.Printf("Error executing query from batch: %v", err.Error())
			tx.Rollback(context.TODO())
			return createdUser, err
		}
	}

	if err := batchResult.Close(); err != nil {
		log.Printf("Error closing address creation queries batch: %v", err.Error())
		tx.Rollback(context.TODO())
		return createdUser, err
	}

	// Finish transaction
	if err = tx.Commit(context.TODO()); err != nil {
		log.Printf("Error when committing tx: %v", err.Error())
		return createdUser, err
	}

	if err := ur.connectionPool.QueryRow(context.Background(), sql.GetUserById, createdUser.ID).Scan(
		&createdUser.ID,
		&createdUser.FirstName,
		&createdUser.LastName,
		&createdUser.BirthDate,
	); err != nil {
		log.Printf("Error retrieving user of id %v: %v", createdUser.ID, err.Error())
		return createdUser, err
	}

	rows, err := ur.connectionPool.Query(context.TODO(), sql.GetAddressesByUserId, createdUser.ID)
	if err != nil {
		log.Printf("Error retrieving addresses of user id %v: %v", createdUser.ID, err.Error())
		return createdUser, err
	}
	defer rows.Close()

	for rows.Next() {
		var address entities.Address
		if err := rows.Scan(&address.ID, &address.UserID, &address.Street, &address.Number, &address.City); err != nil {
			log.Printf("Error scanning address row: %v", err.Error())
			return createdUser, err
		}

		createdUser.Addresses = append(createdUser.Addresses, address)
	}

	return createdUser, nil
}

func (ur *UserRepository) UpdateUser(user entities.User) (entities.User, error) {
	originalUser, err := ur.GetUser(user.ID)
	if err != nil {
		return user, err
	}
	if originalUser == nil {
		return user, errors.New("user not found")
	}

	originalUser.FirstName = user.FirstName
	originalUser.LastName = user.LastName
	originalUser.BirthDate = user.BirthDate

	_, err = ur.connectionPool.Exec(context.Background(), sql.UpdateUser, originalUser.FirstName, originalUser.LastName, originalUser.BirthDate, originalUser.ID)
	if err != nil {
		log.Printf("Error updating user: %v", err.Error())
		return user, err
	}

	return *originalUser, nil
}

func (ur *UserRepository) DeleteUser(id int64) error {
	return nil
}

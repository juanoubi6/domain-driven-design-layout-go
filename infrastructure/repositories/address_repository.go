package repositories

import (
	"context"
	sql2 "database/sql"
	"domain-driven-design-layout/domain/entities"
	"domain-driven-design-layout/infrastructure/repositories/sql"
	"domain-driven-design-layout/infrastructure/repositories/sql/models"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type AddressRepository struct {
	queryExecutor QueryExecutor
}

func NewAddressRepository(db *sqlx.DB) (*AddressRepository, error) {
	return &AddressRepository{queryExecutor: QueryExecutor{db: db, tx: nil}}, nil
}

func (ur *AddressRepository) CreateAddress(userID int64, prototype entities.AddressPrototype) (entities.Address, error) {
	addressModel := models.CreateAddressModelFromPrototype(prototype, userID)

	var addressID int64
	err := ur.queryExecutor.db.QueryRowContext(context.Background(), sql.InsertAddress, addressModel.UserID, addressModel.Street, addressModel.Number, addressModel.City).Scan(&addressID)
	if err != nil {
		log.Printf("Error creating address: %v", err.Error())
		return entities.Address{}, err
	}

	addressModel.ID = addressID

	return addressModel.ToAddress(), nil
}

func (ur *AddressRepository) DeleteAddress(id int64) error {
	_, err := ur.queryExecutor.Exec(context.Background(), sql.DeleteAddress, id)
	if err != nil {
		log.Printf("Error deleting address: %v", err.Error())
		return err
	}

	return nil
}

func (ur *AddressRepository) GetAddress(id int64) (*entities.Address, error) {
	var address entities.Address
	var addressModel models.AddressModel

	start := time.Now()

	err := ur.queryExecutor.db.QueryRowxContext(context.TODO(), sql.GetAddressById, id).StructScan(&addressModel)
	if err != nil {
		if err == sql2.ErrNoRows {
			return nil, nil
		}

		log.Printf("Error retrieving address of id %v: %v", id, err.Error())
		return nil, err
	}

	sql.QueryTimeHistogram.WithLabelValues("GetAddress").Observe(time.Since(start).Seconds())

	address = addressModel.ToAddress()

	return &address, nil
}

func (ur *AddressRepository) DeleteUserAddresses(userID int64) error {
	_, err := ur.queryExecutor.Exec(context.Background(), sql.DeleteUserAddresses, userID)
	if err != nil {
		log.Printf("Error deleting user addresses: %v", err.Error())
		return err
	}

	return nil
}

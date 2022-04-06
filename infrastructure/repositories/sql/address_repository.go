package sql

import (
	"context"
	sql2 "database/sql"
	"domain-driven-design-layout/domain/entities"
	"domain-driven-design-layout/infrastructure/repositories/sql/models"
	"log"
	"time"
)

func (qe *QueryExecutor) CreateAddress(userID int64, prototype entities.AddressPrototype) (entities.Address, error) {
	addressModel := models.CreateAddressModelFromPrototype(prototype, userID)

	var addressID int64
	err := qe.db.QueryRowContext(context.Background(), InsertAddress, addressModel.UserID, addressModel.Street, addressModel.Number, addressModel.City).Scan(&addressID)
	if err != nil {
		log.Printf("Error creating address: %v", err.Error())
		return entities.Address{}, err
	}

	addressModel.ID = addressID

	return addressModel.ToAddress(), nil
}

func (qe *QueryExecutor) DeleteAddress(id int64) error {
	_, err := qe.Exec(context.Background(), DeleteAddress, id)
	if err != nil {
		log.Printf("Error deleting address: %v", err.Error())
		return err
	}

	return nil
}

func (qe *QueryExecutor) GetAddress(id int64) (*entities.Address, error) {
	var address entities.Address
	var addressModel models.AddressModel

	start := time.Now()

	err := qe.db.QueryRowxContext(context.TODO(), GetAddressById, id).StructScan(&addressModel)
	if err != nil {
		if err == sql2.ErrNoRows {
			return nil, nil
		}

		log.Printf("Error retrieving address of id %v: %v", id, err.Error())
		return nil, err
	}

	QueryTimeHistogram.WithLabelValues("GetAddress").Observe(time.Since(start).Seconds())

	address = addressModel.ToAddress()

	return &address, nil
}

func (qe *QueryExecutor) DeleteUserAddresses(userID int64) error {
	_, err := qe.Exec(context.Background(), DeleteUserAddresses, userID)
	if err != nil {
		log.Printf("Error deleting user addresses: %v", err.Error())
		return err
	}

	return nil
}

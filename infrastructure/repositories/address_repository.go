package repositories

import (
	"context"
	"domain-driven-design-layout/domain/entities"
	"domain-driven-design-layout/infrastructure/repositories/sql"
	"domain-driven-design-layout/infrastructure/repositories/sql/models"
	"github.com/jmoiron/sqlx"
	"log"
)

type AddressRepository struct {
	db *sqlx.DB
}

func NewAddressRepository(db *sqlx.DB) (*AddressRepository, error) {
	return &AddressRepository{db: db}, nil
}

func (ur *AddressRepository) CreateAddress(userID int64, prototype entities.AddressPrototype) (entities.Address, error) {
	addressModel := models.CreateAddressModelFromPrototype(prototype, userID)

	var addressID int64
	err := ur.db.QueryRowContext(context.Background(), sql.InsertAddress, addressModel.UserID, addressModel.Street, addressModel.Number, addressModel.City).Scan(&addressID)
	if err != nil {
		log.Printf("Error creating address: %v", err.Error())
		return entities.Address{}, err
	}

	addressModel.ID = addressID

	return addressModel.ToAddress(), nil
}

func (ur *AddressRepository) DeleteAddress(id int64) error {
	_, err := ur.db.ExecContext(context.Background(), sql.DeleteAddress, id)
	if err != nil {
		log.Printf("Error deleting address: %v", err.Error())
		return err
	}

	return nil
}

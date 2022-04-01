package db

import (
	"errors"

	"github.com/ekharisma/web-service-pp/model"
)

type Database interface {
	StoreTemperature(data model.Temperature) error
	GetLastTemperatures() (model.Temperature, error)
}

type InMemoryDatabase struct {
	database []model.Temperature
}

func NewInMemoryDatabase() Database {
	return &InMemoryDatabase{}
}

func (db *InMemoryDatabase) StoreTemperature(data model.Temperature) error {
	db.database = append(db.database, data)
	return nil
}

func (db *InMemoryDatabase) GetLastTemperatures() (model.Temperature, error) {
	if len(db.database) > 0 {
		return db.database[len(db.database)-1], nil
	}
	return model.Temperature{}, errors.New("No Data Found")
}

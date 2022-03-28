package db

import "github.com/ekharisma/web-service-pp/model"

var database []model.Temperature

func StoreTemperature(data model.Temperature) {
	database = append(database, data)
}

func GetLastTemperatures() []float32 {
	return database[len(database)-2:]
}

package db

import (
	"testing"
	"time"

	"github.com/ekharisma/web-service-pp/model"
)

func TestStoreData(t *testing.T) {
	db := NewInMemoryDatabase()
	mockData := model.Temperature{
		Timestamp:   time.Time{},
		Temperature: [2]float32{1, 2},
	}
	err := db.StoreTemperature(mockData)
	if err != nil {
		t.Error("Error. Reason : ", err.Error())
	}
}

func TestGetDataShouldSuccess(t *testing.T) {
	db := NewInMemoryDatabase()
	mockDatas := []model.Temperature{
		{
			Timestamp:   time.Time{},
			Temperature: [2]float32{1, 2},
		},
		{
			Timestamp:   time.Time{},
			Temperature: [2]float32{1, 2},
		},
	}
	for _, mockData := range mockDatas {
		db.StoreTemperature(mockData)
	}
	_, err := db.GetLastTemperatures()
	if err != nil {
		t.Error("Error. Reason : ", err.Error())
	}
}
func TestGetDataShouldError(t *testing.T) {
	db := NewInMemoryDatabase()
	_, err := db.GetLastTemperatures()
	if err == nil {
		t.Error("Should return nil")
	}
}

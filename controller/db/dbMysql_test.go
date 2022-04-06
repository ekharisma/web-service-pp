package db

import (
	"testing"
	"time"

	"github.com/ekharisma/web-service-pp/constant"
	"github.com/ekharisma/web-service-pp/model"
)

func TestStoreDataDB(t *testing.T) {
	db := NewMySQLDatabase(
		constant.UsernameDB, constant.PasswordDB, constant.HostDB,
		constant.NameDB, constant.PortDB,
	)
	mockData := model.Temperature{
		Timestamp:   time.Time{},
		Temperature: [2]float32{1, 2},
	}
	err := db.StoreTemperature(mockData)
	if err != nil {
		t.Error("Error. Reason : ", err.Error())
	}
}

func TestGetDataShouldSuccessDB(t *testing.T) {
	db := NewMySQLDatabase(
		constant.UsernameDB, constant.PasswordDB, constant.HostDB,
		constant.NameDB, constant.PortDB,
	)
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

// func TestGetDataShouldErrorDB(t *testing.T) {
// 	db := NewMySQLDatabase(
// 		constant.UsernameDB, constant.PasswordDB, constant.HostDB,
// 		constant.NameDB, constant.PortDB,
// 	)
// 	_, err := db.GetLastTemperatures()
// 	if err == nil {
// 		t.Error("Should return nil")
// 	}
// }

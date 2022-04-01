package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/ekharisma/web-service-pp/constant"
	"github.com/ekharisma/web-service-pp/model"
	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQLDatabase(username, password, host, dbName string, port uint) Database {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%d)/%v?parseTime=true", username, password, host, port, dbName)
	fmt.Println("Try to connect with dsn : ", dsn)
	db, err := sql.Open(constant.Driver, dsn)
	if err != nil {
		log.Panic("Can't connect to mysql. Reason : ", err.Error())
	}
	tx, txErr := db.BeginTx(context.Background(), &sql.TxOptions{})
	if txErr != nil {
		log.Fatal("Begin Transaction Error", txErr.Error())
	}
	query := `CREATE TABLE IF NOT EXISTS temperature_table(
		id INT AUTO_INCREMENT PRIMARY KEY,
		timestamp DATETIME,
		temperature1 FLOAT,
		temperature2 FLOAT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	) ENGINE=INNODB`
	_, txErr = tx.Exec(query)
	if txErr != nil {
		log.Fatal("Transaction Error, Rolling Back. Reason : ", txErr.Error())
		tx.Rollback()
	}
	if txErr = tx.Commit(); txErr != nil {
		log.Fatal("Commit Transaction Error. Reason : ", txErr.Error())
	}
	return &MySQL{
		db: db,
	}
}

func (db *MySQL) StoreTemperature(data model.Temperature) error {
	query := "INSERT INTO temperature_table(timestamp, temperature1, temperature2) VALUES(?, ?, ?)"
	fmt.Println("Query :", query)
	tx, txErr := db.db.BeginTx(context.Background(), &sql.TxOptions{})
	if txErr != nil {
		log.Fatal("Begin Transaction Error", txErr.Error())
		return txErr
	}
	_, txErr = tx.Exec(query, data.Timestamp.Format("2006-01-02 03:04:05"), data.Temperature[0], data.Temperature[1])
	if txErr != nil {
		tx.Rollback()
		log.Fatal("Transaction Error, Rolling Back. Reason : ", txErr.Error())
		return txErr
	}
	if txErr = tx.Commit(); txErr != nil {
		log.Fatal("Commit Transaction Error. Reason : ", txErr.Error())
		return txErr
	}
	return nil
}

func (db *MySQL) GetLastTemperatures() (model.Temperature, error) {
	query := "SELECT timestamp, temperature1, temperature2 FROM temperature_table ORDER BY Id DESC LIMIT 1"
	var temperature model.Temperature
	tx, txErr := db.db.BeginTx(context.Background(), &sql.TxOptions{})
	if txErr != nil {
		log.Fatal("Begin Transaction Error", txErr.Error())
		return model.Temperature{}, txErr
	}
	err := tx.QueryRow(query).Scan(&temperature.Timestamp, &temperature.Temperature[0], &temperature.Temperature[1])
	if err != nil {
		log.Fatal("Ca't Query DB. Reason : ", err.Error())
		return model.Temperature{}, err
	}
	if err = tx.Commit(); err != nil {
		log.Fatal("Cant Commit. Reason : ", err.Error())
		return model.Temperature{}, err
	}
	return temperature, nil
}

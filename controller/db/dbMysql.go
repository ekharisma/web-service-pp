package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ekharisma/web-service-pp/constant"
	"github.com/ekharisma/web-service-pp/model"
	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQLDatabase(username, password, host, dbName string, port uint) Database {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%d)/%v", username, password, host, port, dbName)
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

func (db *MySQL) StoreTemperature(data model.Temperature) {
	query := fmt.Sprintf("INSERT INTO temperature_table(timestamp, temperature1, temperature2) VALUES('%v', %v, %v)", dateParser(data.Timestamp), data.Temperature[0], data.Temperature[1])
	fmt.Println("Query :", query)
	tx, txErr := db.db.BeginTx(context.Background(), &sql.TxOptions{})
	if txErr != nil {
		log.Fatal("Begin Transaction Error", txErr.Error())
	}
	_, txErr = tx.Exec(query)
	if txErr != nil {
		tx.Rollback()
		log.Fatal("Transaction Error, Rolling Back. Reason : ", txErr.Error())
	}
	if txErr = tx.Commit(); txErr != nil {
		log.Fatal("Commit Transaction Error. Reason : ", txErr.Error())
	}
}

func (db *MySQL) GetLastTemperatures() (model.Temperature, error) {
	query := "SELECT timestamp, temperature1, temperature2 FROM temperature_table ORDER BY Id DESC LIMIT 1"
	tx, txErr := db.db.BeginTx(context.Background(), &sql.TxOptions{})
	if txErr != nil {
		log.Fatal("Begin Transaction Error", txErr.Error())
		return model.Temperature{}, txErr
	}
	results, txErr := tx.Query(query)
	if txErr != nil {
		tx.Rollback()
		log.Fatal("Transaction Error, Rolling Back. Reason : ", txErr.Error())
		return model.Temperature{}, txErr
	}
	if txErr = tx.Commit(); txErr != nil {
		log.Fatal("Commit Transaction Error. Reason : ", txErr.Error())
		return model.Temperature{}, txErr
	}
	var result = model.Temperature{}
	for results.Next() {
		var err = results.Scan(&result.Timestamp, &result.Temperature[0], &result.Temperature[1])
		if err != nil {
			log.Fatal(err.Error())
			return model.Temperature{}, err
		}
	}
	return result, nil
}

func dateParser(t time.Time) string {
	split := strings.Split(t.String(), " ")
	formatedStr := split[0] + " " + split[1]
	split = strings.Split(formatedStr, ".")
	return split[0]
}

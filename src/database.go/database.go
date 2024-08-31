package database

import (
	"database/sql"

	"api/src/config.go"

	_ "github.com/go-sql-driver/mysql" //driver
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DBConnString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

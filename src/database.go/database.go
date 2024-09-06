package database

import (
	"context"
	"database/sql"
	"time"

	"api/src/config.go"

	_ "github.com/go-sql-driver/mysql" //driver
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DBConnString)
	if err != nil {
		return nil, err
	}

	// Definindo o timeout para a conexão
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Verifica se a conexão com o banco está ativa
	if err = db.PingContext(ctx); err != nil {
		db.Close()

		return nil, err
	}

	return db, nil
}

package dto

import (
	"context"
	"database/sql"
	"log"
)

var db *sql.DB

const connectionString = "host=0.0.0.0 port=5432 dbname=postgres user=postgres password=password"

func Connect() error {
	var err error
	db, err = sql.Open("pgx", connectionString)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = db.Close()
		}
	}()

	return nil
}

func AccountExists(name string, db *sql.DB) bool {
	ctx := context.Background()

	rows, err := db.QueryContext(ctx, "SELECT name FROM accounts WHERE name = $1", name)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		return true
	}

	return false
}

func GetDB() *sql.DB {
	return db
}

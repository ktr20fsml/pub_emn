package database

import (
	_ "github.com/lib/pq"

	"fmt"

	"github.com/jmoiron/sqlx"
)

func NewDatabase() (*sqlx.DB, error) {
	db, errDB := sqlx.Connect("postgres", "host=postgres port=5432 user=postgres password=uoijlkm<?> dbname=emn sslmode=disable")
	if errDB != nil {
		return nil, fmt.Errorf("DATABASE CONNECTION ERROR: %s", errDB)
	}

	return db, nil
}

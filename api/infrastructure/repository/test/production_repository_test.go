package repository_test

import (
	"api/infrastructure/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func Test_NewProductionRepository(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("FAILED TO CREATE SQL MOCK: %s", err)
	}
	dbx := sqlx.NewDb(db, "sqlmock")
	defer db.Close()
	defer dbx.Close()

	repo := repository.NewProductionRepository(dbx)
	if repo == nil {
		t.Errorf("FAILED TO CREATE \"machine repository\".")
	}

}

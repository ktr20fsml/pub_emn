package gateway

import (
	"api/domain/repository"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

var txKey = struct{}{}

type TransactionRepository struct {
	conn *sqlx.DB
}

func (tr *TransactionRepository) GetDBConnection() *sqlx.DB {
	return tr.conn
}

func NewTransactionRepository(conn *sqlx.DB) repository.TransactionRepository {
	return &TransactionRepository{
		conn: conn,
	}
}

func (tr *TransactionRepository) ExecWtihTx(ctx context.Context, f func(ctx context.Context) (interface{}, error)) (interface{}, error) {
	conn := tr.GetDBConnection()
	// Create the transaction object.
	tx, err := conn.Beginx()
	if err != nil {
		return nil, fmt.Errorf("FAILED TO BEGIN TRANSACTION: %s", err.Error())
	}

	// Execute the following function with the transaction object as the argument.
	ctx = context.WithValue(ctx, &txKey, tx)
	v, err := f(ctx)
	if err != nil {
		_ = tx.Rollback()
		return v, fmt.Errorf("FAILED TO EXECUTE TRANSACTION: %s", err.Error())
	}

	// Execute commit transaction.
	if err := tx.Commit(); err != nil {
		_ = tx.Rollback()
		return v, fmt.Errorf("FAILED TO COMMIT: ROLLBACK: %s", err.Error())
	}

	return v, nil
}

func GetTx(ctx context.Context) (*sqlx.Tx, bool) {
	tx, ok := ctx.Value(&txKey).(*sqlx.Tx)
	return tx, ok
}

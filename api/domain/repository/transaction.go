package repository

import "context"

type TransactionRepository interface {
	ExecWtihTx(ctx context.Context, f func(ctx context.Context) (interface{}, error)) (interface{}, error)
}

package mock_repository

import (
	"api/domain/repository"
	"context"
)

type MockTransactionRepository struct {
	repository.TransactionRepository
	MockExecWithTx func(context.Context, func(context.Context) (interface{}, error)) (interface{}, error)
}

func (mtr *MockTransactionRepository) ExecWtihTx(ctx context.Context, f func(context.Context) (interface{}, error)) (interface{}, error) {
	return mtr.MockExecWithTx(ctx, f)
}

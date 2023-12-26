package mock_repository

import (
	domainGeneral "api/domain/model/general"
	"api/domain/repository"
	"context"
)

type MockGeneralRepository struct {
	repository.GeneralRepository
	MockCreateTableInformation func(context.Context, *domainGeneral.TableInformation) error
	MockUpdateTableInformation func(context.Context, *domainGeneral.TableInformation) error
}

func (mgr *MockGeneralRepository) CreateTableInformation(ctx context.Context, tableInfo *domainGeneral.TableInformation) error {
	return mgr.MockCreateTableInformation(ctx, tableInfo)
}

func (mgr *MockGeneralRepository) UpdateTableInformation(ctx context.Context, tableInfo *domainGeneral.TableInformation) error {
	return mgr.MockUpdateTableInformation(ctx, tableInfo)
}

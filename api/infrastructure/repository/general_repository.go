package repository

import (
	domainGeneral "api/domain/model/general"
	"api/domain/repository"
	"api/infrastructure/database/sql"
	dtoGeneral "api/infrastructure/dto/general"
	"api/interface/adapter/gateway"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type generalRepository struct {
	db *sqlx.DB
}

func NewGeneralRepository(db *sqlx.DB) repository.GeneralRepository {
	return &generalRepository{
		db: db,
	}
}

func (gr *generalRepository) CreateTableInformation(ctx context.Context, tableInfo *domainGeneral.TableInformation) error {
	dao, ok := gateway.GetTx(ctx)
	if !ok {
		_, err := gr.db.NamedExec(sql.InsertTableInformation, dtoGeneral.ConvertToTableInformationData(tableInfo))
		if err != nil {
			return fmt.Errorf("FAILED TO INSERT TABLE INFORMATION: %s", err.Error())
		}

		return nil
	}

	_, err := dao.NamedExec(sql.InsertTableInformation, dtoGeneral.ConvertToTableInformationData(tableInfo))
	if err != nil {
		return fmt.Errorf("FAILED TO INSERT TABLE INFORMATION: %s", err.Error())
	}

	return nil
}

func (gr *generalRepository) CreateTableInformations(ctx context.Context, tableInfos []*domainGeneral.TableInformation) error {
	dao, ok := gateway.GetTx(ctx)
	if !ok {
		_, err := gr.db.NamedExec(sql.InsertTableInformation, dtoGeneral.ConvertToTableInformationsDatas(tableInfos))
		if err != nil {
			return fmt.Errorf("FAILED TO INSERT TABLE INFORMATION: %s", err.Error())
		}

		return nil
	}

	_, err := dao.NamedExec(sql.InsertTableInformation, dtoGeneral.ConvertToTableInformationsDatas(tableInfos))
	if err != nil {
		return fmt.Errorf("FAILED TO INSERT TABLE INFORMATION: %s", err.Error())
	}

	return nil
}

func (gr *generalRepository) UpdateTableInformation(ctx context.Context, tableInfo *domainGeneral.TableInformation) error {
	dao, ok := gateway.GetTx(ctx)
	if !ok {
		_, err := gr.db.NamedExec(sql.UpdateTableInformation, dtoGeneral.ConvertToTableInformationData(tableInfo))
		if err != nil {
			return fmt.Errorf("FAILED TO UPDATE TABLE INFORMATION: %s", err.Error())
		}

		return nil
	}

	_, err := dao.NamedExec(sql.UpdateTableInformation, dtoGeneral.ConvertToTableInformationData(tableInfo))
	if err != nil {
		return fmt.Errorf("FAILED TO UPDATE TABLE INFORMATION: %s", err.Error())
	}

	return nil
}

func (gr *generalRepository) UpdateTableInformations(ctx context.Context, tableInfos []*domainGeneral.TableInformation) error {
	dao, ok := gateway.GetTx(ctx)
	if !ok {
		_, err := gr.db.NamedExec(sql.UpdateTableInformation, dtoGeneral.ConvertToTableInformationsDatas(tableInfos))
		if err != nil {
			return fmt.Errorf("FAILED TO UPDATE TABLE INFORMATIONS: %s", err.Error())
		}

		return nil
	}

	_, err := dao.NamedExec(sql.UpdateTableInformation, dtoGeneral.ConvertToTableInformationsDatas(tableInfos))
	if err != nil {
		return fmt.Errorf("FAILED TO UPDATE TABLE INFORMATIONS: %s", err.Error())
	}

	return nil
}

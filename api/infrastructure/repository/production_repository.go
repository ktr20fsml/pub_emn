package repository

import (
	domainProduction "api/domain/model/production"
	"api/domain/repository"
	stmt "api/infrastructure/database/sql"
	dtoProduction "api/infrastructure/dto/production"
	"api/interface/adapter/gateway"
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type productionRepository struct {
	db *sqlx.DB
}

func NewProductionRepository(db *sqlx.DB) repository.ProductionRepository {
	return &productionRepository{db}
}

func (pr *productionRepository) FindAllProductions() ([]*domainProduction.Production, error) {
	productions := []*dtoProduction.Production{}
	err := pr.db.Select(&productions, stmt.FindAllProductions)
	if err != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", err.Error())
	}

	return dtoProduction.ConvertToProductionsDomains(productions), nil
}

func (pr *productionRepository) FindProductionByID(id string) (*domainProduction.Production, error) {
	production := &dtoProduction.Production{}

	err := pr.db.Get(production, stmt.FindProductionByID, id)
	if err != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", err.Error())
	}

	return dtoProduction.ConvertToProductionDomain(production), nil
}

func (pr *productionRepository) FindProductionByItemID(id string) ([]*domainProduction.Production, error) {
	productions := []*dtoProduction.Production{}
	errDB := pr.db.Select(&productions, stmt.FindProductionByItemID, id)
	if errDB != nil {
		return nil, fmt.Errorf("SQL ERROR: %s", errDB)
	}

	return dtoProduction.ConvertToProductionsDomains(productions), nil
}

func (pr *productionRepository) CreateProduction(ctx context.Context, production *domainProduction.Production) error {
	dao, ok := gateway.GetTx(ctx)
	if !ok {
		_, err := pr.db.NamedExec(stmt.InsertProduction, dtoProduction.ConvertToProductionData(production))
		if err != nil {
			return fmt.Errorf("FAILED TO INSERT PRODUCTION: %s", err.Error())
		}

		return nil
	}

	_, err := dao.NamedExec(stmt.InsertProduction, dtoProduction.ConvertToProductionData(production))
	if err != nil {
		return fmt.Errorf("FAILED TO INSERT PRODUCTION: %s", err.Error())
	}

	return nil
}

func (pr *productionRepository) CreateConsumptionListID(ctx context.Context, id domainProduction.ConsumptionListID) error {
	var dao interface {
		Exec(query string, args ...interface{}) (sql.Result, error)
	}

	dao, ok := gateway.GetTx(ctx)
	if !ok {
		dao = pr.db
		// _, err := pr.db.Exec(stmt.InsertConsumptionListID, string(id))
		// if err != nil {
		// 	return fmt.Errorf("FAILED TO INSERT BSS CONSUMPTION LIST ID: %s", err.Error())
		// }

		// return nil
	}

	_, err := dao.Exec(stmt.InsertConsumptionListID, string(id))
	if err != nil {
		return fmt.Errorf("FAILED TO INSERT BSS CONSUMPTION LIST ID: %s", err.Error())
	}

	return nil
}

func (pr *productionRepository) CreateConsumptionList(ctx context.Context, csmpList []*domainProduction.ConsumptionList) error {
	var dao interface {
		NamedExec(query string, arg interface{}) (sql.Result, error)
	}

	dao, ok := gateway.GetTx(ctx)
	if !ok {
		dao = pr.db
		// _, err := pr.db.NamedExec(stmt.InsertConsumptionList, dtoProduction.ConvertToConsumptionListsDatas(csmpList))
		// if err != nil {
		// 	return fmt.Errorf("FAILED TO INSERT CONSUMPTION LIST: %s", err.Error())
		// }

		// return nil
	}

	_, err := dao.NamedExec(stmt.InsertConsumptionList, dtoProduction.ConvertToConsumptionListsDatas(csmpList))
	if err != nil {
		return fmt.Errorf("FAILED TO INSERT CONSUMPTION LIST: %s", err.Error())
	}

	return nil
}

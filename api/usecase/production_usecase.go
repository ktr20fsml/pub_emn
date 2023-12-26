package usecase

import (
	"api/domain/model/production"
	"api/domain/repository"
	"api/domain/service"

	"context"
)

type productionUsecase struct {
	transactionRepository repository.TransactionRepository
	productionRepository  repository.ProductionRepository
	productionService     service.ProductionService
	inventoryRepository   repository.InventoryRepository
	inventoryService      service.InventoryService
	generalRepository     repository.GeneralRepository
}

type ProductionUsecase interface {
	FindAllProductions() ([]*production.Production, error)
	FindProductionByID(id string) (*production.Production, error)
	FindProductionByItemID(itemID string) ([]*production.Production, error)
	CreateProduction(context.Context, *production.Production) error
}

func NewProductionUsecase(
	txRepo repository.TransactionRepository,
	productionRepo repository.ProductionRepository,
	productionServ service.ProductionService,
	inventoryRepo repository.InventoryRepository,
	inventoryServ service.InventoryService,
	generalRepo repository.GeneralRepository,
) ProductionUsecase {
	return &productionUsecase{
		transactionRepository: txRepo,
		productionRepository:  productionRepo,
		productionService:     productionServ,
		inventoryRepository:   inventoryRepo,
		inventoryService:      inventoryServ,
		generalRepository:     generalRepo,
	}
}

func (pu *productionUsecase) FindAllProductions() ([]*production.Production, error) {
	productions, err := pu.productionRepository.FindAllProductions()
	if err != nil {
		return nil, err
	}

	return productions, err
}

func (pu *productionUsecase) FindProductionByID(id string) (*production.Production, error) {
	production, err := pu.productionRepository.FindProductionByID(id)
	if err != nil {
		return nil, err
	}

	return production, err
}

func (pu *productionUsecase) FindProductionByItemID(id string) ([]*production.Production, error) {
	productions, err := pu.productionRepository.FindProductionByItemID(id)
	if err != nil {
		return nil, err
	}

	return productions, err
}

func (pu *productionUsecase) CreateProduction(ctx context.Context, p *production.Production) error {
	for _, c := range p.ConsumptionList {
		ok, errExist := pu.inventoryService.CheckExists(c.ItemID, c.ProcessID, c.Lot, c.Branch)
		if errExist != nil {
			return errExist
		}
		if !ok {
			return errExist
		}
	}
	preInventory, err := pu.productionService.Consump(p)
	if err != nil {
		return err
	}

	_, errTx := pu.transactionRepository.ExecWtihTx(ctx, func(ctx context.Context) (interface{}, error) {
		var err error

		err = pu.generalRepository.CreateTableInformation(ctx, &p.TableInformation)
		if err != nil {
			return nil, err
		}

		err = pu.productionRepository.CreateConsumptionListID(ctx, p.ConsumptionListID)
		if err != nil {
			return nil, err
		}

		err = pu.productionRepository.CreateConsumptionList(ctx, p.ConsumptionList)
		if err != nil {
			return nil, err
		}

		err = pu.productionRepository.CreateProduction(ctx, p)
		if err != nil {
			return nil, err
		}

		err = pu.inventoryRepository.UpsertInventories(ctx, preInventory)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})

	if errTx != nil {
		return errTx
	}

	return nil
}

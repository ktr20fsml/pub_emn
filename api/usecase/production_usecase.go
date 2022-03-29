package usecase

import (
	"api/domain/model/inventory"
	domainInventory "api/domain/model/inventory"
	"api/domain/model/production"
	"api/domain/repository"
	"api/domain/service"

	"context"
)

type productionUsecase struct {
	transactionRepository repository.TransactionRepository
	productionRepository  repository.ProductionRepository
	productionService     service.ProductionService
	itemRepository        repository.ItemRepository
	inventoryRepository   repository.InventoryRepository
	generalRepository     repository.GeneralRepository
}

type ProductionUsecase interface {
	FindAllProductions() ([]*production.Production, error)
	FindProductionByID(id string) (*production.Production, error)
	FindProductionByItemID(itemID string) ([]*production.Production, error)
	Consump(*production.Production) ([]*inventory.Inventory, error)
	CreateProduction(context.Context, *production.Production) error
}

func NewProductionUsecase(
	txRepo repository.TransactionRepository,
	productionRepo repository.ProductionRepository,
	productionServ service.ProductionService,
	itemRepo repository.ItemRepository,
	inventoryRepo repository.InventoryRepository,
	generalRepo repository.GeneralRepository,
) ProductionUsecase {
	return &productionUsecase{
		transactionRepository: txRepo,
		productionRepository:  productionRepo,
		productionService:     productionServ,
		itemRepository:        itemRepo,
		inventoryRepository:   inventoryRepo,
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

/*
	This "consump" method should not impliment in repository layer.
	Since this method has multiple repository layer's method.
	And not only necessary to write the process flow, but also to calculate the comsumption of materials.
*/
func (pu *productionUsecase) Consump(p *production.Production) ([]*inventory.Inventory, error) {
	// Define the variables "inventories" to return and "tmpInventories" to calculate.
	// The reason length of these slice add 1 is to store the result.
	inventories := make([]*inventory.Inventory, len(p.ConsumptionList)+1)

	// 1. Subtract consumptions from inventory.
	for i, c := range p.ConsumptionList {
		tmpInventory, errFindInventory := pu.inventoryRepository.FindInventory(c.ItemID, c.ProcessID, c.Lot, c.Branch)
		if errFindInventory != nil {
			return nil, errFindInventory
		}

		tmpInventory.NonDefectiveQty -= c.NonDefectiveQty
		tmpInventory.DefectiveQty -= c.DefectiveQty
		tmpInventory.SuspendedQty -= c.SuspendedQty
		tmpInventory.IsUsed = true
		tmpInventory.IsUsedUp = c.IsUsedUp

		inventories[i] = tmpInventory
	}

	tmpItem, errFindItem := pu.itemRepository.FindItemByID(p.ItemID)
	if errFindItem != nil {
		return nil, errFindItem
	}
	expirationDate := p.ProducedAt.AddDate(0, 0, int(tmpItem.ValidityDays))

	// 2. Add production in inventory.
	inventories[len(inventories)-1] = &domainInventory.Inventory{
		ItemID:          p.ItemID,
		ProcessID:       p.ProcessID,
		WarehouseID:     tmpItem.WarehouseID,
		Lot:             p.Lot,
		Branch:          p.Branch,
		NonDefectiveQty: p.NonDefectiveQty,
		DefectiveQty:    p.DefectiveQty,
		SuspendedQty:    p.SuspendedQty,
		ExpirationDate:  expirationDate,
		IsUsed:          false,
		IsUsedUp:        false,
	}

	return inventories, nil
}

func (pu *productionUsecase) CreateProduction(ctx context.Context, p *production.Production) error {
	for _, c := range p.ConsumptionList {
		ok, errExist := pu.productionService.CheckExistsInInventory(c.ItemID, c.ProcessID, c.Lot, c.Branch)
		if errExist != nil {
			return errExist
		}
		if !ok {
			return errExist
		}
	}
	preInventory, err := pu.Consump(p)
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

package service

import (
	domainInventory "api/domain/model/inventory"
	domainItem "api/domain/model/item"
	domainProduction "api/domain/model/production"
	"api/domain/repository"
	"api/domain/service"
	"api/infrastructure/database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type productionService struct {
	db                  *sqlx.DB
	itemRepository      repository.ItemRepository
	inventoryRepository repository.InventoryRepository
}

func NewProductionService(db *sqlx.DB, itemRepo repository.ItemRepository, inventoryRepo repository.InventoryRepository) service.ProductionService {
	return &productionService{
		db:                  db,
		itemRepository:      itemRepo,
		inventoryRepository: inventoryRepo,
	}
}

func (ps *productionService) Consump(production *domainProduction.Production) (*domainProduction.Production, []*domainInventory.Inventory, error) {
	inventories := make([]*domainInventory.Inventory, len(production.ConsumptionList)+1)
	for i, c := range production.ConsumptionList {
		tmpInventory, errFindInventory := ps.inventoryRepository.FindInventory(c.ItemID, c.ProcessID, c.Lot, c.Branch)
		if errFindInventory != nil {
			return nil, nil, errFindInventory
		}

		tmpInventory.NonDefectiveQty -= c.NonDefectiveQty
		tmpInventory.DefectiveQty -= c.DefectiveQty
		tmpInventory.SuspendedQty -= c.SuspendedQty
		tmpInventory.IsUsed = true
		tmpInventory.IsUsedUp = c.IsUsedUp

		inventories[i] = tmpInventory
	}

	item, errFindItem := ps.itemRepository.FindItemByID(production.ItemID)
	if errFindItem != nil {
		return nil, nil, errFindItem
	}
	expirationDate := production.ProducedAt.AddDate(0, 0, int(item.ValidityDays))

	tmpInventory := &domainInventory.Inventory{
		ItemID:          production.ItemID,
		ProcessID:       production.ProcessID,
		WarehouseID:     item.WarehouseID,
		Lot:             production.Lot,
		Branch:          production.Branch,
		NonDefectiveQty: production.NonDefectiveQty,
		DefectiveQty:    production.DefectiveQty,
		SuspendedQty:    production.SuspendedQty,
		ExpirationDate:  expirationDate,
		IsUsed:          false,
		IsUsedUp:        false,
	}
	inventories[len(inventories)-1] = tmpInventory

	return production, inventories, nil
}

func (ps *productionService) CheckExistsInInventory(itemID domainItem.ItemID, processID domainItem.ProcessID, lot string, branch string) (bool, error) {
	var count int

	err := ps.db.Get(&count, sql.CountItemInInventory, itemID, processID, lot, branch)
	if err != nil {
		return false, fmt.Errorf("FAILED TO FIND ITEM IN INVENTORY: %s", err.Error())
	}

	if count == 0 {
		return false, fmt.Errorf("NO ITEM")
	}

	return true, nil
}

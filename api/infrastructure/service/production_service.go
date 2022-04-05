package service

import (
	domainInventory "api/domain/model/inventory"
	domainProduction "api/domain/model/production"
	"api/domain/repository"
	"api/domain/service"
)

type productionService struct {
	itemRepository      repository.ItemRepository
	inventoryRepository repository.InventoryRepository
}

func NewProductionService(itemRepo repository.ItemRepository, inventoryRepo repository.InventoryRepository) service.ProductionService {
	return &productionService{
		itemRepository:      itemRepo,
		inventoryRepository: inventoryRepo,
	}
}

func (ps *productionService) Consump(production *domainProduction.Production) ([]*domainInventory.Inventory, error) {
	// Define the variables "inventories" to return and "tmpInventories" to calculate.
	// The reason length of these slice add 1 is to store the result.
	inventories := make([]*domainInventory.Inventory, len(production.ConsumptionList)+1)

	// 1. Subtract consumptions from inventory.
	for i, c := range production.ConsumptionList {
		tmpInventory, errFindInventory := ps.inventoryRepository.FindInventory(c.ItemID, c.ProcessID, c.Lot, c.Branch)
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

	tmpItem, errFindItem := ps.itemRepository.FindItemByID(production.ItemID)
	if errFindItem != nil {
		return nil, errFindItem
	}
	expirationDate := production.ProducedAt.AddDate(0, 0, int(tmpItem.ValidityDays))

	// 2. Add production in inventory.
	inventories[len(inventories)-1] = &domainInventory.Inventory{
		ItemID:          production.ItemID,
		ProcessID:       production.ProcessID,
		WarehouseID:     tmpItem.WarehouseID,
		Lot:             production.Lot,
		Branch:          production.Branch,
		NonDefectiveQty: production.NonDefectiveQty,
		DefectiveQty:    production.DefectiveQty,
		SuspendedQty:    production.SuspendedQty,
		ExpirationDate:  expirationDate,
		IsUsed:          false,
		IsUsedUp:        false,
	}

	return inventories, nil
}

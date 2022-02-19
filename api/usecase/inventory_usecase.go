package usecase

import (
	domainInventory "api/domain/model/inventory"
	"api/domain/repository"
)

type inventoryUsecase struct {
	inventoryRepository repository.InventoryRepository
	itemRepository      repository.ItemRepository
	locationRepository  repository.LocationRepository
	machineRepository   repository.MachineRepository
}

type InventoryUsecase interface {
	FindAllInventories() ([]*domainInventory.Inventory, error)
}

func NewInventoryUsecase(
	inventoryRepo repository.InventoryRepository,
	itemRepo repository.ItemRepository,
	machineRepo repository.MachineRepository,
	locationRepo repository.LocationRepository,
) InventoryUsecase {
	return &inventoryUsecase{
		inventoryRepository: inventoryRepo,
		itemRepository:      itemRepo,
		machineRepository:   machineRepo,
		locationRepository:  locationRepo,
	}
}

func (iu *inventoryUsecase) FindAllInventories() ([]*domainInventory.Inventory, error) {
	var err error

	inventories, err := iu.inventoryRepository.FindAllInventories()
	if err != nil {
		return nil, err
	}

	return inventories, nil
}

package service

import (
	domainItem "api/domain/model/item"
	"api/domain/repository"
	"api/domain/service"
)

type inventoryService struct {
	inventoryRepository repository.InventoryRepository
}

func NewInventoryService(inventoryRepo repository.InventoryRepository) service.InventoryService {
	return &inventoryService{
		inventoryRepository: inventoryRepo,
	}
}

func (is *inventoryService) CheckExists(itemID domainItem.ItemID, processID domainItem.ProcessID, lot string, branch string) (bool, error) {
	_, err := is.inventoryRepository.CheckExistsInventory(itemID, processID, lot, branch)
	if err != nil {
		return false, err
	}

	return true, nil
}

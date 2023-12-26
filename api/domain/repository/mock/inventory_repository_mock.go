package mock_repository

import (
	"api/domain/model/inventory"
	"api/domain/model/item"
	"api/domain/repository"
)

type MockInventoryRepository struct {
	repository.InventoryRepository
	MockFindInventory func(item.ItemID, item.ProcessID, string, string) (*inventory.Inventory, error)
}

func (mir *MockInventoryRepository) FindInventory(itemID item.ItemID, processID item.ProcessID, lot string, branch string) (*inventory.Inventory, error) {
	return mir.MockFindInventory(itemID, processID, lot, branch)
}

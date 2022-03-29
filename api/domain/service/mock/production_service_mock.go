package mock_service

import (
	"api/domain/model/item"
	"api/domain/service"
)

type MockProductionService struct {
	service.ProductionService
	MockCheckExistsInInventory func(item.ItemID, item.ProcessID, string, string) (bool, error)
}

func (mps *MockProductionService) CheckExistsInInventory(itemID item.ItemID, processID item.ProcessID, lot string, branch string) (bool, error) {
	return mps.MockCheckExistsInInventory(itemID, processID, lot, branch)
}

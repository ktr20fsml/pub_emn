package mock_service

import (
	domainInventory "api/domain/model/inventory"
	domainItem "api/domain/model/item"
	domainProduction "api/domain/model/production"
	"api/domain/service"
)

type MockProductionService struct {
	service.ProductionService
	MockCheckExistsInInventory func(domainItem.ItemID, domainItem.ProcessID, string, string) (bool, error)
	MockConsump                func(*domainProduction.Production) ([]*domainInventory.Inventory, error)
}

func (mps *MockProductionService) CheckExistsInInventory(itemID domainItem.ItemID, processID domainItem.ProcessID, lot string, branch string) (bool, error) {
	return mps.MockCheckExistsInInventory(itemID, processID, lot, branch)
}

func (mps *MockProductionService) Consump(production *domainProduction.Production) ([]*domainInventory.Inventory, error) {
	return mps.MockConsump(production)
}

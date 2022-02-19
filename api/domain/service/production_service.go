package service

import (
	domainInventory "api/domain/model/inventory"
	domainItem "api/domain/model/item"
	domainProduction "api/domain/model/production"
)

type ProductionService interface {
	Consump(*domainProduction.Production) (*domainProduction.Production, []*domainInventory.Inventory, error)
	CheckExistsInInventory(itemID domainItem.ItemID, processID domainItem.ProcessID, lot string, branch string) (bool, error)
}

package service

import (
	domainInventory "api/domain/model/inventory"
	domainProduction "api/domain/model/production"
)

type ProductionService interface {
	Consump(*domainProduction.Production) ([]*domainInventory.Inventory, error)
}

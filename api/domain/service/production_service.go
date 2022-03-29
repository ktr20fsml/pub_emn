package service

import (
	domainItem "api/domain/model/item"
)

type ProductionService interface {
	CheckExistsInInventory(itemID domainItem.ItemID, processID domainItem.ProcessID, lot string, branch string) (bool, error)
}

package service

import (
	domainItem "api/domain/model/item"
)

type InventoryService interface {
	CheckExists(domainItem.ItemID, domainItem.ProcessID, string, string) (bool, error)
}

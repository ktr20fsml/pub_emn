package service

import (
	domainItem "api/domain/model/item"
)

type InventoryService interface {
	CheckItemExists(itemID domainItem.Item, processID domainItem.ProcessID, lot string, branch string) bool
}

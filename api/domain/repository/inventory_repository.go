package repository

import (
	"api/domain/model/inventory"
	"api/domain/model/item"
	"context"
)

type InventoryRepository interface {
	FindInventory(itemID item.ItemID, processID item.ProcessID, lot string, branch string) (*inventory.Inventory, error)
	FindAllInventories() ([]*inventory.Inventory, error)
	CheckExistsInventory(item.ItemID, item.ProcessID, string, string) (bool, error)
	UpsertInventories(context.Context, []*inventory.Inventory) error
}

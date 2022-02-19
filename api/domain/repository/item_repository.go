package repository

import (
	domainItem "api/domain/model/item"
	"context"
)

type ItemRepository interface {
	// repository interfaces related to the item.
	FindItemByID(domainItem.ItemID) (*domainItem.Item, error)
	FindItemInDetailByID(domainItem.ItemID) (*domainItem.Item, error)
	FindItemsByName(string) ([]*domainItem.Item, error)
	FindAllItems() ([]*domainItem.Item, error)
	FindAllItemsInDetail() ([]*domainItem.Item, error)
	CreateItem(context.Context, *domainItem.Item) error
	UpdateItem(*domainItem.Item) error
	StopUsingItem(*domainItem.Item) error

	// repository interfaces related to the item category.
	FindItemCategoryByID(domainItem.CategoryID) (*domainItem.Category, error)
	FindAllItemCategories() ([]*domainItem.Category, error)
	CreateItemCategory(*domainItem.Category) error
	UpdateItemCategory(*domainItem.Category) error
	StopUsingItemCategory(*domainItem.Category) error

	// repository interfaces related to the item process.
	FindItemStatusByID(domainItem.StatusID) (*domainItem.Status, error)
	FindAllItemStatuses() ([]*domainItem.Status, error)
	CreateItemStatus(*domainItem.Status) error
	UpdateItemStatus(*domainItem.Status) error
	StopUsingItemStatus(*domainItem.Status) error

	// repository interfaces related to the item unit.
	FindItemUnitByID(domainItem.UnitID) (*domainItem.Unit, error)
	FindAllItemUnits() ([]*domainItem.Unit, error)
	CreateItemUnit(*domainItem.Unit) error
	UpdateItemUnit(*domainItem.Unit) error
	StopUsingItemUnit(*domainItem.Unit) error

	// repository interfaces related to the item process.
	FindItemProcessByID(domainItem.ProcessID) (*domainItem.Process, error)
	FindAllItemProcesses() ([]*domainItem.Process, error)
	CreateItemProcess(*domainItem.Process) error
	UpdateItemProcess(*domainItem.Process) error
	StopUsingItemProcess(*domainItem.Process) error
}

package inventory

import (
	domainInventory "api/domain/model/inventory"
	domainItem "api/domain/model/item"
	domainWarehouse "api/domain/model/location"
	"api/infrastructure/dto/item"
	"api/infrastructure/dto/location"
	"time"
	// count "api/infrastructure/dto/entity/master/inventorycount"
)

type Inventory struct {
	ItemID          string             `db:"mst_item_id"`
	Item            item.Item          `db:"item"`
	ProcessID       string             `db:"mst_process_id"`
	Process         item.Process       `db:"process"`
	Lot             string             `db:"lot"`
	Branch          string             `db:"branch"`
	WarehouseID     string             `db:"mst_warehouse_id"`
	Warehouse       location.Warehouse `db:"warehouse"`
	NonDefectiveQty float32            `db:"non_defective_qty"`
	DefectiveQty    float32            `db:"defective_qty"`
	SuspendedQty    float32            `db:"suspended_qty"`
	ExpirationDate  time.Time          `db:"expiration_date"`
	IsUsed          bool               `db:"is_used"`
	IsUsedUp        bool               `db:"is_used_up"`
	RegulateID      string             `db:"bss_regulate_inventory_id"`
	Regulates       []*Regulate        `db:""`
	// InventoryCount count.ID
}

func ConvertToInventoryData(reqInventory *domainInventory.Inventory) *Inventory {
	return &Inventory{
		ItemID:          string(reqInventory.ItemID),
		Item:            *item.ConvertToItemData(&reqInventory.Item),
		ProcessID:       string(reqInventory.ProcessID),
		Process:         *item.ConvertToProcessData(&reqInventory.Process),
		Lot:             reqInventory.Lot,
		Branch:          reqInventory.Branch,
		WarehouseID:     string(reqInventory.WarehouseID),
		Warehouse:       *location.ConvertToWarehouseData(&reqInventory.Warehouse),
		NonDefectiveQty: reqInventory.NonDefectiveQty,
		DefectiveQty:    reqInventory.DefectiveQty,
		SuspendedQty:    reqInventory.SuspendedQty,
		ExpirationDate:  reqInventory.ExpirationDate,
		IsUsed:          reqInventory.IsUsed,
		IsUsedUp:        reqInventory.IsUsedUp,
		RegulateID:      string(reqInventory.RegulateID),
		Regulates:       ConvertToRegulatesDatas(reqInventory.Regulates),
	}
}

func ConvertToInventoriesDatas(reqInventories []*domainInventory.Inventory) []*Inventory {
	inventories := make([]*Inventory, len(reqInventories))

	for i, reqInventory := range reqInventories {
		inventories[i] = ConvertToInventoryData(reqInventory)
	}

	return inventories
}

func ConvertToInventoryDomain(inventory *Inventory) *domainInventory.Inventory {
	return &domainInventory.Inventory{
		ItemID:          domainItem.ItemID(inventory.ItemID),
		Item:            *item.ConvertToItemDomain(&inventory.Item),
		ProcessID:       domainItem.ProcessID(inventory.ProcessID),
		Process:         *item.ConvertToProcessDomain(&inventory.Process),
		Lot:             inventory.Lot,
		Branch:          inventory.Branch,
		WarehouseID:     domainWarehouse.WarehouseID(inventory.WarehouseID),
		Warehouse:       *location.ConvertToWarehouseDomain(&inventory.Warehouse),
		NonDefectiveQty: inventory.NonDefectiveQty,
		DefectiveQty:    inventory.DefectiveQty,
		SuspendedQty:    inventory.SuspendedQty,
		ExpirationDate:  inventory.ExpirationDate,
		IsUsed:          inventory.IsUsed,
		IsUsedUp:        inventory.IsUsedUp,
		RegulateID:      domainInventory.RegulateID(inventory.RegulateID),
		Regulates:       ConvertToRegulatesDomains(inventory.Regulates),
	}
}

func ConvertToInventoriesDomains(inventories []*Inventory) []*domainInventory.Inventory {
	resInventories := make([]*domainInventory.Inventory, len(inventories))

	for i, inventory := range inventories {
		resInventories[i] = ConvertToInventoryDomain(inventory)
	}

	return resInventories
}

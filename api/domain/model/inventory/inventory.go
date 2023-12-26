package inventory

import (
	"api/domain/model/item"
	"api/domain/model/location"
	"time"
)

type Inventory struct {
	ItemID          item.ItemID          `json:"inventoryItemID"`
	Item            item.Item            `json:"inventoryItem"`
	ProcessID       item.ProcessID       `json:"inventoryProcessID"`
	Process         item.Process         `json:"inventoryProcess"`
	Lot             string               `json:"inventoryLot"`
	Branch          string               `json:"inventoryBranch"`
	WarehouseID     location.WarehouseID `json:"inventoryWarehouseID"`
	Warehouse       location.Warehouse   `json:"inventoryWarehouse"`
	NonDefectiveQty float32              `json:"inventoryNonDefectiveQty"`
	DefectiveQty    float32              `json:"inventoryDefectiveQty"`
	SuspendedQty    float32              `json:"inventorySuspendedQty"`
	ExpirationDate  time.Time            `json:"inventoryExpirationDate"`
	IsUsed          bool                 `json:"inventoryIsUsed"`
	IsUsedUp        bool                 `json:"inventoryIsUsedUp"`
	RegulateID      RegulateID           `json:"inventoryRegulateID"`
	Regulates       []*Regulate          `json:"inventoryRegulates"`
	// InventoryCount count.ID
}
